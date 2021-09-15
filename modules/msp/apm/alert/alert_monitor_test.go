// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package alert

import (
	context "context"
	reflect "reflect"

	pb "github.com/erda-project/erda-proto-go/core/monitor/alert/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockAlertServiceServer is a mock of AlertServiceServer interface.
type MockAlertServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAlertServiceServerMockRecorder
}

// MockAlertServiceServerMockRecorder is the mock recorder for MockAlertServiceServer.
type MockAlertServiceServerMockRecorder struct {
	mock *MockAlertServiceServer
}

// NewMockAlertServiceServer creates a new mock instance.
func NewMockAlertServiceServer(ctrl *gomock.Controller) *MockAlertServiceServer {
	mock := &MockAlertServiceServer{ctrl: ctrl}
	mock.recorder = &MockAlertServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertServiceServer) EXPECT() *MockAlertServiceServerMockRecorder {
	return m.recorder
}

// CreateAlert mocks base method.
func (m *MockAlertServiceServer) CreateAlert(arg0 context.Context, arg1 *pb.CreateAlertRequest) (*pb.CreateAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlert indicates an expected call of CreateAlert.
func (mr *MockAlertServiceServerMockRecorder) CreateAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateAlert), arg0, arg1)
}

// CreateAlertIssue mocks base method.
func (m *MockAlertServiceServer) CreateAlertIssue(arg0 context.Context, arg1 *pb.CreateAlertIssueRequest) (*pb.CreateAlertIssueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlertIssue", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateAlertIssueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlertIssue indicates an expected call of CreateAlertIssue.
func (mr *MockAlertServiceServerMockRecorder) CreateAlertIssue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlertIssue", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateAlertIssue), arg0, arg1)
}

// CreateCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) CreateCustomizeAlert(arg0 context.Context, arg1 *pb.CreateCustomizeAlertRequest) (*pb.CreateCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomizeAlert indicates an expected call of CreateCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) CreateCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateCustomizeAlert), arg0, arg1)
}

// CreateOrgAlert mocks base method.
func (m *MockAlertServiceServer) CreateOrgAlert(arg0 context.Context, arg1 *pb.CreateOrgAlertRequest) (*pb.CreateOrgAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateOrgAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAlert indicates an expected call of CreateOrgAlert.
func (mr *MockAlertServiceServerMockRecorder) CreateOrgAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateOrgAlert), arg0, arg1)
}

// CreateOrgAlertIssue mocks base method.
func (m *MockAlertServiceServer) CreateOrgAlertIssue(arg0 context.Context, arg1 *pb.CreateOrgAlertIssueRequest) (*pb.CreateOrgAlertIssueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgAlertIssue", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateOrgAlertIssueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAlertIssue indicates an expected call of CreateOrgAlertIssue.
func (mr *MockAlertServiceServerMockRecorder) CreateOrgAlertIssue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAlertIssue", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateOrgAlertIssue), arg0, arg1)
}

// CreateOrgCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) CreateOrgCustomizeAlert(arg0 context.Context, arg1 *pb.CreateOrgCustomizeAlertRequest) (*pb.CreateOrgCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateOrgCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgCustomizeAlert indicates an expected call of CreateOrgCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) CreateOrgCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).CreateOrgCustomizeAlert), arg0, arg1)
}

// DeleteAlert mocks base method.
func (m *MockAlertServiceServer) DeleteAlert(arg0 context.Context, arg1 *pb.DeleteAlertRequest) (*pb.DeleteAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAlert indicates an expected call of DeleteAlert.
func (mr *MockAlertServiceServerMockRecorder) DeleteAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).DeleteAlert), arg0, arg1)
}

// DeleteCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) DeleteCustomizeAlert(arg0 context.Context, arg1 *pb.DeleteCustomizeAlertRequest) (*pb.DeleteCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCustomizeAlert indicates an expected call of DeleteCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) DeleteCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).DeleteCustomizeAlert), arg0, arg1)
}

