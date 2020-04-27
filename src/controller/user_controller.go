package controller

import (
	"encoding/json"
	"friend/dto"
	"friend/response"
	"friend/service"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
}

// @Summary GetAllUsers
// @Description API to get all emails
// @Tags User
// @Accept json
// @Produce json
// @success 200 {object} response.Response
// @Router /users [get]
func (u UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	friends := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: friends, Count: len(friends)}
	response.SuccessResponse(w, res)
}

// @Summary CreateUser
// @Description API to create new user
// @Tags User
// @Accept json
// @Produce json
// @Param emailDto body dto.EmailDto true "emailDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /users/create-user [post]
func (u UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	emailDto := dto.EmailDto{}

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := u.UserService.CreateUser(emailDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}