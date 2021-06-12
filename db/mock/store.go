// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MadhavaAdiga/grpc-hrm-server/db (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/MadhavaAdiga/grpc-hrm-server/db"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateOrganization mocks base method.
func (m *MockStore) CreateOrganization(arg0 context.Context, arg1 db.CreateOrganizationParam) (db.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0, arg1)
	ret0, _ := ret[0].(db.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization.
func (mr *MockStoreMockRecorder) CreateOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockStore)(nil).CreateOrganization), arg0, arg1)
}

// CreateRole mocks base method.
func (m *MockStore) CreateRole(arg0 context.Context, arg1 db.CreateRoleParam) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockStoreMockRecorder) CreateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockStore)(nil).CreateRole), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParam) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// FindOrganizationByName mocks base method.
func (m *MockStore) FindOrganizationByName(arg0 context.Context, arg1 string) (db.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrganizationByName", arg0, arg1)
	ret0, _ := ret[0].(db.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationByName indicates an expected call of FindOrganizationByName.
func (mr *MockStoreMockRecorder) FindOrganizationByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationByName", reflect.TypeOf((*MockStore)(nil).FindOrganizationByName), arg0, arg1)
}

// FindRoleByOrganization mocks base method.
func (m *MockStore) FindRoleByOrganization(arg0 context.Context, arg1 db.FindRoleByOrgParam) (db.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoleByOrganization", arg0, arg1)
	ret0, _ := ret[0].(db.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoleByOrganization indicates an expected call of FindRoleByOrganization.
func (mr *MockStoreMockRecorder) FindRoleByOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoleByOrganization", reflect.TypeOf((*MockStore)(nil).FindRoleByOrganization), arg0, arg1)
}

// FindUserByName mocks base method.
func (m *MockStore) FindUserByName(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName.
func (mr *MockStoreMockRecorder) FindUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockStore)(nil).FindUserByName), arg0, arg1)
}
