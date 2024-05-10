// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	dtos "happyBill/dtos"
	models "happyBill/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// CreateBillboard mocks base method.
func (m *MockRepository) CreateBillboard(product models.Product) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillboard", product)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBillboard indicates an expected call of CreateBillboard.
func (mr *MockRepositoryMockRecorder) CreateBillboard(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillboard", reflect.TypeOf((*MockRepository)(nil).CreateBillboard), product)
}

// CreateClient mocks base method.
func (m *MockRepository) CreateClient(client models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", client)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient.
func (mr *MockRepositoryMockRecorder) CreateClient(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockRepository)(nil).CreateClient), client)
}

// CreateManager mocks base method.
func (m *MockRepository) CreateManager(manager models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManager", manager)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateManager indicates an expected call of CreateManager.
func (mr *MockRepositoryMockRecorder) CreateManager(manager interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManager", reflect.TypeOf((*MockRepository)(nil).CreateManager), manager)
}

// CreateOrder mocks base method.
func (m *MockRepository) CreateOrder(clientId int, order models.Order) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", clientId, order)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockRepositoryMockRecorder) CreateOrder(clientId, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockRepository)(nil).CreateOrder), clientId, order)
}

// DeleteBillboard mocks base method.
func (m *MockRepository) DeleteBillboard(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBillboard", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBillboard indicates an expected call of DeleteBillboard.
func (mr *MockRepositoryMockRecorder) DeleteBillboard(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBillboard", reflect.TypeOf((*MockRepository)(nil).DeleteBillboard), id)
}

// GetAllBillboards mocks base method.
func (m *MockRepository) GetAllBillboards(page int) ([]dtos.Product, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBillboards", page)
	ret0, _ := ret[0].([]dtos.Product)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllBillboards indicates an expected call of GetAllBillboards.
func (mr *MockRepositoryMockRecorder) GetAllBillboards(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBillboards", reflect.TypeOf((*MockRepository)(nil).GetAllBillboards), page)
}

// GetAllManagerOrders mocks base method.
func (m *MockRepository) GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllManagerOrders", id, page)
	ret0, _ := ret[0].([]dtos.ManagerOrder)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllManagerOrders indicates an expected call of GetAllManagerOrders.
func (mr *MockRepositoryMockRecorder) GetAllManagerOrders(id, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllManagerOrders", reflect.TypeOf((*MockRepository)(nil).GetAllManagerOrders), id, page)
}

// GetAllManagers mocks base method.
func (m *MockRepository) GetAllManagers(page int) ([]dtos.User, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllManagers", page)
	ret0, _ := ret[0].([]dtos.User)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllManagers indicates an expected call of GetAllManagers.
func (mr *MockRepositoryMockRecorder) GetAllManagers(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllManagers", reflect.TypeOf((*MockRepository)(nil).GetAllManagers), page)
}

// GetAllOrders mocks base method.
func (m *MockRepository) GetAllOrders() ([]dtos.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders")
	ret0, _ := ret[0].([]dtos.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockRepositoryMockRecorder) GetAllOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockRepository)(nil).GetAllOrders))
}

