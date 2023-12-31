// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sijanstha/electronic-voting-system/internal/core/ports (interfaces: TokenService)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -package=mockservice -destination internal/core/services/mock/mock.service.go github.com/sijanstha/electronic-voting-system/internal/core/ports TokenService
//
// Package mockservice is a generated GoMock package.
package mockservice

import (
	reflect "reflect"

	domain "github.com/sijanstha/electronic-voting-system/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTokenService) Generate(arg0 domain.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTokenServiceMockRecorder) Generate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTokenService)(nil).Generate), arg0)
}

// Validate mocks base method.
func (m *MockTokenService) Validate(arg0 string) (*domain.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(*domain.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenServiceMockRecorder) Validate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenService)(nil).Validate), arg0)
}
