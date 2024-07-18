// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	context "context"

	flow "github.com/onflow/flow-go/model/flow"
	metrics "github.com/slok/go-http-metrics/metrics"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AccessMetrics is an autogenerated mock type for the AccessMetrics type
type AccessMetrics struct {
	mock.Mock
}

// AddInflightRequests provides a mock function with given fields: ctx, props, quantity
func (_m *AccessMetrics) AddInflightRequests(ctx context.Context, props metrics.HTTPProperties, quantity int) {
	_m.Called(ctx, props, quantity)
}

// AddTotalRequests provides a mock function with given fields: ctx, method, routeName
func (_m *AccessMetrics) AddTotalRequests(ctx context.Context, method string, routeName string) {
	_m.Called(ctx, method, routeName)
}

// ConnectionAddedToPool provides a mock function with given fields:
func (_m *AccessMetrics) ConnectionAddedToPool() {
	_m.Called()
}

// ConnectionFromPoolEvicted provides a mock function with given fields:
func (_m *AccessMetrics) ConnectionFromPoolEvicted() {
	_m.Called()
}

// ConnectionFromPoolInvalidated provides a mock function with given fields:
func (_m *AccessMetrics) ConnectionFromPoolInvalidated() {
	_m.Called()
}

// ConnectionFromPoolReused provides a mock function with given fields:
func (_m *AccessMetrics) ConnectionFromPoolReused() {
	_m.Called()
}

// ConnectionFromPoolUpdated provides a mock function with given fields:
func (_m *AccessMetrics) ConnectionFromPoolUpdated() {
	_m.Called()
}

// NewConnectionEstablished provides a mock function with given fields:
func (_m *AccessMetrics) NewConnectionEstablished() {
	_m.Called()
}

// ObserveHTTPRequestDuration provides a mock function with given fields: ctx, props, duration
func (_m *AccessMetrics) ObserveHTTPRequestDuration(ctx context.Context, props metrics.HTTPReqProperties, duration time.Duration) {
	_m.Called(ctx, props, duration)
}

// ObserveHTTPResponseSize provides a mock function with given fields: ctx, props, sizeBytes
func (_m *AccessMetrics) ObserveHTTPResponseSize(ctx context.Context, props metrics.HTTPReqProperties, sizeBytes int64) {
	_m.Called(ctx, props, sizeBytes)
}

// ScriptExecuted provides a mock function with given fields: dur, size
func (_m *AccessMetrics) ScriptExecuted(dur time.Duration, size int) {
	_m.Called(dur, size)
}

// ScriptExecutionErrorLocal provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionErrorLocal() {
	_m.Called()
}

// ScriptExecutionErrorMatch provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionErrorMatch() {
	_m.Called()
}

// ScriptExecutionErrorMismatch provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionErrorMismatch() {
	_m.Called()
}

// ScriptExecutionErrorOnExecutionNode provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionErrorOnExecutionNode() {
	_m.Called()
}

// ScriptExecutionNotIndexed provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionNotIndexed() {
	_m.Called()
}

// ScriptExecutionResultMatch provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionResultMatch() {
	_m.Called()
}

// ScriptExecutionResultMismatch provides a mock function with given fields:
func (_m *AccessMetrics) ScriptExecutionResultMismatch() {
	_m.Called()
}

// TotalConnectionsInPool provides a mock function with given fields: connectionCount, connectionPoolSize
func (_m *AccessMetrics) TotalConnectionsInPool(connectionCount uint, connectionPoolSize uint) {
	_m.Called(connectionCount, connectionPoolSize)
}

// TransactionExecuted provides a mock function with given fields: txID, when
func (_m *AccessMetrics) TransactionExecuted(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionExpired provides a mock function with given fields: txID
func (_m *AccessMetrics) TransactionExpired(txID flow.Identifier) {
	_m.Called(txID)
}

// TransactionFinalized provides a mock function with given fields: txID, when
func (_m *AccessMetrics) TransactionFinalized(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionReceived provides a mock function with given fields: txID, when
func (_m *AccessMetrics) TransactionReceived(txID flow.Identifier, when time.Time) {
	_m.Called(txID, when)
}

// TransactionResultFetched provides a mock function with given fields: dur, size
func (_m *AccessMetrics) TransactionResultFetched(dur time.Duration, size int) {
	_m.Called(dur, size)
}

// TransactionSubmissionFailed provides a mock function with given fields:
func (_m *AccessMetrics) TransactionSubmissionFailed() {
	_m.Called()
}

// TransactionValidated provides a mock function with given fields: txID
func (_m *AccessMetrics) TransactionValidated(txID flow.Identifier) {
	_m.Called(txID)
}

// TransactionValidationFailed provides a mock function with given fields: txID, reason
func (_m *AccessMetrics) TransactionValidationFailed(txID flow.Identifier, reason string) {
	_m.Called(txID, reason)
}

// UpdateExecutionReceiptMaxHeight provides a mock function with given fields: height
func (_m *AccessMetrics) UpdateExecutionReceiptMaxHeight(height uint64) {
	_m.Called(height)
}

// UpdateLastFullBlockHeight provides a mock function with given fields: height
func (_m *AccessMetrics) UpdateLastFullBlockHeight(height uint64) {
	_m.Called(height)
}

// NewAccessMetrics creates a new instance of AccessMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccessMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccessMetrics {
	mock := &AccessMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
