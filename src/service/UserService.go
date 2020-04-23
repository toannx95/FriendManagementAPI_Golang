package service

import (
	"friend/dto"
	exp "friend/exception"
)

type UserService interface {
	GetAllUsers() []string
	CreateUser(emailDto dto.EmailDto) (bool, *exp.Exception)
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) (int64, error)
	FindByIds(ids []int64) []string
}