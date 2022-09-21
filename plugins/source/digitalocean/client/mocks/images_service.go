// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/digitalocean/client (interfaces: ImagesService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockImagesService is a mock of ImagesService interface.
type MockImagesService struct {
	ctrl     *gomock.Controller
	recorder *MockImagesServiceMockRecorder
}

// MockImagesServiceMockRecorder is the mock recorder for MockImagesService.
type MockImagesServiceMockRecorder struct {
	mock *MockImagesService
}

// NewMockImagesService creates a new mock instance.
func NewMockImagesService(ctrl *gomock.Controller) *MockImagesService {
	mock := &MockImagesService{ctrl: ctrl}
	mock.recorder = &MockImagesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImagesService) EXPECT() *MockImagesServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockImagesService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Image, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]godo.Image)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockImagesServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockImagesService)(nil).List), arg0, arg1)
}
