// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda/modules/orchestrator (interfaces: SharedCronjobRunner)

// Package orchestrator is a generated GoMock package.
package orchestrator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSharedCronjobRunner is a mock of SharedCronjobRunner interface.
type MockSharedCronjobRunner struct {
	ctrl     *gomock.Controller
	recorder *MockSharedCronjobRunnerMockRecorder
}

// MockSharedCronjobRunnerMockRecorder is the mock recorder for MockSharedCronjobRunner.
type MockSharedCronjobRunnerMockRecorder struct {
	mock *MockSharedCronjobRunner
}

// NewMockSharedCronjobRunner creates a new mock instance.
func NewMockSharedCronjobRunner(ctrl *gomock.Controller) *MockSharedCronjobRunner {
	mock := &MockSharedCronjobRunner{ctrl: ctrl}
	mock.recorder = &MockSharedCronjobRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedCronjobRunner) EXPECT() *MockSharedCronjobRunnerMockRecorder {
	return m.recorder
}

// SyncAddons mocks base method.
func (m *MockSharedCronjobRunner) SyncAddons() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncAddons")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncAddons indicates an expected call of SyncAddons.
func (mr *MockSharedCronjobRunnerMockRecorder) SyncAddons() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncAddons", reflect.TypeOf((*MockSharedCronjobRunner)(nil).SyncAddons))
}

// SyncProjects mocks base method.
func (m *MockSharedCronjobRunner) SyncProjects() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProjects")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncProjects indicates an expected call of SyncProjects.
func (mr *MockSharedCronjobRunnerMockRecorder) SyncProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProjects", reflect.TypeOf((*MockSharedCronjobRunner)(nil).SyncProjects))
}
