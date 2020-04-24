package service

import (
	"friend/dto"
	"friend/entity"
	"friend/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFriend(t *testing.T) {
	friendDto := dto.FriendDto{Friends: []string{"a@gmail.com", "b@gmail.com"}}
	firstEmailId := int64(1)
	secondEmailId := int64(2)
	relationship := entity.Relationship{FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: enum.FRIEND}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", friendDto.Friends[0]).Return(firstEmailId)
	userServiceMock.On("FindUserIdByEmail", friendDto.Friends[1]).Return(secondEmailId)
	relationshipServiceMock.On("IsFriendedOrBlocked", firstEmailId, secondEmailId).Return(false)
	relationshipServiceMock.On("CreateRelationship", relationship).Return(true)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.CreateFriend(friendDto)
	assert.Equal(t, true, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}

func TestCreateSubscribe(t *testing.T) {
	requestDto := dto.RequestDto{Requestor: "a@gmail.com", Target: "b@gmail.com"}
	requestorId := int64(1)
	targetId := int64(2)
	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.SUBSCRIBE}
	subscriberEmailIds := []int64{3,4,5}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", requestDto.Requestor).Return(requestorId)
	userServiceMock.On("FindUserIdByEmail", requestDto.Target).Return(targetId)
	relationshipServiceMock.On("FindSubscribersByEmailId", targetId).Return(subscriberEmailIds)
	relationshipServiceMock.On("IsBlocked", requestorId, targetId).Return(false)
	relationshipServiceMock.On("CreateRelationship", relationship).Return(true)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.CreateSubscribe(requestDto)
	assert.Equal(t, true, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}

func TestCreateBlock(t *testing.T) {
	requestDto := dto.RequestDto{Requestor: "a@gmail.com", Target: "b@gmail.com"}
	requestorId := int64(1)
	targetId := int64(2)
	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.BLOCK}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", requestDto.Requestor).Return(requestorId)
	userServiceMock.On("FindUserIdByEmail", requestDto.Target).Return(targetId)
	relationshipServiceMock.On("IsBlocked", requestorId, targetId).Return(false)
	relationshipServiceMock.On("CreateRelationship", relationship).Return(true)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.CreateBlock(requestDto)
	assert.Equal(t, true, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}

func TestGetFriendsListByEmail(t *testing.T) {
	emailDto := dto.EmailDto{Email: "a@gmail.com"}
	emailId := int64(1)
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 1, SecondEmailId: 3, Status: enum.FRIEND}
	relationships := []entity.Relationship{relationship1, relationship2}
	emailIds := []int64{2, 3}
	emails := []string{"b@gmail.com", "c@gmail.com"}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", emailDto.Email).Return(emailId)
	relationshipServiceMock.On("FindByEmailIdAndStatus", emailId, []int64{enum.FRIEND}).Return(relationships)
	userServiceMock.On("FindByIds", emailIds).Return(emails)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.GetFriendsListByEmail(emailDto)
	assert.Equal(t, emails, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}

func TestGetCommonFriends(t *testing.T) {
	friendDto := dto.FriendDto{Friends: []string{"b@gmail.com", "c@gmail.com"}}
	firstEmailId := int64(2)
	secondEmailId := int64(3)
	commonEmailIds := []int64{1}
	commonEmails := []string{"a@gmail.com"}

	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 1, SecondEmailId: 3, Status: enum.FRIEND}
	relationships := []entity.Relationship{relationship1, relationship2}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", friendDto.Friends[0]).Return(firstEmailId)
	userServiceMock.On("FindUserIdByEmail", friendDto.Friends[1]).Return(secondEmailId)
	relationshipServiceMock.On("FindByFirstOrSecondEmailIdAndStatus", firstEmailId, secondEmailId, []int64{enum.FRIEND}).Return(relationships)
	userServiceMock.On("FindByIds", commonEmailIds).Return(commonEmails)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.GetCommonFriends(friendDto)
	assert.Equal(t, commonEmails, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}

func TestGetReceiversList(t *testing.T) {
	senderDto := dto.SenderDto{Sender: "a@gmail.com", Text: "Hi hi!! x@gmail.com"}
	senderEmailId := int64(1)

	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 1, SecondEmailId: 2, Status: enum.BLOCK}

	friendRelationships := []entity.Relationship{relationship1}
	blockedRelationships := []entity.Relationship{relationship2}

	subscriberEmailIds := []int64{3}
	receiverEmailIds := []int64{3}
	emails := []string{"c@gmail.com"}
	expectedEmails := []string{"c@gmail.com", "x@gmail.com"}

	userServiceMock := UserServiceMock{}
	relationshipServiceMock := RelationshipServiceMock{}

	userServiceMock.On("FindUserIdByEmail", senderDto.Sender).Return(senderEmailId)
	relationshipServiceMock.On("FindByEmailIdAndStatus", senderEmailId, []int64{enum.FRIEND}).Return(friendRelationships)
	relationshipServiceMock.On("FindSubscribersByEmailId", senderEmailId).Return(subscriberEmailIds)
	relationshipServiceMock.On("FindByEmailIdAndStatus", senderEmailId, []int64{enum.BLOCK}).Return(blockedRelationships)
	userServiceMock.On("FindByIds", receiverEmailIds).Return(emails)

	friendService := FriendService{relationshipServiceMock, userServiceMock}

	result, _ := friendService.GetReceiversList(senderDto)
	assert.Equal(t, expectedEmails, result)
	userServiceMock.AssertExpectations(t)
	relationshipServiceMock.AssertExpectations(t)
}