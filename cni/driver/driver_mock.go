// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/cni/driver (interfaces: NetLink,NS,IP)

// Package driver is a generated GoMock package.
package driver

import (
	ns "github.com/containernetworking/cni/pkg/ns"
	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
	net "net"
	reflect "reflect"
)

// MockNetLink is a mock of NetLink interface
type MockNetLink struct {
	ctrl     *gomock.Controller
	recorder *MockNetLinkMockRecorder
}

// MockNetLinkMockRecorder is the mock recorder for MockNetLink
type MockNetLinkMockRecorder struct {
	mock *MockNetLink
}

// NewMockNetLink creates a new mock instance
func NewMockNetLink(ctrl *gomock.Controller) *MockNetLink {
	mock := &MockNetLink{ctrl: ctrl}
	mock.recorder = &MockNetLinkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetLink) EXPECT() *MockNetLinkMockRecorder {
	return m.recorder
}

// AddrAdd mocks base method
func (m *MockNetLink) AddrAdd(arg0 netlink.Link, arg1 *netlink.Addr) error {
	ret := m.ctrl.Call(m, "AddrAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd
func (mr *MockNetLinkMockRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockNetLink)(nil).AddrAdd), arg0, arg1)
}

// LinkAdd mocks base method
func (m *MockNetLink) LinkAdd(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkAdd indicates an expected call of LinkAdd
func (mr *MockNetLinkMockRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkAdd", reflect.TypeOf((*MockNetLink)(nil).LinkAdd), arg0)
}

// LinkByName mocks base method
func (m *MockNetLink) LinkByName(arg0 string) (netlink.Link, error) {
	ret := m.ctrl.Call(m, "LinkByName", arg0)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByName indicates an expected call of LinkByName
func (mr *MockNetLinkMockRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByName", reflect.TypeOf((*MockNetLink)(nil).LinkByName), arg0)
}

// LinkDel mocks base method
func (m *MockNetLink) LinkDel(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDel indicates an expected call of LinkDel
func (mr *MockNetLinkMockRecorder) LinkDel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDel", reflect.TypeOf((*MockNetLink)(nil).LinkDel), arg0)
}

// LinkSetNsFd mocks base method
func (m *MockNetLink) LinkSetNsFd(arg0 netlink.Link, arg1 int) error {
	ret := m.ctrl.Call(m, "LinkSetNsFd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetNsFd indicates an expected call of LinkSetNsFd
func (mr *MockNetLinkMockRecorder) LinkSetNsFd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetNsFd", reflect.TypeOf((*MockNetLink)(nil).LinkSetNsFd), arg0, arg1)
}

// LinkSetUp mocks base method
func (m *MockNetLink) LinkSetUp(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkSetUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp
func (mr *MockNetLinkMockRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockNetLink)(nil).LinkSetUp), arg0)
}

// NeighAdd mocks base method
func (m *MockNetLink) NeighAdd(arg0 *netlink.Neigh) error {
	ret := m.ctrl.Call(m, "NeighAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NeighAdd indicates an expected call of NeighAdd
func (mr *MockNetLinkMockRecorder) NeighAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighAdd", reflect.TypeOf((*MockNetLink)(nil).NeighAdd), arg0)
}

// RouteAdd mocks base method
func (m *MockNetLink) RouteAdd(arg0 *netlink.Route) error {
	ret := m.ctrl.Call(m, "RouteAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteAdd indicates an expected call of RouteAdd
func (mr *MockNetLinkMockRecorder) RouteAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteAdd", reflect.TypeOf((*MockNetLink)(nil).RouteAdd), arg0)
}

// RuleAdd mocks base method
func (m *MockNetLink) RuleAdd(arg0 *netlink.Rule) error {
	ret := m.ctrl.Call(m, "RuleAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd
func (mr *MockNetLinkMockRecorder) RuleAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockNetLink)(nil).RuleAdd), arg0)
}

// RuleDel mocks base method
func (m *MockNetLink) RuleDel(arg0 *netlink.Rule) error {
	ret := m.ctrl.Call(m, "RuleDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleDel indicates an expected call of RuleDel
func (mr *MockNetLinkMockRecorder) RuleDel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleDel", reflect.TypeOf((*MockNetLink)(nil).RuleDel), arg0)
}

// MockNS is a mock of NS interface
type MockNS struct {
	ctrl     *gomock.Controller
	recorder *MockNSMockRecorder
}

// MockNSMockRecorder is the mock recorder for MockNS
type MockNSMockRecorder struct {
	mock *MockNS
}

// NewMockNS creates a new mock instance
func NewMockNS(ctrl *gomock.Controller) *MockNS {
	mock := &MockNS{ctrl: ctrl}
	mock.recorder = &MockNSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNS) EXPECT() *MockNSMockRecorder {
	return m.recorder
}

// WithNetNSPath mocks base method
func (m *MockNS) WithNetNSPath(arg0 string, arg1 func(ns.NetNS) error) error {
	ret := m.ctrl.Call(m, "WithNetNSPath", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithNetNSPath indicates an expected call of WithNetNSPath
func (mr *MockNSMockRecorder) WithNetNSPath(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithNetNSPath", reflect.TypeOf((*MockNS)(nil).WithNetNSPath), arg0, arg1)
}

// MockIP is a mock of IP interface
type MockIP struct {
	ctrl     *gomock.Controller
	recorder *MockIPMockRecorder
}

// MockIPMockRecorder is the mock recorder for MockIP
type MockIPMockRecorder struct {
	mock *MockIP
}

// NewMockIP creates a new mock instance
func NewMockIP(ctrl *gomock.Controller) *MockIP {
	mock := &MockIP{ctrl: ctrl}
	mock.recorder = &MockIPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIP) EXPECT() *MockIPMockRecorder {
	return m.recorder
}

// AddDefaultRoute mocks base method
func (m *MockIP) AddDefaultRoute(arg0 net.IP, arg1 netlink.Link) error {
	ret := m.ctrl.Call(m, "AddDefaultRoute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDefaultRoute indicates an expected call of AddDefaultRoute
func (mr *MockIPMockRecorder) AddDefaultRoute(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDefaultRoute", reflect.TypeOf((*MockIP)(nil).AddDefaultRoute), arg0, arg1)
}
