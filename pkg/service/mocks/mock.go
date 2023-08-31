// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	enty "github.com/pasha1coil/testingavito/pkg/service/system"
)

// MockCommands is a mock of Commands interface.
type MockCommands struct {
	ctrl     *gomock.Controller
	recorder *MockCommandsMockRecorder
}

// MockCommandsMockRecorder is the mock recorder for MockCommands.
type MockCommandsMockRecorder struct {
	mock *MockCommands
}

// NewMockCommands creates a new mock instance.
func NewMockCommands(ctrl *gomock.Controller) *MockCommands {
	mock := &MockCommands{ctrl: ctrl}
	mock.recorder = &MockCommandsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommands) EXPECT() *MockCommandsMockRecorder {
	return m.recorder
}

// CreateSegment mocks base method.
func (m *MockCommands) CreateSegment(segment enty.Segment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSegment", segment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSegment indicates an expected call of CreateSegment.
func (mr *MockCommandsMockRecorder) CreateSegment(segment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSegment", reflect.TypeOf((*MockCommands)(nil).CreateSegment), segment)
}

// CreateUser mocks base method.
func (m *MockCommands) CreateUser(user enty.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockCommandsMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockCommands)(nil).CreateUser), user)
}

// DelSegment mocks base method.
func (m *MockCommands) DelSegment(segment enty.Segment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelSegment", segment)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelSegment indicates an expected call of DelSegment.
func (mr *MockCommandsMockRecorder) DelSegment(segment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelSegment", reflect.TypeOf((*MockCommands)(nil).DelSegment), segment)
}

// DeleteSemUser mocks base method.
func (m *MockCommands) DeleteSemUser(NameSegment []string, UserID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSemUser", NameSegment, UserID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSemUser indicates an expected call of DeleteSemUser.
func (mr *MockCommandsMockRecorder) DeleteSemUser(NameSegment, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSemUser", reflect.TypeOf((*MockCommands)(nil).DeleteSemUser), NameSegment, UserID)
}

// GetActiveSlugs mocks base method.
func (m *MockCommands) GetActiveSlugs(user enty.User) ([]enty.ListNames, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveSlugs", user)
	ret0, _ := ret[0].([]enty.ListNames)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveSlugs indicates an expected call of GetActiveSlugs.
func (mr *MockCommandsMockRecorder) GetActiveSlugs(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveSlugs", reflect.TypeOf((*MockCommands)(nil).GetActiveSlugs), user)
}

// GetCsvHistory mocks base method.
func (m *MockCommands) GetCsvHistory(userId int, startDate, endDate string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCsvHistory", userId, startDate, endDate)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCsvHistory indicates an expected call of GetCsvHistory.
func (mr *MockCommandsMockRecorder) GetCsvHistory(userId, startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCsvHistory", reflect.TypeOf((*MockCommands)(nil).GetCsvHistory), userId, startDate, endDate)
}

// GetSlugHistory mocks base method.
func (m *MockCommands) GetSlugHistory(userId int, startDate, endDate string) ([]enty.History, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlugHistory", userId, startDate, endDate)
	ret0, _ := ret[0].([]enty.History)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlugHistory indicates an expected call of GetSlugHistory.
func (mr *MockCommandsMockRecorder) GetSlugHistory(userId, startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlugHistory", reflect.TypeOf((*MockCommands)(nil).GetSlugHistory), userId, startDate, endDate)
}

// InsertSemUser mocks base method.
func (m *MockCommands) InsertSemUser(NameSegment []string, UserID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSemUser", NameSegment, UserID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertSemUser indicates an expected call of InsertSemUser.
func (mr *MockCommandsMockRecorder) InsertSemUser(NameSegment, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSemUser", reflect.TypeOf((*MockCommands)(nil).InsertSemUser), NameSegment, UserID)
}
