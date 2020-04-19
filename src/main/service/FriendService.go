package service

import (
	ex "main/exception"
	"net/http"
)

type FriendService interface {
	CreateFriend(r *http.Request) (bool, *ex.Exception)
	CreateSubscribe(r *http.Request) (bool, *ex.Exception)
	CreateBlock(r *http.Request) (bool, *ex.Exception)
	GetFriendsListByEmail(r *http.Request) ([]string, *ex.Exception)
	GetCommonFriends(r *http.Request) ([]string, *ex.Exception)
	GetReceiversList(r *http.Request) ([]string, *ex.Exception)
}