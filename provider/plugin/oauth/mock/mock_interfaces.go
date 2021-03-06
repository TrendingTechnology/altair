// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	entity "github.com/codefluence-x/altair/provider/plugin/oauth/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOauthApplicationModel is a mock of OauthApplicationModel interface
type MockOauthApplicationModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationModelMockRecorder
}

// MockOauthApplicationModelMockRecorder is the mock recorder for MockOauthApplicationModel
type MockOauthApplicationModelMockRecorder struct {
	mock *MockOauthApplicationModel
}

// NewMockOauthApplicationModel creates a new mock instance
func NewMockOauthApplicationModel(ctrl *gomock.Controller) *MockOauthApplicationModel {
	mock := &MockOauthApplicationModel{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationModel) EXPECT() *MockOauthApplicationModelMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockOauthApplicationModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOauthApplicationModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOauthApplicationModel)(nil).Name))
}

// Paginate mocks base method
func (m *MockOauthApplicationModel) Paginate(ctx context.Context, offset, limit int) ([]entity.OauthApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paginate", ctx, offset, limit)
	ret0, _ := ret[0].([]entity.OauthApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockOauthApplicationModelMockRecorder) Paginate(ctx, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockOauthApplicationModel)(nil).Paginate), ctx, offset, limit)
}

// One mocks base method
func (m *MockOauthApplicationModel) One(ctx context.Context, ID int) (entity.OauthApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, ID)
	ret0, _ := ret[0].(entity.OauthApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthApplicationModelMockRecorder) One(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthApplicationModel)(nil).One), ctx, ID)
}

// OneByUIDandSecret mocks base method
func (m *MockOauthApplicationModel) OneByUIDandSecret(ctx context.Context, clientUID, clientSecret string) (entity.OauthApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByUIDandSecret", ctx, clientUID, clientSecret)
	ret0, _ := ret[0].(entity.OauthApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OneByUIDandSecret indicates an expected call of OneByUIDandSecret
func (mr *MockOauthApplicationModelMockRecorder) OneByUIDandSecret(ctx, clientUID, clientSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByUIDandSecret", reflect.TypeOf((*MockOauthApplicationModel)(nil).OneByUIDandSecret), ctx, clientUID, clientSecret)
}

// Count mocks base method
func (m *MockOauthApplicationModel) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockOauthApplicationModelMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOauthApplicationModel)(nil).Count), ctx)
}

// Create mocks base method
func (m *MockOauthApplicationModel) Create(ctx context.Context, data entity.OauthApplicationInsertable, txs ...*sql.Tx) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, data}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthApplicationModelMockRecorder) Create(ctx, data interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, data}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthApplicationModel)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockOauthApplicationModel) Update(ctx context.Context, ID int, data entity.OauthApplicationUpdateable, txs ...*sql.Tx) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ID, data}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockOauthApplicationModelMockRecorder) Update(ctx, ID, data interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ID, data}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOauthApplicationModel)(nil).Update), varargs...)
}

// MockOauthAccessTokenModel is a mock of OauthAccessTokenModel interface
type MockOauthAccessTokenModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessTokenModelMockRecorder
}

// MockOauthAccessTokenModelMockRecorder is the mock recorder for MockOauthAccessTokenModel
type MockOauthAccessTokenModelMockRecorder struct {
	mock *MockOauthAccessTokenModel
}

// NewMockOauthAccessTokenModel creates a new mock instance
func NewMockOauthAccessTokenModel(ctrl *gomock.Controller) *MockOauthAccessTokenModel {
	mock := &MockOauthAccessTokenModel{ctrl: ctrl}
	mock.recorder = &MockOauthAccessTokenModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthAccessTokenModel) EXPECT() *MockOauthAccessTokenModelMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockOauthAccessTokenModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOauthAccessTokenModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).Name))
}

// One mocks base method
func (m *MockOauthAccessTokenModel) One(ctx context.Context, ID int) (entity.OauthAccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, ID)
	ret0, _ := ret[0].(entity.OauthAccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthAccessTokenModelMockRecorder) One(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).One), ctx, ID)
}

// OneByToken mocks base method
func (m *MockOauthAccessTokenModel) OneByToken(ctx context.Context, token string) (entity.OauthAccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByToken", ctx, token)
	ret0, _ := ret[0].(entity.OauthAccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OneByToken indicates an expected call of OneByToken
func (mr *MockOauthAccessTokenModelMockRecorder) OneByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByToken", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).OneByToken), ctx, token)
}

// Create mocks base method
func (m *MockOauthAccessTokenModel) Create(ctx context.Context, data entity.OauthAccessTokenInsertable, txs ...*sql.Tx) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, data}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthAccessTokenModelMockRecorder) Create(ctx, data interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, data}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).Create), varargs...)
}

