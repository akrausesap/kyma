// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

import v1 "k8s.io/api/core/v1"

// addonsCfgUpdater is an autogenerated mock type for the addonsCfgUpdater type
type addonsCfgUpdater struct {
	mock.Mock
}

// AddRepos provides a mock function with given fields: name, url
func (_m *addonsCfgUpdater) AddRepos(name string, url []string) (*v1.ConfigMap, error) {
	ret := _m.Called(name, url)

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func(string, []string) *v1.ConfigMap); ok {
		r0 = rf(name, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(name, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRepos provides a mock function with given fields: name, urls
func (_m *addonsCfgUpdater) RemoveRepos(name string, urls []string) (*v1.ConfigMap, error) {
	ret := _m.Called(name, urls)

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func(string, []string) *v1.ConfigMap); ok {
		r0 = rf(name, urls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(name, urls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}