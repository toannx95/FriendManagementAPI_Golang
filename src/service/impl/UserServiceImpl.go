package impl

import (
	"dto"
	"encoding/json"
	"errors"
	exp "exception"
	"net/http"
	"repository"
	"utils"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (f UserServiceImpl) GetAllUsers() []string {
	return f.UserRepository.GetAllUsers()
}

func (f UserServiceImpl) CreateUser(r *http.Request) (bool, *exp.Exception) {
	emailDto := dto.EmailDto{}

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		return false, &exp.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	if !utils.IsFormatEmail(emailDto.Email) {
		return false, &exp.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	if f.UserRepository.ExistsByEmail(emailDto.Email) {
		return false, &exp.Exception{Code: http.StatusInternalServerError, Message: "Email already existed!"}
	}

	user := f.UserRepository.CreateUser(emailDto.Email)
	if user != true {
		return false, &exp.Exception{Code: http.StatusInternalServerError, Message: "Error when create user!"}
	}
	return true, nil
}

func (f UserServiceImpl) ExistsByEmail(email string) bool {
	return f.UserRepository.ExistsByEmail(email)
}

func (f UserServiceImpl) FindUserIdByEmail(email string) (int64, error) {
	id := f.UserRepository.FindUserIdByEmail(email)
	if id == -1 {
		return -1, errors.New("User not found!")
	}
	return id, nil
}

func (f UserServiceImpl) FindByIds(ids []int64) []string {
	return f.UserRepository.FindByIds(ids)
}

