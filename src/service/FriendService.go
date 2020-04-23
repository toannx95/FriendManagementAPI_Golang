package service

import (
	"friend/dto"
	ex "friend/exception"
)

type FriendService interface {
	CreateFriend(friendDto dto.FriendDto) (bool, *ex.Exception)
	CreateSubscribe(requestDto dto.RequestDto) (bool, *ex.Exception)
	CreateBlock(requestDto dto.RequestDto) (bool, *ex.Exception)
	GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *ex.Exception)
	GetCommonFriends(friendDto dto.FriendDto) ([]string, *ex.Exception)
	GetReceiversList(senderDto dto.SenderDto) ([]string, *ex.Exception)
}