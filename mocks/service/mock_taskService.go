// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/emrekursatt/JuniorProject/services (interfaces: TaskService)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dto "github.com/emrekursatt/JuniorProject/dto"
	models "github.com/emrekursatt/JuniorProject/models"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTaskService) Delete(arg0 models.TaskEntity) (*dto.TaskDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*dto.TaskDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskService)(nil).Delete), arg0)
}

// GetAllTasks mocks base method.
func (m *MockTaskService) GetAllTasks() ([]models.TaskEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks")
	ret0, _ := ret[0].([]models.TaskEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockTaskServiceMockRecorder) GetAllTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockTaskService)(nil).GetAllTasks))
}

// Insert mocks base method.
func (m *MockTaskService) Insert(arg0 models.TaskEntity) (*dto.TaskDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(*dto.TaskDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTaskServiceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTaskService)(nil).Insert), arg0)
}

// Update mocks base method.
func (m *MockTaskService) Update(arg0 models.TaskEntity) (*dto.TaskDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*dto.TaskDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskService)(nil).Update), arg0)
}