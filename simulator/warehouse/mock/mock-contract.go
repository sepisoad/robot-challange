// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock_warehouse is a generated GoMock package.
package mock_warehouse

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	warehouse "github.com/sepisoad/robot-challange/simulator/warehouse"
)

// MockRobotInterface is a mock of RobotInterface interface.
type MockRobotInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRobotInterfaceMockRecorder
}

// MockRobotInterfaceMockRecorder is the mock recorder for MockRobotInterface.
type MockRobotInterfaceMockRecorder struct {
	mock *MockRobotInterface
}

// NewMockRobotInterface creates a new mock instance.
func NewMockRobotInterface(ctrl *gomock.Controller) *MockRobotInterface {
	mock := &MockRobotInterface{ctrl: ctrl}
	mock.recorder = &MockRobotInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRobotInterface) EXPECT() *MockRobotInterfaceMockRecorder {
	return m.recorder
}

// CancelTask mocks base method.
func (m *MockRobotInterface) CancelTask(taskId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTask", taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTask indicates an expected call of CancelTask.
func (mr *MockRobotInterfaceMockRecorder) CancelTask(taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTask", reflect.TypeOf((*MockRobotInterface)(nil).CancelTask), taskId)
}

// CurrentState mocks base method.
func (m *MockRobotInterface) CurrentState() warehouse.RobotState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentState")
	ret0, _ := ret[0].(warehouse.RobotState)
	return ret0
}

// CurrentState indicates an expected call of CurrentState.
func (mr *MockRobotInterfaceMockRecorder) CurrentState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentState", reflect.TypeOf((*MockRobotInterface)(nil).CurrentState))
}

// EnqueueTask mocks base method.
func (m *MockRobotInterface) EnqueueTask(commands string) (int64, chan warehouse.RobotState, chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueTask", commands)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(chan warehouse.RobotState)
	ret2, _ := ret[2].(chan error)
	return ret0, ret1, ret2
}

// EnqueueTask indicates an expected call of EnqueueTask.
func (mr *MockRobotInterfaceMockRecorder) EnqueueTask(commands interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueTask", reflect.TypeOf((*MockRobotInterface)(nil).EnqueueTask), commands)
}
