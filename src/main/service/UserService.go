package service

import (
	exp "main/exception"
	"net/http"
)

type UserService interface {
	GetAllUsers() []string
	CreateUser(r *http.Request) (bool, *exp.Exception)
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) (int64, error)
	FindByIds(ids []int64) []string
}