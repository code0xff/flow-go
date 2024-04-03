// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"
)

// StateMutator is an autogenerated mock type for the StateMutator type
type StateMutator struct {
	mock.Mock
}

// ApplyServiceEventsFromValidatedSeals provides a mock function with given fields: seals
func (_m *StateMutator) ApplyServiceEventsFromValidatedSeals(seals []*flow.Seal) error {
	ret := _m.Called(seals)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*flow.Seal) error); ok {
		r0 = rf(seals)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Build provides a mock function with given fields:
func (_m *StateMutator) Build() (flow.Identifier, protocol.DeferredDBUpdates, error) {
	ret := _m.Called()

	var r0 flow.Identifier
	var r1 protocol.DeferredDBUpdates
	var r2 error
	if rf, ok := ret.Get(0).(func() (flow.Identifier, protocol.DeferredDBUpdates, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	if rf, ok := ret.Get(1).(func() protocol.DeferredDBUpdates); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(protocol.DeferredDBUpdates)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewStateMutator interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateMutator creates a new instance of StateMutator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateMutator(t mockConstructorTestingTNewStateMutator) *StateMutator {
	mock := &StateMutator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
