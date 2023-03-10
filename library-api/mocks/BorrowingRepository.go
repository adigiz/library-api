// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "library-api/models"
	mock "github.com/stretchr/testify/mock"
)

// BorrowingRepository is an autogenerated mock type for the BorrowingRepository type
type BorrowingRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: b
func (_m *BorrowingRepository) Create(b models.BorrowingRecord) (*models.BorrowingRecord, error) {
	ret := _m.Called(b)

	var r0 *models.BorrowingRecord
	if rf, ok := ret.Get(0).(func(models.BorrowingRecord) *models.BorrowingRecord); ok {
		r0 = rf(b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BorrowingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.BorrowingRecord) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnBook provides a mock function with given fields: recordID
func (_m *BorrowingRepository) ReturnBook(recordID int) (*models.BorrowingRecord, error) {
	ret := _m.Called(recordID)

	var r0 *models.BorrowingRecord
	if rf, ok := ret.Get(0).(func(int) *models.BorrowingRecord); ok {
		r0 = rf(recordID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BorrowingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(recordID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
