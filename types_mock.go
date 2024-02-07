// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/suyono3484/r3d3 (interfaces: Armament,SpaceCraft)
//
// Generated by this command:
//
//	mockgen -package r3d3 -destination types_mock.go . Armament,SpaceCraft
//

// Package r3d3 is a generated GoMock package.
package r3d3

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockArmament is a mock of Armament interface.
type MockArmament struct {
	ctrl     *gomock.Controller
	recorder *MockArmamentMockRecorder
}

// MockArmamentMockRecorder is the mock recorder for MockArmament.
type MockArmamentMockRecorder struct {
	mock *MockArmament
}

// NewMockArmament creates a new mock instance.
func NewMockArmament(ctrl *gomock.Controller) *MockArmament {
	mock := &MockArmament{ctrl: ctrl}
	mock.recorder = &MockArmamentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArmament) EXPECT() *MockArmamentMockRecorder {
	return m.recorder
}

// Qty mocks base method.
func (m *MockArmament) Qty() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Qty")
	ret0, _ := ret[0].(int)
	return ret0
}

// Qty indicates an expected call of Qty.
func (mr *MockArmamentMockRecorder) Qty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qty", reflect.TypeOf((*MockArmament)(nil).Qty))
}

// Title mocks base method.
func (m *MockArmament) Title() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Title")
	ret0, _ := ret[0].(string)
	return ret0
}

// Title indicates an expected call of Title.
func (mr *MockArmamentMockRecorder) Title() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Title", reflect.TypeOf((*MockArmament)(nil).Title))
}

// MockSpaceCraft is a mock of SpaceCraft interface.
type MockSpaceCraft struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceCraftMockRecorder
}

// MockSpaceCraftMockRecorder is the mock recorder for MockSpaceCraft.
type MockSpaceCraftMockRecorder struct {
	mock *MockSpaceCraft
}

// NewMockSpaceCraft creates a new mock instance.
func NewMockSpaceCraft(ctrl *gomock.Controller) *MockSpaceCraft {
	mock := &MockSpaceCraft{ctrl: ctrl}
	mock.recorder = &MockSpaceCraftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceCraft) EXPECT() *MockSpaceCraftMockRecorder {
	return m.recorder
}

// Armament mocks base method.
func (m *MockSpaceCraft) Armament() []Armament {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Armament")
	ret0, _ := ret[0].([]Armament)
	return ret0
}

// Armament indicates an expected call of Armament.
func (mr *MockSpaceCraftMockRecorder) Armament() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Armament", reflect.TypeOf((*MockSpaceCraft)(nil).Armament))
}

// Class mocks base method.
func (m *MockSpaceCraft) Class() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Class")
	ret0, _ := ret[0].(string)
	return ret0
}

// Class indicates an expected call of Class.
func (mr *MockSpaceCraftMockRecorder) Class() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Class", reflect.TypeOf((*MockSpaceCraft)(nil).Class))
}

// Crew mocks base method.
func (m *MockSpaceCraft) Crew() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Crew")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Crew indicates an expected call of Crew.
func (mr *MockSpaceCraftMockRecorder) Crew() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Crew", reflect.TypeOf((*MockSpaceCraft)(nil).Crew))
}

// ID mocks base method.
func (m *MockSpaceCraft) ID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSpaceCraftMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSpaceCraft)(nil).ID))
}

// ImageURL mocks base method.
func (m *MockSpaceCraft) ImageURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// ImageURL indicates an expected call of ImageURL.
func (mr *MockSpaceCraftMockRecorder) ImageURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageURL", reflect.TypeOf((*MockSpaceCraft)(nil).ImageURL))
}

// Name mocks base method.
func (m *MockSpaceCraft) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSpaceCraftMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSpaceCraft)(nil).Name))
}

// Status mocks base method.
func (m *MockSpaceCraft) Status() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(string)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockSpaceCraftMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSpaceCraft)(nil).Status))
}

// Value mocks base method.
func (m *MockSpaceCraft) Value() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockSpaceCraftMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockSpaceCraft)(nil).Value))
}