// Code generated by mockery v2.46.3. DO NOT EDIT.

package host

import (
	context "context"

	host "github.com/Mellanox/doca-driver-build/entrypoint/internal/utils/host"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

type Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *Interface) EXPECT() *Interface_Expecter {
	return &Interface_Expecter{mock: &_m.Mock}
}

// GetDebugInfo provides a mock function with given fields: ctx
func (_m *Interface) GetDebugInfo(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDebugInfo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_GetDebugInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDebugInfo'
type Interface_GetDebugInfo_Call struct {
	*mock.Call
}

// GetDebugInfo is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Interface_Expecter) GetDebugInfo(ctx interface{}) *Interface_GetDebugInfo_Call {
	return &Interface_GetDebugInfo_Call{Call: _e.mock.On("GetDebugInfo", ctx)}
}

func (_c *Interface_GetDebugInfo_Call) Run(run func(ctx context.Context)) *Interface_GetDebugInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Interface_GetDebugInfo_Call) Return(_a0 string, _a1 error) *Interface_GetDebugInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_GetDebugInfo_Call) RunAndReturn(run func(context.Context) (string, error)) *Interface_GetDebugInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetOSType provides a mock function with given fields: ctx
func (_m *Interface) GetOSType(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetOSType")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_GetOSType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOSType'
type Interface_GetOSType_Call struct {
	*mock.Call
}

// GetOSType is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Interface_Expecter) GetOSType(ctx interface{}) *Interface_GetOSType_Call {
	return &Interface_GetOSType_Call{Call: _e.mock.On("GetOSType", ctx)}
}

func (_c *Interface_GetOSType_Call) Run(run func(ctx context.Context)) *Interface_GetOSType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Interface_GetOSType_Call) Return(_a0 string, _a1 error) *Interface_GetOSType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_GetOSType_Call) RunAndReturn(run func(context.Context) (string, error)) *Interface_GetOSType_Call {
	_c.Call.Return(run)
	return _c
}

// LsMod provides a mock function with given fields: ctx
func (_m *Interface) LsMod(ctx context.Context) (map[string]host.LoadedModule, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LsMod")
	}

	var r0 map[string]host.LoadedModule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]host.LoadedModule, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]host.LoadedModule); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]host.LoadedModule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_LsMod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LsMod'
type Interface_LsMod_Call struct {
	*mock.Call
}

// LsMod is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Interface_Expecter) LsMod(ctx interface{}) *Interface_LsMod_Call {
	return &Interface_LsMod_Call{Call: _e.mock.On("LsMod", ctx)}
}

func (_c *Interface_LsMod_Call) Run(run func(ctx context.Context)) *Interface_LsMod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Interface_LsMod_Call) Return(_a0 map[string]host.LoadedModule, _a1 error) *Interface_LsMod_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_LsMod_Call) RunAndReturn(run func(context.Context) (map[string]host.LoadedModule, error)) *Interface_LsMod_Call {
	_c.Call.Return(run)
	return _c
}

// RmMod provides a mock function with given fields: ctx, module
func (_m *Interface) RmMod(ctx context.Context, module string) error {
	ret := _m.Called(ctx, module)

	if len(ret) == 0 {
		panic("no return value specified for RmMod")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Interface_RmMod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RmMod'
type Interface_RmMod_Call struct {
	*mock.Call
}

// RmMod is a helper method to define mock.On call
//   - ctx context.Context
//   - module string
func (_e *Interface_Expecter) RmMod(ctx interface{}, module interface{}) *Interface_RmMod_Call {
	return &Interface_RmMod_Call{Call: _e.mock.On("RmMod", ctx, module)}
}

func (_c *Interface_RmMod_Call) Run(run func(ctx context.Context, module string)) *Interface_RmMod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Interface_RmMod_Call) Return(_a0 error) *Interface_RmMod_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Interface_RmMod_Call) RunAndReturn(run func(context.Context, string) error) *Interface_RmMod_Call {
	_c.Call.Return(run)
	return _c
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
