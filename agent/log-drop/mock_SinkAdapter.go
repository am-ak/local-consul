// Code generated by mockery v2.12.2. DO NOT EDIT.

package logdrop

import (
	hclog "github.com/hashicorp/go-hclog"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockSinkAdapter is an autogenerated mock type for the SinkAdapter type
type MockSinkAdapter struct {
	mock.Mock
}

// Accept provides a mock function with given fields: name, level, msg, args
func (_m *MockSinkAdapter) Accept(name string, level hclog.Level, msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, level, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// NewMockSinkAdapter creates a new instance of MockSinkAdapter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSinkAdapter(t testing.TB) *MockSinkAdapter {
	mock := &MockSinkAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
