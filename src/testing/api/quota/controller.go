// Code generated by mockery v1.0.0. DO NOT EDIT.

package quota

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/quota/models"
	mock "github.com/stretchr/testify/mock"

	quota "github.com/goharbor/harbor/src/api/quota"

	types "github.com/goharbor/harbor/src/pkg/types"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, reference, referenceID, hardLimits, used
func (_m *Controller) Create(ctx context.Context, reference string, referenceID string, hardLimits types.ResourceList, used ...types.ResourceList) (int64, error) {
	_va := make([]interface{}, len(used))
	for _i := range used {
		_va[_i] = used[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reference, referenceID, hardLimits)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, types.ResourceList, ...types.ResourceList) int64); ok {
		r0 = rf(ctx, reference, referenceID, hardLimits, used...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, types.ResourceList, ...types.ResourceList) error); ok {
		r1 = rf(ctx, reference, referenceID, hardLimits, used...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Controller) Get(ctx context.Context, id int64) (*models.Quota, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Quota
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Quota); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRef provides a mock function with given fields: ctx, reference, referenceID
func (_m *Controller) GetByRef(ctx context.Context, reference string, referenceID string) (*models.Quota, error) {
	ret := _m.Called(ctx, reference, referenceID)

	var r0 *models.Quota
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Quota); ok {
		r0 = rf(ctx, reference, referenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, reference, referenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEnabled provides a mock function with given fields: ctx, reference, referenceID
func (_m *Controller) IsEnabled(ctx context.Context, reference string, referenceID string) (bool, error) {
	ret := _m.Called(ctx, reference, referenceID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, reference, referenceID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, reference, referenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx, reference, referenceID, options
func (_m *Controller) Refresh(ctx context.Context, reference string, referenceID string, options ...quota.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reference, referenceID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...quota.Option) error); ok {
		r0 = rf(ctx, reference, referenceID, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Request provides a mock function with given fields: ctx, reference, referenceID, resources, f
func (_m *Controller) Request(ctx context.Context, reference string, referenceID string, resources types.ResourceList, f func() error) error {
	ret := _m.Called(ctx, reference, referenceID, resources, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, types.ResourceList, func() error) error); ok {
		r0 = rf(ctx, reference, referenceID, resources, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
