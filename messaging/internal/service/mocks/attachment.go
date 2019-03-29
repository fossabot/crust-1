// Code generated by MockGen. DO NOT EDIT.
// Source: messaging/internal/service/attachment.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	service "github.com/crusttech/crust/messaging/internal/service"
	types "github.com/crusttech/crust/messaging/types"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockAttachmentService is a mock of AttachmentService interface
type MockAttachmentService struct {
	ctrl     *gomock.Controller
	recorder *MockAttachmentServiceMockRecorder
}

// MockAttachmentServiceMockRecorder is the mock recorder for MockAttachmentService
type MockAttachmentServiceMockRecorder struct {
	mock *MockAttachmentService
}

// NewMockAttachmentService creates a new mock instance
func NewMockAttachmentService(ctrl *gomock.Controller) *MockAttachmentService {
	mock := &MockAttachmentService{ctrl: ctrl}
	mock.recorder = &MockAttachmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttachmentService) EXPECT() *MockAttachmentServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockAttachmentService) With(ctx context.Context) service.AttachmentService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(service.AttachmentService)
	return ret0
}

// With indicates an expected call of With
func (mr *MockAttachmentServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockAttachmentService)(nil).With), ctx)
}

// FindByID mocks base method
func (m *MockAttachmentService) FindByID(id uint64) (*types.Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockAttachmentServiceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAttachmentService)(nil).FindByID), id)
}

// Create mocks base method
func (m *MockAttachmentService) Create(name string, size int64, fh io.ReadSeeker, channelId, replyTo uint64) (*types.Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, size, fh, channelId, replyTo)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAttachmentServiceMockRecorder) Create(name, size, fh, channelId, replyTo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAttachmentService)(nil).Create), name, size, fh, channelId, replyTo)
}

// OpenOriginal mocks base method
func (m *MockAttachmentService) OpenOriginal(att *types.Attachment) (io.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenOriginal", att)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenOriginal indicates an expected call of OpenOriginal
func (mr *MockAttachmentServiceMockRecorder) OpenOriginal(att interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenOriginal", reflect.TypeOf((*MockAttachmentService)(nil).OpenOriginal), att)
}

// OpenPreview mocks base method
func (m *MockAttachmentService) OpenPreview(att *types.Attachment) (io.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPreview", att)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPreview indicates an expected call of OpenPreview
func (mr *MockAttachmentServiceMockRecorder) OpenPreview(att interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPreview", reflect.TypeOf((*MockAttachmentService)(nil).OpenPreview), att)
}