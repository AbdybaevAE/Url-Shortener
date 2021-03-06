// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repos/link/type.go

// Package mock_link is a generated GoMock package.
package mock_link

import (
	reflect "reflect"

	models "github.com/abdybaevae/url-shortener/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockLinkRepo is a mock of LinkRepo interface.
type MockLinkRepo struct {
	ctrl     *gomock.Controller
	recorder *MockLinkRepoMockRecorder
}

// MockLinkRepoMockRecorder is the mock recorder for MockLinkRepo.
type MockLinkRepoMockRecorder struct {
	mock *MockLinkRepo
}

// NewMockLinkRepo creates a new mock instance.
func NewMockLinkRepo(ctrl *gomock.Controller) *MockLinkRepo {
	mock := &MockLinkRepo{ctrl: ctrl}
	mock.recorder = &MockLinkRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkRepo) EXPECT() *MockLinkRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockLinkRepo) Get(key string) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLinkRepoMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLinkRepo)(nil).Get), key)
}

// Save mocks base method.
func (m *MockLinkRepo) Save(link *models.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockLinkRepoMockRecorder) Save(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockLinkRepo)(nil).Save), link)
}