// DeleteOrgAlert mocks base method.
func (m *MockAlertServiceServer) DeleteOrgAlert(arg0 context.Context, arg1 *pb.DeleteOrgAlertRequest) (*pb.DeleteOrgAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteOrgAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgAlert indicates an expected call of DeleteOrgAlert.
func (mr *MockAlertServiceServerMockRecorder) DeleteOrgAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).DeleteOrgAlert), arg0, arg1)
}

// DeleteOrgCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) DeleteOrgCustomizeAlert(arg0 context.Context, arg1 *pb.DeleteOrgCustomizeAlertRequest) (*pb.DeleteOrgCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteOrgCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgCustomizeAlert indicates an expected call of DeleteOrgCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) DeleteOrgCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).DeleteOrgCustomizeAlert), arg0, arg1)
}

// GetAlert mocks base method.
func (m *MockAlertServiceServer) GetAlert(arg0 context.Context, arg1 *pb.GetAlertRequest) (*pb.GetAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlert indicates an expected call of GetAlert.
func (mr *MockAlertServiceServerMockRecorder) GetAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).GetAlert), arg0, arg1)
}

// GetAlertDetail mocks base method.
func (m *MockAlertServiceServer) GetAlertDetail(arg0 context.Context, arg1 *pb.GetAlertDetailRequest) (*pb.GetAlertDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlertDetail", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAlertDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertDetail indicates an expected call of GetAlertDetail.
func (mr *MockAlertServiceServerMockRecorder) GetAlertDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertDetail", reflect.TypeOf((*MockAlertServiceServer)(nil).GetAlertDetail), arg0, arg1)
}

// GetAlertRecord mocks base method.
func (m *MockAlertServiceServer) GetAlertRecord(arg0 context.Context, arg1 *pb.GetAlertRecordRequest) (*pb.GetAlertRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlertRecord", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAlertRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertRecord indicates an expected call of GetAlertRecord.
func (mr *MockAlertServiceServerMockRecorder) GetAlertRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertRecord", reflect.TypeOf((*MockAlertServiceServer)(nil).GetAlertRecord), arg0, arg1)
}

// GetAlertRecordAttr mocks base method.
func (m *MockAlertServiceServer) GetAlertRecordAttr(arg0 context.Context, arg1 *pb.GetAlertRecordAttrRequest) (*pb.GetAlertRecordAttrResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlertRecordAttr", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAlertRecordAttrResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertRecordAttr indicates an expected call of GetAlertRecordAttr.
func (mr *MockAlertServiceServerMockRecorder) GetAlertRecordAttr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertRecordAttr", reflect.TypeOf((*MockAlertServiceServer)(nil).GetAlertRecordAttr), arg0, arg1)
}

// GetCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) GetCustomizeAlert(arg0 context.Context, arg1 *pb.GetCustomizeAlertRequest) (*pb.GetCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomizeAlert indicates an expected call of GetCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) GetCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).GetCustomizeAlert), arg0, arg1)
}

// GetCustomizeAlertDetail mocks base method.
func (m *MockAlertServiceServer) GetCustomizeAlertDetail(arg0 context.Context, arg1 *pb.GetCustomizeAlertDetailRequest) (*pb.GetCustomizeAlertDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomizeAlertDetail", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetCustomizeAlertDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomizeAlertDetail indicates an expected call of GetCustomizeAlertDetail.
func (mr *MockAlertServiceServerMockRecorder) GetCustomizeAlertDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomizeAlertDetail", reflect.TypeOf((*MockAlertServiceServer)(nil).GetCustomizeAlertDetail), arg0, arg1)
}

