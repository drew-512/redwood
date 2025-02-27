// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	blob "redwood.dev/blob"

	mock "github.com/stretchr/testify/mock"

	process "redwood.dev/process"

	protoblob "redwood.dev/swarm/protoblob"
)

// BlobProtocol is an autogenerated mock type for the BlobProtocol type
type BlobProtocol struct {
	mock.Mock
}

// Autoclose provides a mock function with given fields:
func (_m *BlobProtocol) Autoclose() {
	_m.Called()
}

// AutocloseWithCleanup provides a mock function with given fields: closeFn
func (_m *BlobProtocol) AutocloseWithCleanup(closeFn func()) {
	_m.Called(closeFn)
}

// Close provides a mock function with given fields:
func (_m *BlobProtocol) Close() error {
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
func (_m *BlobProtocol) Ctx() context.Context {
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
func (_m *BlobProtocol) Done() <-chan struct{} {
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
func (_m *BlobProtocol) Go(ctx context.Context, name string, fn func(context.Context)) <-chan struct{} {
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
func (_m *BlobProtocol) Name() string {
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
func (_m *BlobProtocol) NewChild(ctx context.Context, name string) *process.Process {
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

// ProcessTree provides a mock function with given fields:
func (_m *BlobProtocol) ProcessTree() map[string]interface{} {
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
func (_m *BlobProtocol) ProvidersOfBlob(ctx context.Context, blobID blob.ID) <-chan protoblob.BlobPeerConn {
	ret := _m.Called(ctx, blobID)

	var r0 <-chan protoblob.BlobPeerConn
	if rf, ok := ret.Get(0).(func(context.Context, blob.ID) <-chan protoblob.BlobPeerConn); ok {
		r0 = rf(ctx, blobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan protoblob.BlobPeerConn)
		}
	}

	return r0
}

// SpawnChild provides a mock function with given fields: ctx, child
func (_m *BlobProtocol) SpawnChild(ctx context.Context, child process.Spawnable) error {
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
func (_m *BlobProtocol) Start() error {
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
func (_m *BlobProtocol) State() process.State {
	ret := _m.Called()

	var r0 process.State
	if rf, ok := ret.Get(0).(func() process.State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(process.State)
	}

	return r0
}
