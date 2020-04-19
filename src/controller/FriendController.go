package controller

import (
	"net/http"
	"response"
	"service"
)

type FriendController struct {
	FriendService service.FriendService
}

func (f FriendController) CreateUser(w http.ResponseWriter, r *http.Request) {
	result, err := f.FriendService.CreateFriend(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	result, err := f.FriendService.CreateSubscribe(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) CreateBlock(w http.ResponseWriter, r *http.Request) {
	result, err := f.FriendService.CreateBlock(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetFriendsListByEmail(w http.ResponseWriter, r *http.Request) {
	results, err := f.FriendService.GetFriendsListByEmail(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	results, err := f.FriendService.GetCommonFriends(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetReceiversList(w http.ResponseWriter, r *http.Request) {
	results, err := f.FriendService.GetReceiversList(r)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}