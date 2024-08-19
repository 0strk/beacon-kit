// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	types "github.com/berachain/beacon-kit/mod/async/pkg/types"
	pruner "github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
	mock "github.com/stretchr/testify/mock"
)

// BlockEvent is an autogenerated mock type for the BlockEvent type
type BlockEvent[BeaconBlockT pruner.BeaconBlock] struct {
	mock.Mock
}

type BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock] struct {
	mock *mock.Mock
}

func (_m *BlockEvent[BeaconBlockT]) EXPECT() *BlockEvent_Expecter[BeaconBlockT] {
	return &BlockEvent_Expecter[BeaconBlockT]{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *BlockEvent[BeaconBlockT]) Data() BeaconBlockT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 BeaconBlockT
	if rf, ok := ret.Get(0).(func() BeaconBlockT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BeaconBlockT)
	}

	return r0
}

// BlockEvent_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type BlockEvent_Data_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[BeaconBlockT]) Data() *BlockEvent_Data_Call[BeaconBlockT] {
	return &BlockEvent_Data_Call[BeaconBlockT]{Call: _e.mock.On("Data")}
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) Run(run func()) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) Return(_a0 BeaconBlockT) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) RunAndReturn(run func() BeaconBlockT) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// Is provides a mock function with given fields: _a0
func (_m *BlockEvent[BeaconBlockT]) Is(_a0 types.EventID) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Is")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.EventID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BlockEvent_Is_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Is'
type BlockEvent_Is_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// Is is a helper method to define mock.On call
//   - _a0 types.EventID
func (_e *BlockEvent_Expecter[BeaconBlockT]) Is(_a0 interface{}) *BlockEvent_Is_Call[BeaconBlockT] {
	return &BlockEvent_Is_Call[BeaconBlockT]{Call: _e.mock.On("Is", _a0)}
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) Run(run func(_a0 types.EventID)) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.EventID))
	})
	return _c
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) Return(_a0 bool) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) RunAndReturn(run func(types.EventID) bool) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// NewBlockEvent creates a new instance of BlockEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockEvent[BeaconBlockT pruner.BeaconBlock](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockEvent[BeaconBlockT] {
	mock := &BlockEvent[BeaconBlockT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
