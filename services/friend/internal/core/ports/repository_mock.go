// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/popeskul/awesome-messanger/services/friend/internal/core/ports (interfaces: Repository,FriendRepository)
//
// Generated by this command:
//
//	mockgen -destination=repository_mock.go -package=ports github.com/popeskul/awesome-messanger/services/friend/internal/core/ports Repository,FriendRepository
//

// Package ports is a generated GoMock package.
package ports

import (
	context "context"
	reflect "reflect"

	models "github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Friend mocks base method.
func (m *MockRepository) Friend() FriendRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Friend")
	ret0, _ := ret[0].(FriendRepository)
	return ret0
}

// Friend indicates an expected call of Friend.
func (mr *MockRepositoryMockRecorder) Friend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Friend", reflect.TypeOf((*MockRepository)(nil).Friend))
}

// MockFriendRepository is a mock of FriendRepository interface.
type MockFriendRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFriendRepositoryMockRecorder
}

// MockFriendRepositoryMockRecorder is the mock recorder for MockFriendRepository.
type MockFriendRepositoryMockRecorder struct {
	mock *MockFriendRepository
}

// NewMockFriendRepository creates a new mock instance.
func NewMockFriendRepository(ctrl *gomock.Controller) *MockFriendRepository {
	mock := &MockFriendRepository{ctrl: ctrl}
	mock.recorder = &MockFriendRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFriendRepository) EXPECT() *MockFriendRepositoryMockRecorder {
	return m.recorder
}

// AddFriend mocks base method.
func (m *MockFriendRepository) AddFriend(arg0 context.Context, arg1 *models.Friend) (*models.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(*models.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFriend indicates an expected call of AddFriend.
func (mr *MockFriendRepositoryMockRecorder) AddFriend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockFriendRepository)(nil).AddFriend), arg0, arg1)
}

// GetFriends mocks base method.
func (m *MockFriendRepository) GetFriends(arg0 context.Context, arg1 string) ([]*models.Friend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriends", arg0, arg1)
	ret0, _ := ret[0].([]*models.Friend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriends indicates an expected call of GetFriends.
func (mr *MockFriendRepositoryMockRecorder) GetFriends(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriends", reflect.TypeOf((*MockFriendRepository)(nil).GetFriends), arg0, arg1)
}

// RemoveFriend mocks base method.
func (m *MockFriendRepository) RemoveFriend(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFriend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFriend indicates an expected call of RemoveFriend.
func (mr *MockFriendRepositoryMockRecorder) RemoveFriend(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFriend", reflect.TypeOf((*MockFriendRepository)(nil).RemoveFriend), arg0, arg1, arg2)
}
