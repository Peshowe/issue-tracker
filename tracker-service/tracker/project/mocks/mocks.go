// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Peshowe/issue-tracker/tracker-service/tracker/project (interfaces: ProjectService)

// Package mock_project is a generated GoMock package.
package mock_project

import (
	context "context"
	reflect "reflect"

	project "github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
	gomock "github.com/golang/mock/gomock"
)

// MockProjectService is a mock of ProjectService interface.
type MockProjectService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceMockRecorder
}

// MockProjectServiceMockRecorder is the mock recorder for MockProjectService.
type MockProjectServiceMockRecorder struct {
	mock *MockProjectService
}

// NewMockProjectService creates a new mock instance.
func NewMockProjectService(ctrl *gomock.Controller) *MockProjectService {
	mock := &MockProjectService{ctrl: ctrl}
	mock.recorder = &MockProjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectService) EXPECT() *MockProjectServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockProjectService) AddUser(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockProjectServiceMockRecorder) AddUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockProjectService)(nil).AddUser), arg0, arg1, arg2)
}

// CreateProject mocks base method.
func (m *MockProjectService) CreateProject(arg0 context.Context, arg1 *project.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectServiceMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectService)(nil).CreateProject), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockProjectService) DeleteProject(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectServiceMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectService)(nil).DeleteProject), arg0, arg1)
}

// GetProjectById mocks base method.
func (m *MockProjectService) GetProjectById(arg0 context.Context, arg1 string) (*project.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectById", arg0, arg1)
	ret0, _ := ret[0].(*project.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectById indicates an expected call of GetProjectById.
func (mr *MockProjectServiceMockRecorder) GetProjectById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectById", reflect.TypeOf((*MockProjectService)(nil).GetProjectById), arg0, arg1)
}

// GetProjectsAll mocks base method.
func (m *MockProjectService) GetProjectsAll(arg0 context.Context) ([]*project.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsAll", arg0)
	ret0, _ := ret[0].([]*project.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsAll indicates an expected call of GetProjectsAll.
func (mr *MockProjectServiceMockRecorder) GetProjectsAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsAll", reflect.TypeOf((*MockProjectService)(nil).GetProjectsAll), arg0)
}

// GetProjectsByUser mocks base method.
func (m *MockProjectService) GetProjectsByUser(arg0 context.Context, arg1 string) ([]*project.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsByUser", arg0, arg1)
	ret0, _ := ret[0].([]*project.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsByUser indicates an expected call of GetProjectsByUser.
func (mr *MockProjectServiceMockRecorder) GetProjectsByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsByUser", reflect.TypeOf((*MockProjectService)(nil).GetProjectsByUser), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockProjectService) RemoveUser(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockProjectServiceMockRecorder) RemoveUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockProjectService)(nil).RemoveUser), arg0, arg1, arg2)
}

// UserInProject mocks base method.
func (m *MockProjectService) UserInProject(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserInProject", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UserInProject indicates an expected call of UserInProject.
func (mr *MockProjectServiceMockRecorder) UserInProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInProject", reflect.TypeOf((*MockProjectService)(nil).UserInProject), arg0, arg1)
}
