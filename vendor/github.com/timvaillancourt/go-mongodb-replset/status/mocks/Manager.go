// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import status "github.com/timvaillancourt/go-mongodb-replset/status"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// GetMember provides a mock function with given fields: name
func (_m *Manager) GetMember(name string) *status.Member {
	ret := _m.Called(name)

	var r0 *status.Member
	if rf, ok := ret.Get(0).(func(string) *status.Member); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*status.Member)
		}
	}

	return r0
}

// GetMemberId provides a mock function with given fields: id
func (_m *Manager) GetMemberId(id int) *status.Member {
	ret := _m.Called(id)

	var r0 *status.Member
	if rf, ok := ret.Get(0).(func(int) *status.Member); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*status.Member)
		}
	}

	return r0
}

// GetMembersByState provides a mock function with given fields: state, limit
func (_m *Manager) GetMembersByState(state status.MemberState, limit int) []*status.Member {
	ret := _m.Called(state, limit)

	var r0 []*status.Member
	if rf, ok := ret.Get(0).(func(status.MemberState, int) []*status.Member); ok {
		r0 = rf(state, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*status.Member)
		}
	}

	return r0
}

// GetSelf provides a mock function with given fields:
func (_m *Manager) GetSelf() *status.Member {
	ret := _m.Called()

	var r0 *status.Member
	if rf, ok := ret.Get(0).(func() *status.Member); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*status.Member)
		}
	}

	return r0
}

// HasMember provides a mock function with given fields: name
func (_m *Manager) HasMember(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Primary provides a mock function with given fields:
func (_m *Manager) Primary() *status.Member {
	ret := _m.Called()

	var r0 *status.Member
	if rf, ok := ret.Get(0).(func() *status.Member); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*status.Member)
		}
	}

	return r0
}

// Secondaries provides a mock function with given fields:
func (_m *Manager) Secondaries() []*status.Member {
	ret := _m.Called()

	var r0 []*status.Member
	if rf, ok := ret.Get(0).(func() []*status.Member); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*status.Member)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Manager) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ToJSON provides a mock function with given fields:
func (_m *Manager) ToJSON() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}