// GetAllSearchedBillboards mocks base method.
func (m *MockRepository) GetAllSearchedBillboards(page int, search dtos.Search, filter dtos.Filter) ([]dtos.Product, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSearchedBillboards", page, search, filter)
	ret0, _ := ret[0].([]dtos.Product)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllSearchedBillboards indicates an expected call of GetAllSearchedBillboards.
func (mr *MockRepositoryMockRecorder) GetAllSearchedBillboards(page, search, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSearchedBillboards", reflect.TypeOf((*MockRepository)(nil).GetAllSearchedBillboards), page, search, filter)
}

// GetAllSearchedBillboardsFake mocks base method.
func (m *MockRepository) GetAllSearchedBillboardsFake(filter dtos.Filter) ([]dtos.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSearchedBillboardsFake", filter)
	ret0, _ := ret[0].([]dtos.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSearchedBillboardsFake indicates an expected call of GetAllSearchedBillboardsFake.
func (mr *MockRepositoryMockRecorder) GetAllSearchedBillboardsFake(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSearchedBillboardsFake", reflect.TypeOf((*MockRepository)(nil).GetAllSearchedBillboardsFake), filter)
}

// GetBillboardById mocks base method.
func (m *MockRepository) GetBillboardById(id int) (dtos.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillboardById", id)
	ret0, _ := ret[0].(dtos.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillboardById indicates an expected call of GetBillboardById.
func (mr *MockRepositoryMockRecorder) GetBillboardById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillboardById", reflect.TypeOf((*MockRepository)(nil).GetBillboardById), id)
}

// GetClientById mocks base method.
func (m *MockRepository) GetClientById(id int) (dtos.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientById", id)
	ret0, _ := ret[0].(dtos.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientById indicates an expected call of GetClientById.
func (mr *MockRepositoryMockRecorder) GetClientById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientById", reflect.TypeOf((*MockRepository)(nil).GetClientById), id)
}

// GetClientByUserId mocks base method.
func (m *MockRepository) GetClientByUserId(id int) (dtos.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientByUserId", id)
	ret0, _ := ret[0].(dtos.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientByUserId indicates an expected call of GetClientByUserId.
func (mr *MockRepositoryMockRecorder) GetClientByUserId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientByUserId", reflect.TypeOf((*MockRepository)(nil).GetClientByUserId), id)
}

// GetManagerById mocks base method.
func (m *MockRepository) GetManagerById(id int) (dtos.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagerById", id)
	ret0, _ := ret[0].(dtos.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagerById indicates an expected call of GetManagerById.
func (mr *MockRepositoryMockRecorder) GetManagerById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagerById", reflect.TypeOf((*MockRepository)(nil).GetManagerById), id)
}

// GetManagerOrderById mocks base method.
func (m *MockRepository) GetManagerOrderById(id int) (dtos.ManagerOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagerOrderById", id)
	ret0, _ := ret[0].(dtos.ManagerOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagerOrderById indicates an expected call of GetManagerOrderById.
func (mr *MockRepositoryMockRecorder) GetManagerOrderById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagerOrderById", reflect.TypeOf((*MockRepository)(nil).GetManagerOrderById), id)
}

// GetMostFreeManager mocks base method.
func (m *MockRepository) GetMostFreeManager() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMostFreeManager")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMostFreeManager indicates an expected call of GetMostFreeManager.
func (mr *MockRepositoryMockRecorder) GetMostFreeManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMostFreeManager", reflect.TypeOf((*MockRepository)(nil).GetMostFreeManager))
}

// GetMyBillboards mocks base method.
func (m *MockRepository) GetMyBillboards(id, page int) ([]dtos.Product, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyBillboards", id, page)
	ret0, _ := ret[0].([]dtos.Product)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMyBillboards indicates an expected call of GetMyBillboards.
func (mr *MockRepositoryMockRecorder) GetMyBillboards(id, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyBillboards", reflect.TypeOf((*MockRepository)(nil).GetMyBillboards), id, page)
}

// GetMyOrders mocks base method.
func (m *MockRepository) GetMyOrders(clientId, page int, status string) ([]dtos.MyOrder, dtos.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyOrders", clientId, page, status)
	ret0, _ := ret[0].([]dtos.MyOrder)
	ret1, _ := ret[1].(dtos.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMyOrders indicates an expected call of GetMyOrders.
func (mr *MockRepositoryMockRecorder) GetMyOrders(clientId, page, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyOrders", reflect.TypeOf((*MockRepository)(nil).GetMyOrders), clientId, page, status)
}

// GetRoleId mocks base method.
func (m *MockRepository) GetRoleId(role string, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleId", role, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleId indicates an expected call of GetRoleId.
func (mr *MockRepositoryMockRecorder) GetRoleId(role, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleId", reflect.TypeOf((*MockRepository)(nil).GetRoleId), role, userId)
}

// GetRoles mocks base method.
func (m *MockRepository) GetRoles(id int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles.
func (mr *MockRepositoryMockRecorder) GetRoles(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockRepository)(nil).GetRoles), id)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), username)
}

// GetUserById mocks base method.
func (m *MockRepository) GetUserById(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockRepositoryMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockRepository)(nil).GetUserById), id)
}

// LikeBillboard mocks base method.
func (m *MockRepository) LikeBillboard(clientId, productId int, action string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeBillboard", clientId, productId, action)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikeBillboard indicates an expected call of LikeBillboard.
func (mr *MockRepositoryMockRecorder) LikeBillboard(clientId, productId, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeBillboard", reflect.TypeOf((*MockRepository)(nil).LikeBillboard), clientId, productId, action)
}

// UpdateBillboard mocks base method.
func (m *MockRepository) UpdateBillboard(id int, input dtos.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBillboard", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBillboard indicates an expected call of UpdateBillboard.
func (mr *MockRepositoryMockRecorder) UpdateBillboard(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBillboard", reflect.TypeOf((*MockRepository)(nil).UpdateBillboard), id, input)
}

// UpdateManagerOrder mocks base method.
func (m *MockRepository) UpdateManagerOrder(id int, input dtos.UpdateOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateManagerOrder", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateManagerOrder indicates an expected call of UpdateManagerOrder.
func (mr *MockRepositoryMockRecorder) UpdateManagerOrder(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateManagerOrder", reflect.TypeOf((*MockRepository)(nil).UpdateManagerOrder), id, input)
}

// UpdateMyProfile mocks base method.
func (m *MockRepository) UpdateMyProfile(userId int, input dtos.UpdateUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMyProfile", userId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMyProfile indicates an expected call of UpdateMyProfile.
func (mr *MockRepositoryMockRecorder) UpdateMyProfile(userId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMyProfile", reflect.TypeOf((*MockRepository)(nil).UpdateMyProfile), userId, input)
}
