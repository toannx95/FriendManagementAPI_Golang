package controller

import (
	"friend/response"
	"friend/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func (u UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	friends := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: friends, Count: len(friends)}
	response.SuccessResponse(w, res)
}

func (u UserController)  CreateUser(w http.ResponseWriter, r *http.Request) {
	result, err := u.UserService.CreateUser(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}