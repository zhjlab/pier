// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_syncer is a generated GoMock package.
package mock_syncer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	appchain_mgr "github.com/meshplus/bitxhub-core/appchain-mgr"
	pb "github.com/meshplus/bitxhub-model/pb"
	syncer "github.com/meshplus/pier/internal/syncer"
	model "github.com/meshplus/pier/pkg/model"
)

// MockSyncer is a mock of Syncer interface.
type MockSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncerMockRecorder
}

// MockSyncerMockRecorder is the mock recorder for MockSyncer.
type MockSyncerMockRecorder struct {
	mock *MockSyncer
}

// NewMockSyncer creates a new mock instance.
func NewMockSyncer(ctrl *gomock.Controller) *MockSyncer {
	mock := &MockSyncer{ctrl: ctrl}
	mock.recorder = &MockSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncer) EXPECT() *MockSyncerMockRecorder {
	return m.recorder
}

// GetAppchains mocks base method.
func (m *MockSyncer) GetAppchains() ([]*appchain_mgr.Appchain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppchains")
	ret0, _ := ret[0].([]*appchain_mgr.Appchain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppchains indicates an expected call of GetAppchains.
func (mr *MockSyncerMockRecorder) GetAppchains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppchains", reflect.TypeOf((*MockSyncer)(nil).GetAppchains))
}

// GetAssetExchangeSigns mocks base method.
func (m *MockSyncer) GetAssetExchangeSigns(id string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetExchangeSigns", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetExchangeSigns indicates an expected call of GetAssetExchangeSigns.
func (mr *MockSyncerMockRecorder) GetAssetExchangeSigns(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetExchangeSigns", reflect.TypeOf((*MockSyncer)(nil).GetAssetExchangeSigns), id)
}

// GetIBTPSigns mocks base method.
func (m *MockSyncer) GetIBTPSigns(ibtp *pb.IBTP) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIBTPSigns", ibtp)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIBTPSigns indicates an expected call of GetIBTPSigns.
func (mr *MockSyncerMockRecorder) GetIBTPSigns(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIBTPSigns", reflect.TypeOf((*MockSyncer)(nil).GetIBTPSigns), ibtp)
}

// GetInterchainById mocks base method.
func (m *MockSyncer) GetInterchainById(from string) *pb.Interchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterchainById", from)
	ret0, _ := ret[0].(*pb.Interchain)
	return ret0
}

// GetInterchainById indicates an expected call of GetInterchainById.
func (mr *MockSyncerMockRecorder) GetInterchainById(from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterchainById", reflect.TypeOf((*MockSyncer)(nil).GetInterchainById), from)
}

// GetTxStatus mocks base method.
func (m *MockSyncer) GetTxStatus(id string) (pb.TransactionStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxStatus", id)
	ret0, _ := ret[0].(pb.TransactionStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxStatus indicates an expected call of GetTxStatus.
func (mr *MockSyncerMockRecorder) GetTxStatus(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxStatus", reflect.TypeOf((*MockSyncer)(nil).GetTxStatus), id)
}

// ListenIBTP mocks base method.
func (m *MockSyncer) ListenIBTP() <-chan *model.WrappedIBTP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenIBTP")
	ret0, _ := ret[0].(<-chan *model.WrappedIBTP)
	return ret0
}

// ListenIBTP indicates an expected call of ListenIBTP.
func (mr *MockSyncerMockRecorder) ListenIBTP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenIBTP", reflect.TypeOf((*MockSyncer)(nil).ListenIBTP))
}

// QueryIBTP mocks base method.
func (m *MockSyncer) QueryIBTP(ibtpID string, isReq bool) (*pb.IBTP, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryIBTP", ibtpID, isReq)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryIBTP indicates an expected call of QueryIBTP.
func (mr *MockSyncerMockRecorder) QueryIBTP(ibtpID, isReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryIBTP", reflect.TypeOf((*MockSyncer)(nil).QueryIBTP), ibtpID, isReq)
}

// QueryInterchainMeta mocks base method.
func (m *MockSyncer) QueryInterchainMeta(servicePair string) *pb.Interchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryInterchainMeta", servicePair)
	ret0, _ := ret[0].(*pb.Interchain)
	return ret0
}

// QueryInterchainMeta indicates an expected call of QueryInterchainMeta.
func (mr *MockSyncerMockRecorder) QueryInterchainMeta(servicePair interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryInterchainMeta", reflect.TypeOf((*MockSyncer)(nil).QueryInterchainMeta), servicePair)
}

// RegisterAppchainHandler mocks base method.
func (m *MockSyncer) RegisterAppchainHandler(handler syncer.AppchainHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAppchainHandler", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAppchainHandler indicates an expected call of RegisterAppchainHandler.
func (mr *MockSyncerMockRecorder) RegisterAppchainHandler(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAppchainHandler", reflect.TypeOf((*MockSyncer)(nil).RegisterAppchainHandler), handler)
}

// RegisterRecoverHandler mocks base method.
func (m *MockSyncer) RegisterRecoverHandler(arg0 syncer.RecoverUnionHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRecoverHandler", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRecoverHandler indicates an expected call of RegisterRecoverHandler.
func (mr *MockSyncerMockRecorder) RegisterRecoverHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRecoverHandler", reflect.TypeOf((*MockSyncer)(nil).RegisterRecoverHandler), arg0)
}

// RegisterRollbackHandler mocks base method.
func (m *MockSyncer) RegisterRollbackHandler(handler syncer.RollbackHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRollbackHandler", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRollbackHandler indicates an expected call of RegisterRollbackHandler.
func (mr *MockSyncerMockRecorder) RegisterRollbackHandler(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRollbackHandler", reflect.TypeOf((*MockSyncer)(nil).RegisterRollbackHandler), handler)
}

// SendIBTP mocks base method.
func (m *MockSyncer) SendIBTP(ibtp *pb.IBTP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendIBTP", ibtp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendIBTP indicates an expected call of SendIBTP.
func (mr *MockSyncerMockRecorder) SendIBTP(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendIBTP", reflect.TypeOf((*MockSyncer)(nil).SendIBTP), ibtp)
}

// Start mocks base method.
func (m *MockSyncer) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockSyncerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSyncer)(nil).Start))
}

// Stop mocks base method.
func (m *MockSyncer) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockSyncerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSyncer)(nil).Stop))
}
