// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	mock "github.com/stretchr/testify/mock"
)

// BeaconBlock is an autogenerated mock type for the BeaconBlock type
type BeaconBlock struct {
	mock.Mock
}

type BeaconBlock_Expecter struct {
	mock *mock.Mock
}

func (_m *BeaconBlock) EXPECT() *BeaconBlock_Expecter {
	return &BeaconBlock_Expecter{mock: &_m.Mock}
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlock) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlock_GetSlot_Call struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetSlot() *BeaconBlock_GetSlot_Call {
	return &BeaconBlock_GetSlot_Call{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlock_GetSlot_Call) Run(run func()) *BeaconBlock_GetSlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) Return(_a0 math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) RunAndReturn(run func() math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlock creates a new instance of BeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlock {
	mock := &BeaconBlock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
