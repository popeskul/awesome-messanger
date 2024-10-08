// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/popeskul/awesome-messanger/services/platform/app/ports (interfaces: GRPCServer,HTTPServer,SwaggerServer,ServerFactory)
//
// Generated by this command:
//
//	mockgen -destination=server_mock.go -package=ports github.com/popeskul/awesome-messanger/services/platform/app/ports GRPCServer,HTTPServer,SwaggerServer,ServerFactory
//

// Package ports is a generated GoMock package.
package ports

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockGRPCServer is a mock of GRPCServer interface.
type MockGRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCServerMockRecorder
}

// MockGRPCServerMockRecorder is the mock recorder for MockGRPCServer.
type MockGRPCServerMockRecorder struct {
	mock *MockGRPCServer
}

// NewMockGRPCServer creates a new mock instance.
func NewMockGRPCServer(ctrl *gomock.Controller) *MockGRPCServer {
	mock := &MockGRPCServer{ctrl: ctrl}
	mock.recorder = &MockGRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCServer) EXPECT() *MockGRPCServerMockRecorder {
	return m.recorder
}

// GetGrpcServer mocks base method.
func (m *MockGRPCServer) GetGrpcServer() *grpc.Server {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrpcServer")
	ret0, _ := ret[0].(*grpc.Server)
	return ret0
}

// GetGrpcServer indicates an expected call of GetGrpcServer.
func (mr *MockGRPCServerMockRecorder) GetGrpcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrpcServer", reflect.TypeOf((*MockGRPCServer)(nil).GetGrpcServer))
}

// Start mocks base method.
func (m *MockGRPCServer) Start(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockGRPCServerMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockGRPCServer)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockGRPCServer) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockGRPCServerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGRPCServer)(nil).Stop))
}

// MockHTTPServer is a mock of HTTPServer interface.
type MockHTTPServer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPServerMockRecorder
}

// MockHTTPServerMockRecorder is the mock recorder for MockHTTPServer.
type MockHTTPServerMockRecorder struct {
	mock *MockHTTPServer
}

// NewMockHTTPServer creates a new mock instance.
func NewMockHTTPServer(ctrl *gomock.Controller) *MockHTTPServer {
	mock := &MockHTTPServer{ctrl: ctrl}
	mock.recorder = &MockHTTPServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPServer) EXPECT() *MockHTTPServerMockRecorder {
	return m.recorder
}

// ListenAndServe mocks base method.
func (m *MockHTTPServer) ListenAndServe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockHTTPServerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockHTTPServer)(nil).ListenAndServe))
}

// Shutdown mocks base method.
func (m *MockHTTPServer) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockHTTPServerMockRecorder) Shutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockHTTPServer)(nil).Shutdown), arg0)
}

// MockSwaggerServer is a mock of SwaggerServer interface.
type MockSwaggerServer struct {
	ctrl     *gomock.Controller
	recorder *MockSwaggerServerMockRecorder
}

// MockSwaggerServerMockRecorder is the mock recorder for MockSwaggerServer.
type MockSwaggerServerMockRecorder struct {
	mock *MockSwaggerServer
}

// NewMockSwaggerServer creates a new mock instance.
func NewMockSwaggerServer(ctrl *gomock.Controller) *MockSwaggerServer {
	mock := &MockSwaggerServer{ctrl: ctrl}
	mock.recorder = &MockSwaggerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwaggerServer) EXPECT() *MockSwaggerServerMockRecorder {
	return m.recorder
}

// ListenAndServe mocks base method.
func (m *MockSwaggerServer) ListenAndServe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockSwaggerServerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockSwaggerServer)(nil).ListenAndServe))
}

// Shutdown mocks base method.
func (m *MockSwaggerServer) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockSwaggerServerMockRecorder) Shutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockSwaggerServer)(nil).Shutdown), arg0)
}

// MockServerFactory is a mock of ServerFactory interface.
type MockServerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServerFactoryMockRecorder
}

// MockServerFactoryMockRecorder is the mock recorder for MockServerFactory.
type MockServerFactoryMockRecorder struct {
	mock *MockServerFactory
}

// NewMockServerFactory creates a new mock instance.
func NewMockServerFactory(ctrl *gomock.Controller) *MockServerFactory {
	mock := &MockServerFactory{ctrl: ctrl}
	mock.recorder = &MockServerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerFactory) EXPECT() *MockServerFactoryMockRecorder {
	return m.recorder
}

// NewGRPCServer mocks base method.
func (m *MockServerFactory) NewGRPCServer() (GRPCServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGRPCServer")
	ret0, _ := ret[0].(GRPCServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGRPCServer indicates an expected call of NewGRPCServer.
func (mr *MockServerFactoryMockRecorder) NewGRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGRPCServer", reflect.TypeOf((*MockServerFactory)(nil).NewGRPCServer))
}

// NewHTTPServer mocks base method.
func (m *MockServerFactory) NewHTTPServer() (HTTPServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHTTPServer")
	ret0, _ := ret[0].(HTTPServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHTTPServer indicates an expected call of NewHTTPServer.
func (mr *MockServerFactoryMockRecorder) NewHTTPServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHTTPServer", reflect.TypeOf((*MockServerFactory)(nil).NewHTTPServer))
}

// NewSwaggerServer mocks base method.
func (m *MockServerFactory) NewSwaggerServer() (SwaggerServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSwaggerServer")
	ret0, _ := ret[0].(SwaggerServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSwaggerServer indicates an expected call of NewSwaggerServer.
func (mr *MockServerFactoryMockRecorder) NewSwaggerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSwaggerServer", reflect.TypeOf((*MockServerFactory)(nil).NewSwaggerServer))
}
