// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/trainers/library-api/models"
	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// DecreaseStock provides a mock function with given fields: id
func (_m *BookRepository) DecreaseStock(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindBook provides a mock function with given fields: id
func (_m *BookRepository) FindBook(id int) (*models.Book, error) {
	ret := _m.Called(id)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(int) *models.Book); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBooks provides a mock function with given fields:
func (_m *BookRepository) FindBooks() ([]*models.Book, error) {
	ret := _m.Called()

	var r0 []*models.Book
	if rf, ok := ret.Get(0).(func() []*models.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Book)
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

// IncreaseStock provides a mock function with given fields: id
func (_m *BookRepository) IncreaseStock(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: book
func (_m *BookRepository) Save(book *models.Book) (*models.Book, int, error) {
	ret := _m.Called(book)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(*models.Book) *models.Book); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*models.Book) int); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.Book) error); ok {
		r2 = rf(book)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
