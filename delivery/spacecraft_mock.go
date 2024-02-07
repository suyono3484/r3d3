// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/suyono3484/r3d3/delivery (interfaces: UseCase,FilterHelper)
//
// Generated by this command:
//
//	mockgen -package delivery -destination spacecraft_mock.go . UseCase,FilterHelper
//

// Package delivery is a generated GoMock package.
package delivery

import (
	reflect "reflect"

	r3d3 "github.com/suyono3484/r3d3"
	gomock "go.uber.org/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(arg0 r3d3.SpaceCraftCreate, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockUseCase) Get(arg0 int64) (r3d3.SpaceCraft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(r3d3.SpaceCraft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUseCaseMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUseCase)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockUseCase) List(arg0 ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]r3d3.SpaceCraftInList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUseCaseMockRecorder) List(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUseCase)(nil).List), arg0...)
}

// Update mocks base method.
func (m *MockUseCase) Update(arg0 int64, arg1 r3d3.SpaceCraftCreate, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), arg0, arg1, arg2)
}

// MockFilterHelper is a mock of FilterHelper interface.
type MockFilterHelper struct {
	ctrl     *gomock.Controller
	recorder *MockFilterHelperMockRecorder
}

// MockFilterHelperMockRecorder is the mock recorder for MockFilterHelper.
type MockFilterHelperMockRecorder struct {
	mock *MockFilterHelper
}

// NewMockFilterHelper creates a new mock instance.
func NewMockFilterHelper(ctrl *gomock.Controller) *MockFilterHelper {
	mock := &MockFilterHelper{ctrl: ctrl}
	mock.recorder = &MockFilterHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilterHelper) EXPECT() *MockFilterHelperMockRecorder {
	return m.recorder
}

// NewClassFilter mocks base method.
func (m *MockFilterHelper) NewClassFilter(arg0 string) r3d3.ListFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClassFilter", arg0)
	ret0, _ := ret[0].(r3d3.ListFilter)
	return ret0
}

// NewClassFilter indicates an expected call of NewClassFilter.
func (mr *MockFilterHelperMockRecorder) NewClassFilter(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClassFilter", reflect.TypeOf((*MockFilterHelper)(nil).NewClassFilter), arg0)
}

// NewNameFilter mocks base method.
func (m *MockFilterHelper) NewNameFilter(arg0 string) r3d3.ListFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNameFilter", arg0)
	ret0, _ := ret[0].(r3d3.ListFilter)
	return ret0
}

// NewNameFilter indicates an expected call of NewNameFilter.
func (mr *MockFilterHelperMockRecorder) NewNameFilter(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNameFilter", reflect.TypeOf((*MockFilterHelper)(nil).NewNameFilter), arg0)
}

// NewStatusFilter mocks base method.
func (m *MockFilterHelper) NewStatusFilter(arg0 string) r3d3.ListFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStatusFilter", arg0)
	ret0, _ := ret[0].(r3d3.ListFilter)
	return ret0
}

// NewStatusFilter indicates an expected call of NewStatusFilter.
func (mr *MockFilterHelperMockRecorder) NewStatusFilter(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStatusFilter", reflect.TypeOf((*MockFilterHelper)(nil).NewStatusFilter), arg0)
}