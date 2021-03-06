// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/apperrors"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/kyma/apiresources/secrets/model"

import types "k8s.io/apimachinery/pkg/types"

// RequestParametersService is an autogenerated mock type for the RequestParametersService type
type RequestParametersService struct {
	mock.Mock
}

// Create provides a mock function with given fields: application, appUID, serviceID, requestParameters
func (_m *RequestParametersService) Create(application string, appUID types.UID, serviceID string, requestParameters *model.RequestParameters) (string, apperrors.AppError) {
	ret := _m.Called(application, appUID, serviceID, requestParameters)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, types.UID, string, *model.RequestParameters) string); ok {
		r0 = rf(application, appUID, serviceID, requestParameters)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, types.UID, string, *model.RequestParameters) apperrors.AppError); ok {
		r1 = rf(application, appUID, serviceID, requestParameters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: application, serviceId
func (_m *RequestParametersService) Delete(application string, serviceId string) apperrors.AppError {
	ret := _m.Called(application, serviceId)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(application, serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Get provides a mock function with given fields: secretName
func (_m *RequestParametersService) Get(secretName string) (model.RequestParameters, apperrors.AppError) {
	ret := _m.Called(secretName)

	var r0 model.RequestParameters
	if rf, ok := ret.Get(0).(func(string) model.RequestParameters); ok {
		r0 = rf(secretName)
	} else {
		r0 = ret.Get(0).(model.RequestParameters)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(secretName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: application, appUID, serviceID, requestParameters
func (_m *RequestParametersService) Upsert(application string, appUID types.UID, serviceID string, requestParameters *model.RequestParameters) (string, apperrors.AppError) {
	ret := _m.Called(application, appUID, serviceID, requestParameters)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, types.UID, string, *model.RequestParameters) string); ok {
		r0 = rf(application, appUID, serviceID, requestParameters)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, types.UID, string, *model.RequestParameters) apperrors.AppError); ok {
		r1 = rf(application, appUID, serviceID, requestParameters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