// GetOrgAlertDetail mocks base method.
func (m *MockAlertServiceServer) GetOrgAlertDetail(arg0 context.Context, arg1 *pb.GetOrgAlertDetailRequest) (*pb.GetOrgAlertDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgAlertDetail", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetOrgAlertDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgAlertDetail indicates an expected call of GetOrgAlertDetail.
func (mr *MockAlertServiceServerMockRecorder) GetOrgAlertDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgAlertDetail", reflect.TypeOf((*MockAlertServiceServer)(nil).GetOrgAlertDetail), arg0, arg1)
}

// GetOrgAlertRecord mocks base method.
func (m *MockAlertServiceServer) GetOrgAlertRecord(arg0 context.Context, arg1 *pb.GetOrgAlertRecordRequest) (*pb.GetOrgAlertRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgAlertRecord", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetOrgAlertRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgAlertRecord indicates an expected call of GetOrgAlertRecord.
func (mr *MockAlertServiceServerMockRecorder) GetOrgAlertRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgAlertRecord", reflect.TypeOf((*MockAlertServiceServer)(nil).GetOrgAlertRecord), arg0, arg1)
}

// GetOrgAlertRecordAttr mocks base method.
func (m *MockAlertServiceServer) GetOrgAlertRecordAttr(arg0 context.Context, arg1 *pb.GetOrgAlertRecordAttrRequest) (*pb.GetOrgAlertRecordAttrResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgAlertRecordAttr", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetOrgAlertRecordAttrResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgAlertRecordAttr indicates an expected call of GetOrgAlertRecordAttr.
func (mr *MockAlertServiceServerMockRecorder) GetOrgAlertRecordAttr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgAlertRecordAttr", reflect.TypeOf((*MockAlertServiceServer)(nil).GetOrgAlertRecordAttr), arg0, arg1)
}

// GetOrgCustomizeAlertDetail mocks base method.
func (m *MockAlertServiceServer) GetOrgCustomizeAlertDetail(arg0 context.Context, arg1 *pb.GetOrgCustomizeAlertDetailRequest) (*pb.GetOrgCustomizeAlertDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgCustomizeAlertDetail", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetOrgCustomizeAlertDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgCustomizeAlertDetail indicates an expected call of GetOrgCustomizeAlertDetail.
func (mr *MockAlertServiceServerMockRecorder) GetOrgCustomizeAlertDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgCustomizeAlertDetail", reflect.TypeOf((*MockAlertServiceServer)(nil).GetOrgCustomizeAlertDetail), arg0, arg1)
}

// QueryAlert mocks base method.
func (m *MockAlertServiceServer) QueryAlert(arg0 context.Context, arg1 *pb.QueryAlertRequest) (*pb.QueryAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAlert indicates an expected call of QueryAlert.
func (mr *MockAlertServiceServerMockRecorder) QueryAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryAlert), arg0, arg1)
}

// QueryAlertHistory mocks base method.
func (m *MockAlertServiceServer) QueryAlertHistory(arg0 context.Context, arg1 *pb.QueryAlertHistoryRequest) (*pb.QueryAlertHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAlertHistory", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryAlertHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAlertHistory indicates an expected call of QueryAlertHistory.
func (mr *MockAlertServiceServerMockRecorder) QueryAlertHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAlertHistory", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryAlertHistory), arg0, arg1)
}

// QueryAlertRecord mocks base method.
func (m *MockAlertServiceServer) QueryAlertRecord(arg0 context.Context, arg1 *pb.QueryAlertRecordRequest) (*pb.QueryAlertRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAlertRecord", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryAlertRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAlertRecord indicates an expected call of QueryAlertRecord.
func (mr *MockAlertServiceServerMockRecorder) QueryAlertRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAlertRecord", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryAlertRecord), arg0, arg1)
}

// QueryAlertRule mocks base method.
func (m *MockAlertServiceServer) QueryAlertRule(arg0 context.Context, arg1 *pb.QueryAlertRuleRequest) (*pb.QueryAlertRuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAlertRule", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryAlertRuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAlertRule indicates an expected call of QueryAlertRule.
func (mr *MockAlertServiceServerMockRecorder) QueryAlertRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAlertRule", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryAlertRule), arg0, arg1)
}

// QueryCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) QueryCustomizeAlert(arg0 context.Context, arg1 *pb.QueryCustomizeAlertRequest) (*pb.QueryCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCustomizeAlert indicates an expected call of QueryCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) QueryCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryCustomizeAlert), arg0, arg1)
}

// QueryCustomizeMetric mocks base method.
func (m *MockAlertServiceServer) QueryCustomizeMetric(arg0 context.Context, arg1 *pb.QueryCustomizeMetricRequest) (*pb.QueryCustomizeMetricResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCustomizeMetric", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryCustomizeMetricResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCustomizeMetric indicates an expected call of QueryCustomizeMetric.
func (mr *MockAlertServiceServerMockRecorder) QueryCustomizeMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCustomizeMetric", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryCustomizeMetric), arg0, arg1)
}

// QueryCustomizeNotifyTarget mocks base method.
func (m *MockAlertServiceServer) QueryCustomizeNotifyTarget(arg0 context.Context, arg1 *pb.QueryCustomizeNotifyTargetRequest) (*pb.QueryCustomizeNotifyTargetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCustomizeNotifyTarget", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryCustomizeNotifyTargetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCustomizeNotifyTarget indicates an expected call of QueryCustomizeNotifyTarget.
func (mr *MockAlertServiceServerMockRecorder) QueryCustomizeNotifyTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCustomizeNotifyTarget", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryCustomizeNotifyTarget), arg0, arg1)
}

// QueryDashboardByAlert mocks base method.
func (m *MockAlertServiceServer) QueryDashboardByAlert(arg0 context.Context, arg1 *pb.QueryDashboardByAlertRequest) (*pb.QueryDashboardByAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryDashboardByAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryDashboardByAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDashboardByAlert indicates an expected call of QueryDashboardByAlert.
func (mr *MockAlertServiceServerMockRecorder) QueryDashboardByAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDashboardByAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryDashboardByAlert), arg0, arg1)
}

// QueryOrgAlert mocks base method.
func (m *MockAlertServiceServer) QueryOrgAlert(arg0 context.Context, arg1 *pb.QueryOrgAlertRequest) (*pb.QueryOrgAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgAlert indicates an expected call of QueryOrgAlert.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgAlert), arg0, arg1)
}

// QueryOrgAlertHistory mocks base method.
func (m *MockAlertServiceServer) QueryOrgAlertHistory(arg0 context.Context, arg1 *pb.QueryOrgAlertHistoryRequest) (*pb.QueryOrgAlertHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgAlertHistory", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgAlertHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgAlertHistory indicates an expected call of QueryOrgAlertHistory.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgAlertHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgAlertHistory", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgAlertHistory), arg0, arg1)
}

// QueryOrgAlertRecord mocks base method.
func (m *MockAlertServiceServer) QueryOrgAlertRecord(arg0 context.Context, arg1 *pb.QueryOrgAlertRecordRequest) (*pb.QueryOrgAlertRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgAlertRecord", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgAlertRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgAlertRecord indicates an expected call of QueryOrgAlertRecord.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgAlertRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgAlertRecord", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgAlertRecord), arg0, arg1)
}

// QueryOrgAlertRule mocks base method.
func (m *MockAlertServiceServer) QueryOrgAlertRule(arg0 context.Context, arg1 *pb.QueryOrgAlertRuleRequest) (*pb.QueryOrgAlertRuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgAlertRule", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgAlertRuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgAlertRule indicates an expected call of QueryOrgAlertRule.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgAlertRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgAlertRule", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgAlertRule), arg0, arg1)
}

