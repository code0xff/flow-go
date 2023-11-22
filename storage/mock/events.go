// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/onflow/flow-go/storage"
)

// Events is an autogenerated mock type for the Events type
type Events struct {
	mock.Mock
}

// BatchRemoveByBlockID provides a mock function with given fields: blockID, batch
func (_m *Events) BatchRemoveByBlockID(blockID flow.Identifier, batch storage.BatchStorage) error {
	ret := _m.Called(blockID, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, storage.BatchStorage) error); ok {
		r0 = rf(blockID, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchStore provides a mock function with given fields: blockID, events, batch
func (_m *Events) BatchStore(blockID flow.Identifier, events []flow.EventsList, batch storage.BatchStorage) error {
	ret := _m.Called(blockID, events, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, []flow.EventsList, storage.BatchStorage) error); ok {
		r0 = rf(blockID, events, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ByBlockID provides a mock function with given fields: blockID
func (_m *Events) ByBlockID(blockID flow.Identifier) ([]flow.Event, error) {
	ret := _m.Called(blockID)

	var r0 []flow.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) ([]flow.Event, error)); ok {
		return rf(blockID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) []flow.Event); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByBlockIDEventType provides a mock function with given fields: blockID, eventType
func (_m *Events) ByBlockIDEventType(blockID flow.Identifier, eventType flow.EventType) ([]flow.Event, error) {
	ret := _m.Called(blockID, eventType)

	var r0 []flow.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.EventType) ([]flow.Event, error)); ok {
		return rf(blockID, eventType)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.EventType) []flow.Event); ok {
		r0 = rf(blockID, eventType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, flow.EventType) error); ok {
		r1 = rf(blockID, eventType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByBlockIDTransactionID provides a mock function with given fields: blockID, transactionID
func (_m *Events) ByBlockIDTransactionID(blockID flow.Identifier, transactionID flow.Identifier) ([]flow.Event, error) {
	ret := _m.Called(blockID, transactionID)

	var r0 []flow.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) ([]flow.Event, error)); ok {
		return rf(blockID, transactionID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) []flow.Event); ok {
		r0 = rf(blockID, transactionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, flow.Identifier) error); ok {
		r1 = rf(blockID, transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByBlockIDTransactionIndex provides a mock function with given fields: blockID, txIndex
func (_m *Events) ByBlockIDTransactionIndex(blockID flow.Identifier, txIndex uint32) ([]flow.Event, error) {
	ret := _m.Called(blockID, txIndex)

	var r0 []flow.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, uint32) ([]flow.Event, error)); ok {
		return rf(blockID, txIndex)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, uint32) []flow.Event); ok {
		r0 = rf(blockID, txIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, uint32) error); ok {
		r1 = rf(blockID, txIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: blockID, blockEvents
func (_m *Events) Store(blockID flow.Identifier, blockEvents []flow.EventsList) error {
	ret := _m.Called(blockID, blockEvents)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, []flow.EventsList) error); ok {
		r0 = rf(blockID, blockEvents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEvents interface {
	mock.TestingT
	Cleanup(func())
}

// NewEvents creates a new instance of Events. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEvents(t mockConstructorTestingTNewEvents) *Events {
	mock := &Events{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
