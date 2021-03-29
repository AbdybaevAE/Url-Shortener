// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/number/type.go

// Package mock_number is a generated GoMock package.
package mock_number

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNumberService is a mock of NumberService interface.
type MockNumberService struct {
	ctrl     *gomock.Controller
	recorder *MockNumberServiceMockRecorder
}

// MockNumberServiceMockRecorder is the mock recorder for MockNumberService.
type MockNumberServiceMockRecorder struct {
	mock *MockNumberService
}

// NewMockNumberService creates a new mock instance.
func NewMockNumberService(ctrl *gomock.Controller) *MockNumberService {
	mock := &MockNumberService{ctrl: ctrl}
	mock.recorder = &MockNumberServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNumberService) EXPECT() *MockNumberServiceMockRecorder {
	return m.recorder
}

// Increment mocks base method.
func (m *MockNumberService) Increment(number_id, byValue int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Increment", number_id, byValue)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment.
func (mr *MockNumberServiceMockRecorder) Increment(number_id, byValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockNumberService)(nil).Increment), number_id, byValue)
}