package impl

import (
	"friend/dto"
	"friend/entity"
	"friend/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	listEmails := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}

	userRepositoryMock := repository.UserRepositoryMock{}
	userRepositoryMock.On("GetAllUsers").Return(listEmails)
	userService := UserServiceImpl{userRepositoryMock}

	assert.Equal(t, listEmails, userService.GetAllUsers())
	userRepositoryMock.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	emailDto := dto.EmailDto{Email: "user1@gmail.com"}
	email := emailDto.Email

	userRepositoryMock := repository.UserRepositoryMock{}
	userRepositoryMock.On("ExistsByEmail", email).Return(false)
	userRepositoryMock.On("CreateUser", email).Return(true)
	userService := UserServiceImpl{userRepositoryMock}

	result, _ := userService.CreateUser(emailDto)
	assert.Equal(t, true, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestExistsByEmail(t *testing.T) {
	email := "a@gmail.com"

	userRepositoryMock := repository.UserRepositoryMock{}
	userRepositoryMock.On("ExistsByEmail", email).Return(false)
	userService := UserServiceImpl{userRepositoryMock}

	assert.Equal(t, false, userService.ExistsByEmail(email))
	userRepositoryMock.AssertExpectations(t)
}

func TestFindUserIdByEmail(t *testing.T) {
	user := entity.User{Id: 1, Email: "a@gmail.com"}

	userRepositoryMock := repository.UserRepositoryMock{}
	userRepositoryMock.On("FindUserIdByEmail", user.Email).Return(user.Id)
	userService := UserServiceImpl{userRepositoryMock}

	result, _ := userService.FindUserIdByEmail(user.Email)

	assert.Equal(t, user.Id, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestFindByIds(t *testing.T) {
	listIds := []int64{1,2,3}
	listEmails := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}

	userRepositoryMock := repository.UserRepositoryMock{}
	userRepositoryMock.On("FindByIds", listIds).Return(listEmails)
	userService := UserServiceImpl{userRepositoryMock}

	assert.Equal(t, listEmails, userService.FindByIds(listIds))
	userRepositoryMock.AssertExpectations(t)
}