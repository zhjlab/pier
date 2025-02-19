// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/meshplus/bitxhub-model/pb"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAppchainInfo mocks base method.
func (m *MockClient) GetAppchainInfo(chainID string) (string, []byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppchainInfo", chainID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetAppchainInfo indicates an expected call of GetAppchainInfo.
func (mr *MockClientMockRecorder) GetAppchainInfo(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppchainInfo", reflect.TypeOf((*MockClient)(nil).GetAppchainInfo), chainID)
}

// GetCallbackMeta mocks base method.
func (m *MockClient) GetCallbackMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallbackMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallbackMeta indicates an expected call of GetCallbackMeta.
func (mr *MockClientMockRecorder) GetCallbackMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallbackMeta", reflect.TypeOf((*MockClient)(nil).GetCallbackMeta))
}

// GetChainID mocks base method.
func (m *MockClient) GetChainID() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetChainID indicates an expected call of GetChainID.
func (mr *MockClientMockRecorder) GetChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*MockClient)(nil).GetChainID))
}

// GetDirectTransactionMeta mocks base method.
func (m *MockClient) GetDirectTransactionMeta(arg0 string) (uint64, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectTransactionMeta", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetDirectTransactionMeta indicates an expected call of GetDirectTransactionMeta.
func (mr *MockClientMockRecorder) GetDirectTransactionMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectTransactionMeta", reflect.TypeOf((*MockClient)(nil).GetDirectTransactionMeta), arg0)
}

// GetDstRollbackMeta mocks base method.
func (m *MockClient) GetDstRollbackMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDstRollbackMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDstRollbackMeta indicates an expected call of GetDstRollbackMeta.
func (mr *MockClientMockRecorder) GetDstRollbackMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDstRollbackMeta", reflect.TypeOf((*MockClient)(nil).GetDstRollbackMeta))
}

// GetIBTPCh mocks base method.
func (m *MockClient) GetIBTPCh() chan *pb.IBTP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIBTPCh")
	ret0, _ := ret[0].(chan *pb.IBTP)
	return ret0
}

// GetIBTPCh indicates an expected call of GetIBTPCh.
func (mr *MockClientMockRecorder) GetIBTPCh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIBTPCh", reflect.TypeOf((*MockClient)(nil).GetIBTPCh))
}

// GetInMeta mocks base method.
func (m *MockClient) GetInMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInMeta indicates an expected call of GetInMeta.
func (mr *MockClientMockRecorder) GetInMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInMeta", reflect.TypeOf((*MockClient)(nil).GetInMeta))
}

// GetOffChainData mocks base method.
func (m *MockClient) GetOffChainData(request *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffChainData", request)
	ret0, _ := ret[0].(*pb.GetDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffChainData indicates an expected call of GetOffChainData.
func (mr *MockClientMockRecorder) GetOffChainData(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffChainData", reflect.TypeOf((*MockClient)(nil).GetOffChainData), request)
}

// GetOffChainDataReq mocks base method.
func (m *MockClient) GetOffChainDataReq() chan *pb.GetDataRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffChainDataReq")
	ret0, _ := ret[0].(chan *pb.GetDataRequest)
	return ret0
}

// GetOffChainDataReq indicates an expected call of GetOffChainDataReq.
func (mr *MockClientMockRecorder) GetOffChainDataReq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffChainDataReq", reflect.TypeOf((*MockClient)(nil).GetOffChainDataReq))
}

// GetOutMessage mocks base method.
func (m *MockClient) GetOutMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutMessage", servicePair, idx)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutMessage indicates an expected call of GetOutMessage.
func (mr *MockClientMockRecorder) GetOutMessage(servicePair, idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutMessage", reflect.TypeOf((*MockClient)(nil).GetOutMessage), servicePair, idx)
}

// GetOutMeta mocks base method.
func (m *MockClient) GetOutMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutMeta indicates an expected call of GetOutMeta.
func (mr *MockClientMockRecorder) GetOutMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutMeta", reflect.TypeOf((*MockClient)(nil).GetOutMeta))
}

// GetReceiptMessage mocks base method.
func (m *MockClient) GetReceiptMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptMessage", servicePair, idx)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptMessage indicates an expected call of GetReceiptMessage.
func (mr *MockClientMockRecorder) GetReceiptMessage(servicePair, idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptMessage", reflect.TypeOf((*MockClient)(nil).GetReceiptMessage), servicePair, idx)
}

