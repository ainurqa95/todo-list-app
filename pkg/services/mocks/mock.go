// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_services is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	todo "github.com/ainurqa95/todo-list-app"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user todo.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}

// MockItemManager is a mock of ItemManager interface.
type MockItemManager struct {
	ctrl     *gomock.Controller
	recorder *MockItemManagerMockRecorder
}

// MockItemManagerMockRecorder is the mock recorder for MockItemManager.
type MockItemManagerMockRecorder struct {
	mock *MockItemManager
}

// NewMockItemManager creates a new mock instance.
func NewMockItemManager(ctrl *gomock.Controller) *MockItemManager {
	mock := &MockItemManager{ctrl: ctrl}
	mock.recorder = &MockItemManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemManager) EXPECT() *MockItemManagerMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockItemManager) CreateItem(item todo.TodoItem, listId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", item, listId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockItemManagerMockRecorder) CreateItem(item, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemManager)(nil).CreateItem), item, listId)
}

// GetAllItems mocks base method.
func (m *MockItemManager) GetAllItems(listId, userId int) ([]todo.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllItems", listId, userId)
	ret0, _ := ret[0].([]todo.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockItemManagerMockRecorder) GetAllItems(listId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockItemManager)(nil).GetAllItems), listId, userId)
}

// MockListManager is a mock of ListManager interface.
type MockListManager struct {
	ctrl     *gomock.Controller
	recorder *MockListManagerMockRecorder
}

// MockListManagerMockRecorder is the mock recorder for MockListManager.
type MockListManagerMockRecorder struct {
	mock *MockListManager
}

// NewMockListManager creates a new mock instance.
func NewMockListManager(ctrl *gomock.Controller) *MockListManager {
	mock := &MockListManager{ctrl: ctrl}
	mock.recorder = &MockListManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListManager) EXPECT() *MockListManagerMockRecorder {
	return m.recorder
}

// CreateList mocks base method.
func (m *MockListManager) CreateList(list todo.TodoList, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", list, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockListManagerMockRecorder) CreateList(list, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockListManager)(nil).CreateList), list, userId)
}

// DeleteList mocks base method.
func (m *MockListManager) DeleteList(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockListManagerMockRecorder) DeleteList(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockListManager)(nil).DeleteList), id)
}

// GetAllLists mocks base method.
func (m *MockListManager) GetAllLists(userId int) ([]todo.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLists", userId)
	ret0, _ := ret[0].([]todo.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLists indicates an expected call of GetAllLists.
func (mr *MockListManagerMockRecorder) GetAllLists(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLists", reflect.TypeOf((*MockListManager)(nil).GetAllLists), userId)
}

// GetListById mocks base method.
func (m *MockListManager) GetListById(id int) (todo.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListById", id)
	ret0, _ := ret[0].(todo.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListById indicates an expected call of GetListById.
func (mr *MockListManagerMockRecorder) GetListById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListById", reflect.TypeOf((*MockListManager)(nil).GetListById), id)
}

// UpdateList mocks base method.
func (m *MockListManager) UpdateList(id int, input todo.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockListManagerMockRecorder) UpdateList(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockListManager)(nil).UpdateList), id, input)
}