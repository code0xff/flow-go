// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// GlobalParams is an autogenerated mock type for the GlobalParams type
type GlobalParams struct {
	mock.Mock
}

// ChainID provides a mock function with given fields:
func (_m *GlobalParams) ChainID() (flow.ChainID, error) {
	ret := _m.Called()

	var r0 flow.ChainID
	if rf, ok := ret.Get(0).(func() flow.ChainID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flow.ChainID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpochCommitSafetyThreshold provides a mock function with given fields:
func (_m *GlobalParams) EpochCommitSafetyThreshold() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProtocolVersion provides a mock function with given fields:
func (_m *GlobalParams) ProtocolVersion() (uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SporkID provides a mock function with given fields:
func (_m *GlobalParams) SporkID() (flow.Identifier, error) {
	ret := _m.Called()

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewGlobalParamsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewGlobalParams creates a new instance of GlobalParams. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGlobalParams(t NewGlobalParamsT) *GlobalParams {
	mock := &GlobalParams{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
