// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WorkspaceDestroyer is an autogenerated mock type for the WorkspaceDestroyer type
type WorkspaceDestroyer struct {
	mock.Mock
}

// DestroyWorkspaceFunc provides a mock function with given fields:
func (_m *WorkspaceDestroyer) DestroyWorkspaceFunc() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}