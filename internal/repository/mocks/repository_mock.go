// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	entity "diploma/internal/entity/v1"
	reflect "reflect"
	time "time"

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
func (m *MockAuthorization) CreateUser(user entity.User) (int, error) {
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

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(name, password string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", name, password)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), name, password)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeleteAllSiteTexts mocks base method.
func (m *MockService) DeleteAllSiteTexts(ctx context.Context, userID, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSiteTexts", ctx, userID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllSiteTexts indicates an expected call of DeleteAllSiteTexts.
func (mr *MockServiceMockRecorder) DeleteAllSiteTexts(ctx, userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSiteTexts", reflect.TypeOf((*MockService)(nil).DeleteAllSiteTexts), ctx, userID, id)
}

// DeleteAllSites mocks base method.
func (m *MockService) DeleteAllSites(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSites", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllSites indicates an expected call of DeleteAllSites.
func (mr *MockServiceMockRecorder) DeleteAllSites(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSites", reflect.TypeOf((*MockService)(nil).DeleteAllSites), ctx, userID)
}

// DeleteAllTexts mocks base method.
func (m *MockService) DeleteAllTexts(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllTexts", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllTexts indicates an expected call of DeleteAllTexts.
func (mr *MockServiceMockRecorder) DeleteAllTexts(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllTexts", reflect.TypeOf((*MockService)(nil).DeleteAllTexts), ctx, userID)
}

// DeleteSite mocks base method.
func (m *MockService) DeleteSite(ctx context.Context, userID, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSite", ctx, userID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSite indicates an expected call of DeleteSite.
func (mr *MockServiceMockRecorder) DeleteSite(ctx, userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSite", reflect.TypeOf((*MockService)(nil).DeleteSite), ctx, userID, id)
}

// DeleteUsers mocks base method.
func (m *MockService) DeleteUsers(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers.
func (mr *MockServiceMockRecorder) DeleteUsers(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockService)(nil).DeleteUsers), ctx, id)
}

// GetListSites mocks base method.
func (m *MockService) GetListSites(ctx context.Context, userID int) ([]entity.Sites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListSites", ctx, userID)
	ret0, _ := ret[0].([]entity.Sites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListSites indicates an expected call of GetListSites.
func (mr *MockServiceMockRecorder) GetListSites(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListSites", reflect.TypeOf((*MockService)(nil).GetListSites), ctx, userID)
}

// GetMainText mocks base method.
func (m *MockService) GetMainText(ctx context.Context, userID, id, siteID int) (entity.MainText, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainText", ctx, userID, id, siteID)
	ret0, _ := ret[0].(entity.MainText)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMainText indicates an expected call of GetMainText.
func (mr *MockServiceMockRecorder) GetMainText(ctx, userID, id, siteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainText", reflect.TypeOf((*MockService)(nil).GetMainText), ctx, userID, id, siteID)
}

// GetSite mocks base method.
func (m *MockService) GetSite(ctx context.Context, userID, id int) (entity.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSite", ctx, userID, id)
	ret0, _ := ret[0].(entity.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSite indicates an expected call of GetSite.
func (mr *MockServiceMockRecorder) GetSite(ctx, userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSite", reflect.TypeOf((*MockService)(nil).GetSite), ctx, userID, id)
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, id int) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, id)
}

// GetUsers mocks base method.
func (m *MockService) GetUsers(ctx context.Context, id int) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockServiceMockRecorder) GetUsers(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockService)(nil).GetUsers), ctx, id)
}

// ParseSite mocks base method.
func (m *MockService) ParseSite(ctx context.Context, userID int, url, tag string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseSite", ctx, userID, url, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseSite indicates an expected call of ParseSite.
func (mr *MockServiceMockRecorder) ParseSite(ctx, userID, url, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseSite", reflect.TypeOf((*MockService)(nil).ParseSite), ctx, userID, url, tag)
}

// PostMainText mocks base method.
func (m *MockService) PostMainText(ctx context.Context, userID, siteID int, date time.Time, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMainText", ctx, userID, siteID, date, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMainText indicates an expected call of PostMainText.
func (mr *MockServiceMockRecorder) PostMainText(ctx, userID, siteID, date, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMainText", reflect.TypeOf((*MockService)(nil).PostMainText), ctx, userID, siteID, date, text)
}

// PostSite mocks base method.
func (m *MockService) PostSite(ctx context.Context, userID int, url, tag string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostSite", ctx, userID, url, tag)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSite indicates an expected call of PostSite.
func (mr *MockServiceMockRecorder) PostSite(ctx, userID, url, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSite", reflect.TypeOf((*MockService)(nil).PostSite), ctx, userID, url, tag)
}

// UpdateMainText mocks base method.
func (m *MockService) UpdateMainText(ctx context.Context, userID, id, siteID int, date time.Time, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMainText", ctx, userID, id, siteID, date, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMainText indicates an expected call of UpdateMainText.
func (mr *MockServiceMockRecorder) UpdateMainText(ctx, userID, id, siteID, date, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMainText", reflect.TypeOf((*MockService)(nil).UpdateMainText), ctx, userID, id, siteID, date, text)
}

// UpdateUsers mocks base method.
func (m *MockService) UpdateUsers(ctx context.Context, id int, login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsers", ctx, id, login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUsers indicates an expected call of UpdateUsers.
func (mr *MockServiceMockRecorder) UpdateUsers(ctx, id, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsers", reflect.TypeOf((*MockService)(nil).UpdateUsers), ctx, id, login, password)
}
