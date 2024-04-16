// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// JwtInterface is an autogenerated mock type for the JwtInterface type
type JwtInterface struct {
	mock.Mock
}

// DecodeRole provides a mock function with given fields: token
func (_m *JwtInterface) DecodeRole(token *jwt.Token) string {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for DecodeRole")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(*jwt.Token) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DecodeToken provides a mock function with given fields: token
func (_m *JwtInterface) DecodeToken(token *jwt.Token) string {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for DecodeToken")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(*jwt.Token) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateJWT provides a mock function with given fields: email, role
func (_m *JwtInterface) GenerateJWT(email string, role string) (string, error) {
	ret := _m.Called(email, role)

	if len(ret) == 0 {
		panic("no return value specified for GenerateJWT")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(email, role)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJwtInterface creates a new instance of JwtInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtInterface {
	mock := &JwtInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
