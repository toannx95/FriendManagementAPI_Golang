package repository

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	mock.Mock
}

func (u UserRepositoryMock) GetAllUsers() []string {
	returnVals := u.Called()
	return returnVals.Get(0).([]string)
}

func (u UserRepositoryMock) CreateUser(email string) bool {
	returnVals := u.Called(email)
	return returnVals.Get(0).(bool)
}

func (u UserRepositoryMock) ExistsByEmail(email string) bool {
	returnVals := u.Called(email)
	return returnVals.Get(0).(bool)
}

func (u UserRepositoryMock) FindUserIdByEmail(email string) int64 {
	returnVals := u.Called(email)
	return returnVals.Get(0).(int64)
}

func (u UserRepositoryMock) FindByIds(ids []int64) []string {
	returnVals := u.Called(ids)
	return returnVals.Get(0).([]string)
}