// Revoke mocks base method
func (m *MockOauthAccessTokenModel) Revoke(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke
func (mr *MockOauthAccessTokenModelMockRecorder) Revoke(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).Revoke), ctx, token)
}

// MockOauthAccessGrantModel is a mock of OauthAccessGrantModel interface
type MockOauthAccessGrantModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessGrantModelMockRecorder
}

// MockOauthAccessGrantModelMockRecorder is the mock recorder for MockOauthAccessGrantModel
type MockOauthAccessGrantModelMockRecorder struct {
	mock *MockOauthAccessGrantModel
}

// NewMockOauthAccessGrantModel creates a new mock instance
func NewMockOauthAccessGrantModel(ctrl *gomock.Controller) *MockOauthAccessGrantModel {
	mock := &MockOauthAccessGrantModel{ctrl: ctrl}
	mock.recorder = &MockOauthAccessGrantModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthAccessGrantModel) EXPECT() *MockOauthAccessGrantModelMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockOauthAccessGrantModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOauthAccessGrantModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).Name))
}

// One mocks base method
func (m *MockOauthAccessGrantModel) One(ctx context.Context, ID int) (entity.OauthAccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, ID)
	ret0, _ := ret[0].(entity.OauthAccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthAccessGrantModelMockRecorder) One(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).One), ctx, ID)
}

// Create mocks base method
func (m *MockOauthAccessGrantModel) Create(ctx context.Context, data entity.OauthAccessGrantInsertable, txs ...*sql.Tx) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, data}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthAccessGrantModelMockRecorder) Create(ctx, data interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, data}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).Create), varargs...)
}

// OneByCode mocks base method
func (m *MockOauthAccessGrantModel) OneByCode(ctx context.Context, code string) (entity.OauthAccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByCode", ctx, code)
	ret0, _ := ret[0].(entity.OauthAccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OneByCode indicates an expected call of OneByCode
func (mr *MockOauthAccessGrantModelMockRecorder) OneByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByCode", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).OneByCode), ctx, code)
}

// Revoke mocks base method
func (m *MockOauthAccessGrantModel) Revoke(ctx context.Context, code string, txs ...*sql.Tx) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, code}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Revoke", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke
func (mr *MockOauthAccessGrantModelMockRecorder) Revoke(ctx, code interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, code}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).Revoke), varargs...)
}

// MockApplicationManager is a mock of ApplicationManager interface
type MockApplicationManager struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationManagerMockRecorder
}

// MockApplicationManagerMockRecorder is the mock recorder for MockApplicationManager
type MockApplicationManagerMockRecorder struct {
	mock *MockApplicationManager
}

// NewMockApplicationManager creates a new mock instance
func NewMockApplicationManager(ctrl *gomock.Controller) *MockApplicationManager {
	mock := &MockApplicationManager{ctrl: ctrl}
	mock.recorder = &MockApplicationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationManager) EXPECT() *MockApplicationManagerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockApplicationManager) List(ctx context.Context, offset, limit int) ([]entity.OauthApplicationJSON, int, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, offset, limit)
	ret0, _ := ret[0].([]entity.OauthApplicationJSON)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(*entity.Error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockApplicationManagerMockRecorder) List(ctx, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationManager)(nil).List), ctx, offset, limit)
}

// One mocks base method
func (m *MockApplicationManager) One(ctx context.Context, ID int) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, ID)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockApplicationManagerMockRecorder) One(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockApplicationManager)(nil).One), ctx, ID)
}

// Create mocks base method
func (m *MockApplicationManager) Create(ctx context.Context, e entity.OauthApplicationJSON) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, e)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockApplicationManagerMockRecorder) Create(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApplicationManager)(nil).Create), ctx, e)
}

// Update mocks base method
func (m *MockApplicationManager) Update(ctx context.Context, ID int, e entity.OauthApplicationUpdateJSON) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ID, e)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockApplicationManagerMockRecorder) Update(ctx, ID, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockApplicationManager)(nil).Update), ctx, ID, e)
}

