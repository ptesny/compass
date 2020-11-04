// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal2/model"
import packageinstanceauth "github.com/kyma-incubator/compass/components/director/internal2/domain/packageinstanceauth"

// EntityConverter is an autogenerated mock type for the EntityConverter type
type EntityConverter struct {
	mock.Mock
}

// FromEntity provides a mock function with given fields: entity
func (_m *EntityConverter) FromEntity(entity packageinstanceauth.Entity) (model.PackageInstanceAuth, error) {
	ret := _m.Called(entity)

	var r0 model.PackageInstanceAuth
	if rf, ok := ret.Get(0).(func(packageinstanceauth.Entity) model.PackageInstanceAuth); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(model.PackageInstanceAuth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(packageinstanceauth.Entity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToEntity provides a mock function with given fields: in
func (_m *EntityConverter) ToEntity(in model.PackageInstanceAuth) (packageinstanceauth.Entity, error) {
	ret := _m.Called(in)

	var r0 packageinstanceauth.Entity
	if rf, ok := ret.Get(0).(func(model.PackageInstanceAuth) packageinstanceauth.Entity); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(packageinstanceauth.Entity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.PackageInstanceAuth) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}