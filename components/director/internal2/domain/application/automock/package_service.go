// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal2/model"

// PackageService is an autogenerated mock type for the PackageService type
type PackageService struct {
	mock.Mock
}

// CreateMultiple provides a mock function with given fields: ctx, applicationID, in
func (_m *PackageService) CreateMultiple(ctx context.Context, applicationID string, in []*model.PackageCreateInput) error {
	ret := _m.Called(ctx, applicationID, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []*model.PackageCreateInput) error); ok {
		r0 = rf(ctx, applicationID, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetForApplication provides a mock function with given fields: ctx, id, applicationID
func (_m *PackageService) GetForApplication(ctx context.Context, id string, applicationID string) (*model.Package, error) {
	ret := _m.Called(ctx, id, applicationID)

	var r0 *model.Package
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Package); ok {
		r0 = rf(ctx, id, applicationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, applicationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByApplicationID provides a mock function with given fields: ctx, applicationID, pageSize, cursor
func (_m *PackageService) ListByApplicationID(ctx context.Context, applicationID string, pageSize int, cursor string) (*model.PackagePage, error) {
	ret := _m.Called(ctx, applicationID, pageSize, cursor)

	var r0 *model.PackagePage
	if rf, ok := ret.Get(0).(func(context.Context, string, int, string) *model.PackagePage); ok {
		r0 = rf(ctx, applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PackagePage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, string) error); ok {
		r1 = rf(ctx, applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}