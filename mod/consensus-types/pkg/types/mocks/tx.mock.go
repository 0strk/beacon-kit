// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	proto "github.com/cosmos/gogoproto/proto"
	mock "github.com/stretchr/testify/mock"
)

// Tx is an autogenerated mock type for the Tx type
type Tx struct {
	mock.Mock
}

type Tx_Expecter struct {
	mock *mock.Mock
}

func (_m *Tx) EXPECT() *Tx_Expecter {
	return &Tx_Expecter{mock: &_m.Mock}
}

// Bytes provides a mock function with given fields:
func (_m *Tx) Bytes() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Bytes")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Tx_Bytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bytes'
type Tx_Bytes_Call struct {
	*mock.Call
}

// Bytes is a helper method to define mock.On call
func (_e *Tx_Expecter) Bytes() *Tx_Bytes_Call {
	return &Tx_Bytes_Call{Call: _e.mock.On("Bytes")}
}

func (_c *Tx_Bytes_Call) Run(run func()) *Tx_Bytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_Bytes_Call) Return(_a0 []byte) *Tx_Bytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Bytes_Call) RunAndReturn(run func() []byte) *Tx_Bytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetGasLimit provides a mock function with given fields:
func (_m *Tx) GetGasLimit() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGasLimit")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tx_GetGasLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGasLimit'
type Tx_GetGasLimit_Call struct {
	*mock.Call
}

// GetGasLimit is a helper method to define mock.On call
func (_e *Tx_Expecter) GetGasLimit() *Tx_GetGasLimit_Call {
	return &Tx_GetGasLimit_Call{Call: _e.mock.On("GetGasLimit")}
}

func (_c *Tx_GetGasLimit_Call) Run(run func()) *Tx_GetGasLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_GetGasLimit_Call) Return(_a0 uint64, _a1 error) *Tx_GetGasLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tx_GetGasLimit_Call) RunAndReturn(run func() (uint64, error)) *Tx_GetGasLimit_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessages provides a mock function with given fields:
func (_m *Tx) GetMessages() ([]proto.Message, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMessages")
	}

	var r0 []proto.Message
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]proto.Message, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []proto.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]proto.Message)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tx_GetMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessages'
type Tx_GetMessages_Call struct {
	*mock.Call
}

// GetMessages is a helper method to define mock.On call
func (_e *Tx_Expecter) GetMessages() *Tx_GetMessages_Call {
	return &Tx_GetMessages_Call{Call: _e.mock.On("GetMessages")}
}

func (_c *Tx_GetMessages_Call) Run(run func()) *Tx_GetMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_GetMessages_Call) Return(_a0 []proto.Message, _a1 error) *Tx_GetMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tx_GetMessages_Call) RunAndReturn(run func() ([]proto.Message, error)) *Tx_GetMessages_Call {
	_c.Call.Return(run)
	return _c
}

// GetSenders provides a mock function with given fields:
func (_m *Tx) GetSenders() ([][]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSenders")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([][]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() [][]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tx_GetSenders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSenders'
type Tx_GetSenders_Call struct {
	*mock.Call
}

// GetSenders is a helper method to define mock.On call
func (_e *Tx_Expecter) GetSenders() *Tx_GetSenders_Call {
	return &Tx_GetSenders_Call{Call: _e.mock.On("GetSenders")}
}

func (_c *Tx_GetSenders_Call) Run(run func()) *Tx_GetSenders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_GetSenders_Call) Return(_a0 [][]byte, _a1 error) *Tx_GetSenders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tx_GetSenders_Call) RunAndReturn(run func() ([][]byte, error)) *Tx_GetSenders_Call {
	_c.Call.Return(run)
	return _c
}

// Hash provides a mock function with given fields:
func (_m *Tx) Hash() [32]byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Hash")
	}

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func() [32]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	return r0
}

// Tx_Hash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hash'
type Tx_Hash_Call struct {
	*mock.Call
}

// Hash is a helper method to define mock.On call
func (_e *Tx_Expecter) Hash() *Tx_Hash_Call {
	return &Tx_Hash_Call{Call: _e.mock.On("Hash")}
}

func (_c *Tx_Hash_Call) Run(run func()) *Tx_Hash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_Hash_Call) Return(_a0 [32]byte) *Tx_Hash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Hash_Call) RunAndReturn(run func() [32]byte) *Tx_Hash_Call {
	_c.Call.Return(run)
	return _c
}

// NewTx creates a new instance of Tx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tx {
	mock := &Tx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
