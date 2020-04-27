package controller

import (
	"encoding/json"
	"friend/dto"
	"friend/response"
	"friend/service"
	"net/http"
)

type FriendController struct {
	FriendService service.IFriendService
}

// @Summary CreateFriend
// @Description API to create friend connection between two email addresses
// @Tags Friend
// @Accept json
// @Produce json
// @Param friendDto body dto.FriendDto true "friendDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /friends/create-friend [post]
func (f FriendController) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateFriend(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// @Summary CreateSubscribe
// @Description API to subscribe to updates from an email address
// @Tags Friend
// @Accept json
// @Produce json
// @Param requestDto body dto.RequestDto true "requestDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /friends/subscribe [post]
func (f FriendController) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateSubscribe(requestDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// @Summary CreateBlock
// @Description API to block updates from an email address
// @Tags Friend
// @Accept json
// @Produce json
// @Param requestDto body dto.RequestDto true "requestDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /friends/block [post]
func (f FriendController) CreateBlock(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateBlock(requestDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// @Summary GetFriendsListByEmail
// @Description API to retrieve the friends list for an email address
// @Tags Friend
// @Accept json
// @Produce json
// @Param emailDto body dto.EmailDto true "emailDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Router /friends/get-friends-list [post]
func (f FriendController) GetFriendsListByEmail(w http.ResponseWriter, r *http.Request) {
	var emailDto dto.EmailDto
	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	results, err := f.FriendService.GetFriendsListByEmail(emailDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

// @Summary GetCommonFriends
// @Description API to retrieve the common friends list between two email addresses
// @Tags Friend
// @Accept json
// @Produce json
// @Param friendDto body dto.FriendDto true "friendDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Router /friends/get-common-friends-list [post]
func (f FriendController) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	results, err := f.FriendService.GetCommonFriends(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

// @Summary GetReceiversList
// @Description API to retrieve all email addresses that can receive updates from an email address
// @Tags Friend
// @Accept json
// @Produce json
// @Param senderDto body dto.SenderDto true "senderDto"
// @success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Router /friends/get-receivers-list [post]
func (f FriendController) GetReceiversList(w http.ResponseWriter, r *http.Request) {
	var senderDto dto.SenderDto
	if err := json.NewDecoder(r.Body).Decode(&senderDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	results, err := f.FriendService.GetReceiversList(senderDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}