// MockAuthorization is a mock of Authorization interface
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// Grantor mocks base method
func (m *MockAuthorization) Grantor(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (interface{}, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grantor", ctx, authorizationReq)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Grantor indicates an expected call of Grantor
func (mr *MockAuthorizationMockRecorder) Grantor(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grantor", reflect.TypeOf((*MockAuthorization)(nil).Grantor), ctx, authorizationReq)
}

// Grant mocks base method
func (m *MockAuthorization) Grant(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessGrantJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grant", ctx, authorizationReq)
	ret0, _ := ret[0].(entity.OauthAccessGrantJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Grant indicates an expected call of Grant
func (mr *MockAuthorizationMockRecorder) Grant(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grant", reflect.TypeOf((*MockAuthorization)(nil).Grant), ctx, authorizationReq)
}

// GrantToken mocks base method
func (m *MockAuthorization) GrantToken(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantToken", ctx, authorizationReq)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GrantToken indicates an expected call of GrantToken
func (mr *MockAuthorizationMockRecorder) GrantToken(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantToken", reflect.TypeOf((*MockAuthorization)(nil).GrantToken), ctx, authorizationReq)
}

// Token mocks base method
func (m *MockAuthorization) Token(ctx context.Context, accessTokenReq entity.AccessTokenRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", ctx, accessTokenReq)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockAuthorizationMockRecorder) Token(ctx, accessTokenReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockAuthorization)(nil).Token), ctx, accessTokenReq)
}

// RevokeToken mocks base method
func (m *MockAuthorization) RevokeToken(ctx context.Context, revokeAccessTokenReq entity.RevokeAccessTokenRequestJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", ctx, revokeAccessTokenReq)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// RevokeToken indicates an expected call of RevokeToken
func (mr *MockAuthorizationMockRecorder) RevokeToken(ctx, revokeAccessTokenReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockAuthorization)(nil).RevokeToken), ctx, revokeAccessTokenReq)
}

// MockOauthApplicationFormater is a mock of OauthApplicationFormater interface
type MockOauthApplicationFormater struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationFormaterMockRecorder
}

// MockOauthApplicationFormaterMockRecorder is the mock recorder for MockOauthApplicationFormater
type MockOauthApplicationFormaterMockRecorder struct {
	mock *MockOauthApplicationFormater
}

// NewMockOauthApplicationFormater creates a new mock instance
func NewMockOauthApplicationFormater(ctrl *gomock.Controller) *MockOauthApplicationFormater {
	mock := &MockOauthApplicationFormater{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationFormaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationFormater) EXPECT() *MockOauthApplicationFormaterMockRecorder {
	return m.recorder
}

// ApplicationList mocks base method
func (m *MockOauthApplicationFormater) ApplicationList(ctx context.Context, applications []entity.OauthApplication) []entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationList", ctx, applications)
	ret0, _ := ret[0].([]entity.OauthApplicationJSON)
	return ret0
}

// ApplicationList indicates an expected call of ApplicationList
func (mr *MockOauthApplicationFormaterMockRecorder) ApplicationList(ctx, applications interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationList", reflect.TypeOf((*MockOauthApplicationFormater)(nil).ApplicationList), ctx, applications)
}

// Application mocks base method
func (m *MockOauthApplicationFormater) Application(ctx context.Context, application entity.OauthApplication) entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", ctx, application)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockOauthApplicationFormaterMockRecorder) Application(ctx, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockOauthApplicationFormater)(nil).Application), ctx, application)
}

// MockOauthFormatter is a mock of OauthFormatter interface
type MockOauthFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthFormatterMockRecorder
}

// MockOauthFormatterMockRecorder is the mock recorder for MockOauthFormatter
type MockOauthFormatterMockRecorder struct {
	mock *MockOauthFormatter
}

// NewMockOauthFormatter creates a new mock instance
func NewMockOauthFormatter(ctrl *gomock.Controller) *MockOauthFormatter {
	mock := &MockOauthFormatter{ctrl: ctrl}
	mock.recorder = &MockOauthFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthFormatter) EXPECT() *MockOauthFormatterMockRecorder {
	return m.recorder
}

// AccessGrant mocks base method
func (m *MockOauthFormatter) AccessGrant(e entity.OauthAccessGrant) entity.OauthAccessGrantJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrant", e)
	ret0, _ := ret[0].(entity.OauthAccessGrantJSON)
	return ret0
}

// AccessGrant indicates an expected call of AccessGrant
func (mr *MockOauthFormatterMockRecorder) AccessGrant(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrant", reflect.TypeOf((*MockOauthFormatter)(nil).AccessGrant), e)
}

// AccessToken mocks base method
func (m *MockOauthFormatter) AccessToken(e entity.OauthAccessToken, redirectURI string) entity.OauthAccessTokenJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken", e, redirectURI)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	return ret0
}

// AccessToken indicates an expected call of AccessToken
func (mr *MockOauthFormatterMockRecorder) AccessToken(e, redirectURI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockOauthFormatter)(nil).AccessToken), e, redirectURI)
}

// MockModelFormater is a mock of ModelFormater interface
type MockModelFormater struct {
	ctrl     *gomock.Controller
	recorder *MockModelFormaterMockRecorder
}

// MockModelFormaterMockRecorder is the mock recorder for MockModelFormater
type MockModelFormaterMockRecorder struct {
	mock *MockModelFormater
}

// NewMockModelFormater creates a new mock instance
func NewMockModelFormater(ctrl *gomock.Controller) *MockModelFormater {
	mock := &MockModelFormater{ctrl: ctrl}
	mock.recorder = &MockModelFormaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelFormater) EXPECT() *MockModelFormaterMockRecorder {
	return m.recorder
}

