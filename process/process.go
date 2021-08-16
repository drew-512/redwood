package process

import (
	"context"
	"errors"
	"sync"
	"time"

	"redwood.dev/log"
	"redwood.dev/utils"
)

type Process struct {
	log.Logger

	name     string
	mu       sync.RWMutex
	children map[Spawnable]struct{}
	state    State

	goroutines map[*goroutine]struct{}

	closeOnce   sync.Once
	doneOnce    sync.Once
	chStop      chan struct{}
	chDone      chan struct{}
	chAutoclose chan struct{}
	wg          sync.WaitGroup

	debug bool
}

var _ Interface = (*Process)(nil)

type Interface interface {
	Spawnable
	ProcessTreer
	Autoclose()
	Ctx() context.Context
	NewChild(ctx context.Context, name string) *Process
	SpawnChild(ctx context.Context, child Spawnable) error
	Go(name string, fn func(ctx context.Context)) <-chan struct{}
}

type Spawnable interface {
	Name() string
	Start() error
	Close() error
	Done() <-chan struct{}
}

type ProcessTreer interface {
	Spawnable
	ProcessTree() map[string]interface{}
}

var _ Spawnable = (*Process)(nil)

func New(name string) *Process {
	return &Process{
		name:       name + " #" + utils.RandomNumberString(),
		children:   make(map[Spawnable]struct{}),
		goroutines: make(map[*goroutine]struct{}),
		chStop:     make(chan struct{}),
		chDone:     make(chan struct{}),
		Logger:     log.NewLogger(""),
	}
}

func (p *Process) Start() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.state != Unstarted {
		panic("already started")
	}
	p.state = Started

	return nil
}

func (p *Process) Close() error {
	p.closeOnce.Do(func() {
		func() {
			p.mu.Lock()
			defer p.mu.Unlock()

			if p.state != Started {
				panic("not started")
			}
			p.state = Closed
		}()

		close(p.chStop)
		p.wg.Wait()
		close(p.chDone)
	})
	return nil
}

func (p *Process) Autoclose() {
	go func() {
		p.wg.Wait()
		p.Close()
	}()
}

func (p *Process) Name() string {
	return p.name
}

func (p *Process) Ctx() context.Context {
	return utils.ChanContext(p.chStop)
}

func (p *Process) ProcessTree() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var goroutines []string
	for goroutine := range p.goroutines {
		goroutines = append(goroutines, goroutine.name)
	}

	children := make(map[string]interface{}, len(p.children))
	for child := range p.children {
		if child, is := child.(ProcessTreer); is {
			children[child.Name()] = child.ProcessTree()
		} else {
			children[child.Name()] = nil
		}
	}
	return map[string]interface{}{
		"status":     p.state.String(),
		"goroutines": goroutines,
		"children":   children,
	}
}

func (p *Process) NewChild(ctx context.Context, name string) *Process {
	child := New(name)
	_ = p.SpawnChild(ctx, child)
	return child
}

var (
	ErrUnstarted = errors.New("unstarted")
	ErrClosed    = errors.New("closed")
)

func (p *Process) SpawnChild(ctx context.Context, child Spawnable) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.state != Started {
		return nil
	}

	err := child.Start()
	if err != nil {
		return err
	}
	p.children[child] = struct{}{}

	p.wg.Add(1)
	go func() {
		defer func() {
			p.wg.Done()
			p.mu.Lock()
			delete(p.children, child)
			p.mu.Unlock()
		}()

		var ctxDone <-chan struct{}
		if ctx != nil {
			ctxDone = ctx.Done()
		}

		select {
		case <-p.chStop:
			child.Close()
		case <-ctxDone:
			child.Close()
		case <-child.Done():
		}
	}()

	return nil
}

type goroutine struct {
	name   string
	chDone chan struct{}
}

func (p *Process) Go(name string, fn func(ctx context.Context)) <-chan struct{} {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.state != Started {
		ch := make(chan struct{})
		close(ch)
		return ch
	}

	g := &goroutine{name: name + " #" + utils.RandomNumberString(), chDone: make(chan struct{})}
	p.goroutines[g] = struct{}{}

	go func() {
		for {
			select {
			case <-p.chStop:
				// fmt.Println("SHOULD STOP", g.name)
			case <-g.chDone:
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()

	p.wg.Add(1)
	go func() {
		defer func() {
			p.wg.Done()
			p.mu.Lock()
			defer p.mu.Unlock()
			delete(p.goroutines, g)
			close(g.chDone)
		}()

		fn(utils.ChanContext(p.chStop))
	}()

	return g.chDone
}

func (p *Process) Done() <-chan struct{} {
	return p.chDone
}

type State int

const (
	Unstarted State = iota
	Started
	Closed
)

func (s State) String() string {
	switch s {
	case Unstarted:
		return "unstarted"
	case Started:
		return "started"
	case Closed:
		return "closed"
	default:
		return "(err: unknown)"
	}
}