// GetServices mocks base method.
func (m *MockClient) GetServices() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices.
func (mr *MockClientMockRecorder) GetServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockClient)(nil).GetServices))
}

// GetUpdateMeta mocks base method.
func (m *MockClient) GetUpdateMeta() chan *pb.UpdateMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateMeta")
	ret0, _ := ret[0].(chan *pb.UpdateMeta)
	return ret0
}

// GetUpdateMeta indicates an expected call of GetUpdateMeta.
func (mr *MockClientMockRecorder) GetUpdateMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateMeta", reflect.TypeOf((*MockClient)(nil).GetUpdateMeta))
}

// Initialize mocks base method.
func (m *MockClient) Initialize(configPath string, extra []byte, mode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", configPath, extra, mode)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockClientMockRecorder) Initialize(configPath, extra, mode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize), configPath, extra, mode)
}

// Name mocks base method.
func (m *MockClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockClient)(nil).Name))
}

// Start mocks base method.
func (m *MockClient) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start))
}

// Stop mocks base method.
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// SubmitIBTP mocks base method.
func (m *MockClient) SubmitIBTP(from string, index uint64, serviceID string, ibtpType pb.IBTP_Type, content *pb.Content, proof *pb.BxhProof, isEncrypted bool) (*pb.SubmitIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitIBTP", from, index, serviceID, ibtpType, content, proof, isEncrypted)
	ret0, _ := ret[0].(*pb.SubmitIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitIBTP indicates an expected call of SubmitIBTP.
func (mr *MockClientMockRecorder) SubmitIBTP(from, index, serviceID, ibtpType, content, proof, isEncrypted interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitIBTP", reflect.TypeOf((*MockClient)(nil).SubmitIBTP), from, index, serviceID, ibtpType, content, proof, isEncrypted)
}

// SubmitIBTPBatch mocks base method.
func (m *MockClient) SubmitIBTPBatch(from []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, content []*pb.Content, proof []*pb.BxhProof, isEncrypted []bool) (*pb.SubmitIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitIBTPBatch", from, index, serviceID, ibtpType, content, proof, isEncrypted)
	ret0, _ := ret[0].(*pb.SubmitIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitIBTPBatch indicates an expected call of SubmitIBTPBatch.
func (mr *MockClientMockRecorder) SubmitIBTPBatch(from, index, serviceID, ibtpType, content, proof, isEncrypted interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitIBTPBatch", reflect.TypeOf((*MockClient)(nil).SubmitIBTPBatch), from, index, serviceID, ibtpType, content, proof, isEncrypted)
}

// SubmitOffChainData mocks base method.
func (m *MockClient) SubmitOffChainData(response *pb.GetDataResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOffChainData", response)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitOffChainData indicates an expected call of SubmitOffChainData.
func (mr *MockClientMockRecorder) SubmitOffChainData(response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOffChainData", reflect.TypeOf((*MockClient)(nil).SubmitOffChainData), response)
}

// SubmitReceipt mocks base method.
func (m *MockClient) SubmitReceipt(to string, index uint64, serviceID string, ibtpType pb.IBTP_Type, result *pb.Result, proof *pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitReceipt", to, index, serviceID, ibtpType, result, proof)
	ret0, _ := ret[0].(*pb.SubmitIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitReceipt indicates an expected call of SubmitReceipt.
func (mr *MockClientMockRecorder) SubmitReceipt(to, index, serviceID, ibtpType, result, proof interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitReceipt", reflect.TypeOf((*MockClient)(nil).SubmitReceipt), to, index, serviceID, ibtpType, result, proof)
}

// SubmitReceiptBatch mocks base method.
func (m *MockClient) SubmitReceiptBatch(to []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, result []*pb.Result, proof []*pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitReceiptBatch", to, index, serviceID, ibtpType, result, proof)
	ret0, _ := ret[0].(*pb.SubmitIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitReceiptBatch indicates an expected call of SubmitReceiptBatch.
func (mr *MockClientMockRecorder) SubmitReceiptBatch(to, index, serviceID, ibtpType, result, proof interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitReceiptBatch", reflect.TypeOf((*MockClient)(nil).SubmitReceiptBatch), to, index, serviceID, ibtpType, result, proof)
}

// Type mocks base method.
func (m *MockClient) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockClientMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockClient)(nil).Type))
}