// QueryOrgCustomizeAlerts mocks base method.
func (m *MockAlertServiceServer) QueryOrgCustomizeAlerts(arg0 context.Context, arg1 *pb.QueryOrgCustomizeAlertsRequest) (*pb.QueryOrgCustomizeAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgCustomizeAlerts", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgCustomizeAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgCustomizeAlerts indicates an expected call of QueryOrgCustomizeAlerts.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgCustomizeAlerts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgCustomizeAlerts", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgCustomizeAlerts), arg0, arg1)
}

// QueryOrgCustomizeMetric mocks base method.
func (m *MockAlertServiceServer) QueryOrgCustomizeMetric(arg0 context.Context, arg1 *pb.QueryOrgCustomizeMetricRequest) (*pb.QueryOrgCustomizeMetricResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgCustomizeMetric", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgCustomizeMetricResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgCustomizeMetric indicates an expected call of QueryOrgCustomizeMetric.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgCustomizeMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgCustomizeMetric", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgCustomizeMetric), arg0, arg1)
}

// QueryOrgCustomizeNotifyTarget mocks base method.
func (m *MockAlertServiceServer) QueryOrgCustomizeNotifyTarget(arg0 context.Context, arg1 *pb.QueryOrgCustomizeNotifyTargetRequest) (*pb.QueryOrgCustomizeNotifyTargetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgCustomizeNotifyTarget", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgCustomizeNotifyTargetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgCustomizeNotifyTarget indicates an expected call of QueryOrgCustomizeNotifyTarget.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgCustomizeNotifyTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgCustomizeNotifyTarget", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgCustomizeNotifyTarget), arg0, arg1)
}

// QueryOrgDashboardByAlert mocks base method.
func (m *MockAlertServiceServer) QueryOrgDashboardByAlert(arg0 context.Context, arg1 *pb.QueryOrgDashboardByAlertRequest) (*pb.QueryOrgDashboardByAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgDashboardByAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgDashboardByAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgDashboardByAlert indicates an expected call of QueryOrgDashboardByAlert.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgDashboardByAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgDashboardByAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgDashboardByAlert), arg0, arg1)
}

// QueryOrgHostsAlertRecord mocks base method.
func (m *MockAlertServiceServer) QueryOrgHostsAlertRecord(arg0 context.Context, arg1 *pb.QueryOrgHostsAlertRecordRequest) (*pb.QueryOrgAlertRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgHostsAlertRecord", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryOrgAlertRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgHostsAlertRecord indicates an expected call of QueryOrgHostsAlertRecord.
func (mr *MockAlertServiceServerMockRecorder) QueryOrgHostsAlertRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgHostsAlertRecord", reflect.TypeOf((*MockAlertServiceServer)(nil).QueryOrgHostsAlertRecord), arg0, arg1)
}

// UpdateAlert mocks base method.
func (m *MockAlertServiceServer) UpdateAlert(arg0 context.Context, arg1 *pb.UpdateAlertRequest) (*pb.UpdateAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAlert indicates an expected call of UpdateAlert.
func (mr *MockAlertServiceServerMockRecorder) UpdateAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateAlert), arg0, arg1)
}

// UpdateAlertEnable mocks base method.
func (m *MockAlertServiceServer) UpdateAlertEnable(arg0 context.Context, arg1 *pb.UpdateAlertEnableRequest) (*pb.UpdateAlertEnableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlertEnable", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateAlertEnableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAlertEnable indicates an expected call of UpdateAlertEnable.
func (mr *MockAlertServiceServerMockRecorder) UpdateAlertEnable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlertEnable", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateAlertEnable), arg0, arg1)
}

// UpdateAlertIssue mocks base method.
func (m *MockAlertServiceServer) UpdateAlertIssue(arg0 context.Context, arg1 *pb.UpdateAlertIssueRequest) (*pb.UpdateAlertIssueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlertIssue", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateAlertIssueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAlertIssue indicates an expected call of UpdateAlertIssue.
func (mr *MockAlertServiceServerMockRecorder) UpdateAlertIssue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlertIssue", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateAlertIssue), arg0, arg1)
}

// UpdateCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) UpdateCustomizeAlert(arg0 context.Context, arg1 *pb.UpdateCustomizeAlertRequest) (*pb.UpdateCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCustomizeAlert indicates an expected call of UpdateCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) UpdateCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateCustomizeAlert), arg0, arg1)
}

// UpdateCustomizeAlertEnable mocks base method.
func (m *MockAlertServiceServer) UpdateCustomizeAlertEnable(arg0 context.Context, arg1 *pb.UpdateCustomizeAlertEnableRequest) (*pb.UpdateCustomizeAlertEnableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomizeAlertEnable", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateCustomizeAlertEnableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCustomizeAlertEnable indicates an expected call of UpdateCustomizeAlertEnable.
func (mr *MockAlertServiceServerMockRecorder) UpdateCustomizeAlertEnable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomizeAlertEnable", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateCustomizeAlertEnable), arg0, arg1)
}

// UpdateOrgAlert mocks base method.
func (m *MockAlertServiceServer) UpdateOrgAlert(arg0 context.Context, arg1 *pb.UpdateOrgAlertRequest) (*pb.UpdateOrgAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateOrgAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgAlert indicates an expected call of UpdateOrgAlert.
func (mr *MockAlertServiceServerMockRecorder) UpdateOrgAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateOrgAlert), arg0, arg1)
}

// UpdateOrgAlertEnable mocks base method.
func (m *MockAlertServiceServer) UpdateOrgAlertEnable(arg0 context.Context, arg1 *pb.UpdateOrgAlertEnableRequest) (*pb.UpdateOrgAlertEnableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgAlertEnable", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateOrgAlertEnableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgAlertEnable indicates an expected call of UpdateOrgAlertEnable.
func (mr *MockAlertServiceServerMockRecorder) UpdateOrgAlertEnable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgAlertEnable", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateOrgAlertEnable), arg0, arg1)
}

// UpdateOrgAlertIssue mocks base method.
func (m *MockAlertServiceServer) UpdateOrgAlertIssue(arg0 context.Context, arg1 *pb.UpdateOrgAlertIssueRequest) (*pb.UpdateOrgAlertIssueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgAlertIssue", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateOrgAlertIssueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgAlertIssue indicates an expected call of UpdateOrgAlertIssue.
func (mr *MockAlertServiceServerMockRecorder) UpdateOrgAlertIssue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgAlertIssue", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateOrgAlertIssue), arg0, arg1)
}

// UpdateOrgCustomizeAlert mocks base method.
func (m *MockAlertServiceServer) UpdateOrgCustomizeAlert(arg0 context.Context, arg1 *pb.UpdateOrgCustomizeAlertRequest) (*pb.UpdateOrgCustomizeAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgCustomizeAlert", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateOrgCustomizeAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgCustomizeAlert indicates an expected call of UpdateOrgCustomizeAlert.
func (mr *MockAlertServiceServerMockRecorder) UpdateOrgCustomizeAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgCustomizeAlert", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateOrgCustomizeAlert), arg0, arg1)
}

// UpdateOrgCustomizeAlertEnable mocks base method.
func (m *MockAlertServiceServer) UpdateOrgCustomizeAlertEnable(arg0 context.Context, arg1 *pb.UpdateOrgCustomizeAlertEnableRequest) (*pb.UpdateOrgCustomizeAlertEnableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgCustomizeAlertEnable", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateOrgCustomizeAlertEnableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgCustomizeAlertEnable indicates an expected call of UpdateOrgCustomizeAlertEnable.
func (mr *MockAlertServiceServerMockRecorder) UpdateOrgCustomizeAlertEnable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgCustomizeAlertEnable", reflect.TypeOf((*MockAlertServiceServer)(nil).UpdateOrgCustomizeAlertEnable), arg0, arg1)
}