// AccessTokenFromAuthorizationRequest mocks base method
func (m *MockModelFormater) AccessTokenFromAuthorizationRequest(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenFromAuthorizationRequest", r, application)
	ret0, _ := ret[0].(entity.OauthAccessTokenInsertable)
	return ret0
}

// AccessTokenFromAuthorizationRequest indicates an expected call of AccessTokenFromAuthorizationRequest
func (mr *MockModelFormaterMockRecorder) AccessTokenFromAuthorizationRequest(r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenFromAuthorizationRequest", reflect.TypeOf((*MockModelFormater)(nil).AccessTokenFromAuthorizationRequest), r, application)
}

// AccessTokenFromOauthAccessGrant mocks base method
func (m *MockModelFormater) AccessTokenFromOauthAccessGrant(oauthAccessGrant entity.OauthAccessGrant, application entity.OauthApplication) entity.OauthAccessTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenFromOauthAccessGrant", oauthAccessGrant, application)
	ret0, _ := ret[0].(entity.OauthAccessTokenInsertable)
	return ret0
}

// AccessTokenFromOauthAccessGrant indicates an expected call of AccessTokenFromOauthAccessGrant
func (mr *MockModelFormaterMockRecorder) AccessTokenFromOauthAccessGrant(oauthAccessGrant, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenFromOauthAccessGrant", reflect.TypeOf((*MockModelFormater)(nil).AccessTokenFromOauthAccessGrant), oauthAccessGrant, application)
}

// AccessGrantFromAuthorizationRequest mocks base method
func (m *MockModelFormater) AccessGrantFromAuthorizationRequest(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessGrantInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrantFromAuthorizationRequest", r, application)
	ret0, _ := ret[0].(entity.OauthAccessGrantInsertable)
	return ret0
}

// AccessGrantFromAuthorizationRequest indicates an expected call of AccessGrantFromAuthorizationRequest
func (mr *MockModelFormaterMockRecorder) AccessGrantFromAuthorizationRequest(r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrantFromAuthorizationRequest", reflect.TypeOf((*MockModelFormater)(nil).AccessGrantFromAuthorizationRequest), r, application)
}

// OauthApplication mocks base method
func (m *MockModelFormater) OauthApplication(r entity.OauthApplicationJSON) entity.OauthApplicationInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OauthApplication", r)
	ret0, _ := ret[0].(entity.OauthApplicationInsertable)
	return ret0
}

// OauthApplication indicates an expected call of OauthApplication
func (mr *MockModelFormaterMockRecorder) OauthApplication(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OauthApplication", reflect.TypeOf((*MockModelFormater)(nil).OauthApplication), r)
}

// MockOauthValidator is a mock of OauthValidator interface
type MockOauthValidator struct {
	ctrl     *gomock.Controller
	recorder *MockOauthValidatorMockRecorder
}

// MockOauthValidatorMockRecorder is the mock recorder for MockOauthValidator
type MockOauthValidatorMockRecorder struct {
	mock *MockOauthValidator
}

// NewMockOauthValidator creates a new mock instance
func NewMockOauthValidator(ctrl *gomock.Controller) *MockOauthValidator {
	mock := &MockOauthValidator{ctrl: ctrl}
	mock.recorder = &MockOauthValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthValidator) EXPECT() *MockOauthValidatorMockRecorder {
	return m.recorder
}

// ValidateApplication mocks base method
func (m *MockOauthValidator) ValidateApplication(ctx context.Context, data entity.OauthApplicationJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateApplication", ctx, data)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateApplication indicates an expected call of ValidateApplication
func (mr *MockOauthValidatorMockRecorder) ValidateApplication(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateApplication", reflect.TypeOf((*MockOauthValidator)(nil).ValidateApplication), ctx, data)
}

// ValidateAuthorizationGrant mocks base method
func (m *MockOauthValidator) ValidateAuthorizationGrant(ctx context.Context, r entity.AuthorizationRequestJSON, application entity.OauthApplication) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthorizationGrant", ctx, r, application)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateAuthorizationGrant indicates an expected call of ValidateAuthorizationGrant
func (mr *MockOauthValidatorMockRecorder) ValidateAuthorizationGrant(ctx, r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthorizationGrant", reflect.TypeOf((*MockOauthValidator)(nil).ValidateAuthorizationGrant), ctx, r, application)
}

// ValidateTokenGrant mocks base method
func (m *MockOauthValidator) ValidateTokenGrant(ctx context.Context, r entity.AccessTokenRequestJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTokenGrant", ctx, r)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateTokenGrant indicates an expected call of ValidateTokenGrant
func (mr *MockOauthValidatorMockRecorder) ValidateTokenGrant(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTokenGrant", reflect.TypeOf((*MockOauthValidator)(nil).ValidateTokenGrant), ctx, r)
}
