// Code generated by MockGen. DO NOT EDIT.
// Source: port.pb.go

// Package mock_proto_grpc is a generated GoMock package.
package mock_proto_grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/sin-ivan/port-project/proto"
	grpc "google.golang.org/grpc"
)

// MockPortHandlerClient is a mock of PortHandlerClient interface
type MockPortHandlerClient struct {
	ctrl     *gomock.Controller
	recorder *MockPortHandlerClientMockRecorder
}

// MockPortHandlerClientMockRecorder is the mock recorder for MockPortHandlerClient
type MockPortHandlerClientMockRecorder struct {
	mock *MockPortHandlerClient
}

// NewMockPortHandlerClient creates a new mock instance
func NewMockPortHandlerClient(ctrl *gomock.Controller) *MockPortHandlerClient {
	mock := &MockPortHandlerClient{ctrl: ctrl}
	mock.recorder = &MockPortHandlerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortHandlerClient) EXPECT() *MockPortHandlerClientMockRecorder {
	return m.recorder
}

// GetPort mocks base method
func (m *MockPortHandlerClient) GetPort(ctx context.Context, in *proto.SingleRequest, opts ...grpc.CallOption) (*proto.Port, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPort", varargs...)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPort indicates an expected call of GetPort
func (mr *MockPortHandlerClientMockRecorder) GetPort(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockPortHandlerClient)(nil).GetPort), varargs...)
}

// GetAll mocks base method
func (m *MockPortHandlerClient) GetAll(ctx context.Context, in *proto.EmptyRequest, opts ...grpc.CallOption) (*proto.ListPort, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAll", varargs...)
	ret0, _ := ret[0].(*proto.ListPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPortHandlerClientMockRecorder) GetAll(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPortHandlerClient)(nil).GetAll), varargs...)
}

// Store mocks base method
func (m *MockPortHandlerClient) Store(ctx context.Context, in *proto.Port, opts ...grpc.CallOption) (*proto.Port, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Store", varargs...)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockPortHandlerClientMockRecorder) Store(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockPortHandlerClient)(nil).Store), varargs...)
}

// Update mocks base method
func (m *MockPortHandlerClient) Update(ctx context.Context, in *proto.Port, opts ...grpc.CallOption) (*proto.Port, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockPortHandlerClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPortHandlerClient)(nil).Update), varargs...)
}

// Delete mocks base method
func (m *MockPortHandlerClient) Delete(ctx context.Context, in *proto.SingleRequest, opts ...grpc.CallOption) (*proto.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*proto.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockPortHandlerClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPortHandlerClient)(nil).Delete), varargs...)
}

// MockPortHandlerServer is a mock of PortHandlerServer interface
type MockPortHandlerServer struct {
	ctrl     *gomock.Controller
	recorder *MockPortHandlerServerMockRecorder
}

// MockPortHandlerServerMockRecorder is the mock recorder for MockPortHandlerServer
type MockPortHandlerServerMockRecorder struct {
	mock *MockPortHandlerServer
}

// NewMockPortHandlerServer creates a new mock instance
func NewMockPortHandlerServer(ctrl *gomock.Controller) *MockPortHandlerServer {
	mock := &MockPortHandlerServer{ctrl: ctrl}
	mock.recorder = &MockPortHandlerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortHandlerServer) EXPECT() *MockPortHandlerServerMockRecorder {
	return m.recorder
}

// GetPort mocks base method
func (m *MockPortHandlerServer) GetPort(arg0 context.Context, arg1 *proto.SingleRequest) (*proto.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", arg0, arg1)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPort indicates an expected call of GetPort
func (mr *MockPortHandlerServerMockRecorder) GetPort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockPortHandlerServer)(nil).GetPort), arg0, arg1)
}

// GetAll mocks base method
func (m *MockPortHandlerServer) GetAll(arg0 context.Context, arg1 *proto.EmptyRequest) (*proto.ListPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].(*proto.ListPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPortHandlerServerMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPortHandlerServer)(nil).GetAll), arg0, arg1)
}

// Store mocks base method
func (m *MockPortHandlerServer) Store(arg0 context.Context, arg1 *proto.Port) (*proto.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockPortHandlerServerMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockPortHandlerServer)(nil).Store), arg0, arg1)
}

// Update mocks base method
func (m *MockPortHandlerServer) Update(arg0 context.Context, arg1 *proto.Port) (*proto.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*proto.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockPortHandlerServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPortHandlerServer)(nil).Update), arg0, arg1)
}

// Delete mocks base method
func (m *MockPortHandlerServer) Delete(arg0 context.Context, arg1 *proto.SingleRequest) (*proto.DeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockPortHandlerServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPortHandlerServer)(nil).Delete), arg0, arg1)
}
