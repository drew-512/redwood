// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	blob "redwood.dev/blob"

	mock "github.com/stretchr/testify/mock"

	process "redwood.dev/process"

	protoblob "redwood.dev/swarm/protoblob"

	swarm "redwood.dev/swarm"

	types "redwood.dev/types"
)

// BlobTransport is an autogenerated mock type for the BlobTransport type
type BlobTransport struct {
	mock.Mock
}

// AnnounceBlob provides a mock function with given fields: ctx, blobID
func (_m *BlobTransport) AnnounceBlob(ctx context.Context, blobID blob.ID) error {
	ret := _m.Called(ctx, blobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, blob.ID) error); ok {
		r0 = rf(ctx, blobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Autoclose provides a mock function with given fields:
func (_m *BlobTransport) Autoclose() {
	_m.Called()
}

// AutocloseWithCleanup provides a mock function with given fields: closeFn
func (_m *BlobTransport) AutocloseWithCleanup(closeFn func()) {
	_m.Called(closeFn)
}

// Close provides a mock function with given fields:
func (_m *BlobTransport) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ctx provides a mock function with given fields:
func (_m *BlobTransport) Ctx() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Done provides a mock function with given fields:
func (_m *BlobTransport) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Go provides a mock function with given fields: ctx, name, fn
func (_m *BlobTransport) Go(ctx context.Context, name string, fn func(context.Context)) <-chan struct{} {
	ret := _m.Called(ctx, name, fn)

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func(context.Context, string, func(context.Context)) <-chan struct{}); ok {
		r0 = rf(ctx, name, fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *BlobTransport) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewChild provides a mock function with given fields: ctx, name
func (_m *BlobTransport) NewChild(ctx context.Context, name string) *process.Process {
	ret := _m.Called(ctx, name)

	var r0 *process.Process
	if rf, ok := ret.Get(0).(func(context.Context, string) *process.Process); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*process.Process)
		}
	}

	return r0
}

// NewPeerConn provides a mock function with given fields: ctx, dialAddr
func (_m *BlobTransport) NewPeerConn(ctx context.Context, dialAddr string) (swarm.PeerConn, error) {
	ret := _m.Called(ctx, dialAddr)

	var r0 swarm.PeerConn
	if rf, ok := ret.Get(0).(func(context.Context, string) swarm.PeerConn); ok {
		r0 = rf(ctx, dialAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(swarm.PeerConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dialAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnBlobChunkRequest provides a mock function with given fields: handler
func (_m *BlobTransport) OnBlobChunkRequest(handler func(types.Hash, protoblob.BlobPeerConn)) {
	_m.Called(handler)
}

// OnBlobManifestRequest provides a mock function with given fields: handler
func (_m *BlobTransport) OnBlobManifestRequest(handler func(blob.ID, protoblob.BlobPeerConn)) {
	_m.Called(handler)
}

// ProcessTree provides a mock function with given fields:
func (_m *BlobTransport) ProcessTree() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// ProvidersOfBlob provides a mock function with given fields: ctx, blobID
func (_m *BlobTransport) ProvidersOfBlob(ctx context.Context, blobID blob.ID) (<-chan protoblob.BlobPeerConn, error) {
	ret := _m.Called(ctx, blobID)

	var r0 <-chan protoblob.BlobPeerConn
	if rf, ok := ret.Get(0).(func(context.Context, blob.ID) <-chan protoblob.BlobPeerConn); ok {
		r0 = rf(ctx, blobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan protoblob.BlobPeerConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, blob.ID) error); ok {
		r1 = rf(ctx, blobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpawnChild provides a mock function with given fields: ctx, child
func (_m *BlobTransport) SpawnChild(ctx context.Context, child process.Spawnable) error {
	ret := _m.Called(ctx, child)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, process.Spawnable) error); ok {
		r0 = rf(ctx, child)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *BlobTransport) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// State provides a mock function with given fields:
func (_m *BlobTransport) State() process.State {
	ret := _m.Called()

	var r0 process.State
	if rf, ok := ret.Get(0).(func() process.State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(process.State)
	}

	return r0
}
