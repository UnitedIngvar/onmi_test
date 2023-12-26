// Code generated by mockery v2.39.1. DO NOT EDIT.

package client

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// GetLimits provides a mock function with given fields:
func (_m *MockService) GetLimits() (uint64, time.Duration) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLimits")
	}

	var r0 uint64
	var r1 time.Duration
	if rf, ok := ret.Get(0).(func() (uint64, time.Duration)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() time.Duration); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(time.Duration)
	}

	return r0, r1
}

// MockService_GetLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimits'
type MockService_GetLimits_Call struct {
	*mock.Call
}

// GetLimits is a helper method to define mock.On call
func (_e *MockService_Expecter) GetLimits() *MockService_GetLimits_Call {
	return &MockService_GetLimits_Call{Call: _e.mock.On("GetLimits")}
}

func (_c *MockService_GetLimits_Call) Run(run func()) *MockService_GetLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_GetLimits_Call) Return(n uint64, p time.Duration) *MockService_GetLimits_Call {
	_c.Call.Return(n, p)
	return _c
}

func (_c *MockService_GetLimits_Call) RunAndReturn(run func() (uint64, time.Duration)) *MockService_GetLimits_Call {
	_c.Call.Return(run)
	return _c
}

// Process provides a mock function with given fields: ctx, batch
func (_m *MockService) Process(ctx context.Context, batch Batch) error {
	ret := _m.Called(ctx, batch)

	if len(ret) == 0 {
		panic("no return value specified for Process")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Batch) error); ok {
		r0 = rf(ctx, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Process_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Process'
type MockService_Process_Call struct {
	*mock.Call
}

// Process is a helper method to define mock.On call
//   - ctx context.Context
//   - batch Batch
func (_e *MockService_Expecter) Process(ctx interface{}, batch interface{}) *MockService_Process_Call {
	return &MockService_Process_Call{Call: _e.mock.On("Process", ctx, batch)}
}

func (_c *MockService_Process_Call) Run(run func(ctx context.Context, batch Batch)) *MockService_Process_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Batch))
	})
	return _c
}

func (_c *MockService_Process_Call) Return(_a0 error) *MockService_Process_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_Process_Call) RunAndReturn(run func(context.Context, Batch) error) *MockService_Process_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
