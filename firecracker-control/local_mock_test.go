// Code generated by MockGen. DO NOT EDIT.
// Source: local.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	events "github.com/containerd/containerd/events"
	firecracker_go_sdk "github.com/firecracker-microvm/firecracker-go-sdk"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockmachine is a mock of machine interface
type Mockmachine struct {
	ctrl     *gomock.Controller
	recorder *MockmachineMockRecorder
}

// MockmachineMockRecorder is the mock recorder for Mockmachine
type MockmachineMockRecorder struct {
	mock *Mockmachine
}

// NewMockmachine creates a new mock instance
func NewMockmachine(ctrl *gomock.Controller) *Mockmachine {
	mock := &Mockmachine{ctrl: ctrl}
	mock.recorder = &MockmachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockmachine) EXPECT() *MockmachineMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *Mockmachine) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockmachineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockmachine)(nil).Start), arg0)
}

// StopVMM mocks base method
func (m *Mockmachine) StopVMM() error {
	ret := m.ctrl.Call(m, "StopVMM")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopVMM indicates an expected call of StopVMM
func (mr *MockmachineMockRecorder) StopVMM() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopVMM", reflect.TypeOf((*Mockmachine)(nil).StopVMM))
}

// Shutdown mocks base method
func (m *Mockmachine) Shutdown(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockmachineMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*Mockmachine)(nil).Shutdown), arg0)
}

// Wait mocks base method
func (m *Mockmachine) Wait(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockmachineMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*Mockmachine)(nil).Wait), arg0)
}

// SetMetadata mocks base method
func (m *Mockmachine) SetMetadata(arg0 context.Context, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "SetMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetadata indicates an expected call of SetMetadata
func (mr *MockmachineMockRecorder) SetMetadata(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadata", reflect.TypeOf((*Mockmachine)(nil).SetMetadata), arg0, arg1)
}

// UpdateGuestDrive mocks base method
func (m *Mockmachine) UpdateGuestDrive(arg0 context.Context, arg1, arg2 string, arg3 ...firecracker_go_sdk.PatchGuestDriveByIDOpt) error {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGuestDrive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGuestDrive indicates an expected call of UpdateGuestDrive
func (mr *MockmachineMockRecorder) UpdateGuestDrive(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGuestDrive", reflect.TypeOf((*Mockmachine)(nil).UpdateGuestDrive), varargs...)
}

// Mockpublisher is a mock of publisher interface
type Mockpublisher struct {
	ctrl     *gomock.Controller
	recorder *MockpublisherMockRecorder
}

// MockpublisherMockRecorder is the mock recorder for Mockpublisher
type MockpublisherMockRecorder struct {
	mock *Mockpublisher
}

// NewMockpublisher creates a new mock instance
func NewMockpublisher(ctrl *gomock.Controller) *Mockpublisher {
	mock := &Mockpublisher{ctrl: ctrl}
	mock.recorder = &MockpublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockpublisher) EXPECT() *MockpublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *Mockpublisher) Publish(ctx context.Context, topic string, event events.Event) error {
	ret := m.ctrl.Call(m, "Publish", ctx, topic, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockpublisherMockRecorder) Publish(ctx, topic, event interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*Mockpublisher)(nil).Publish), ctx, topic, event)
}
