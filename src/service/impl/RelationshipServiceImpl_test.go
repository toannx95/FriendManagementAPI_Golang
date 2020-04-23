package impl

import (
	"friend/entity"
	"friend/enum"
	"friend/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRelationship(t *testing.T) {
	relationship := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("CreateRelationship", relationship).Return(true)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, true, relationshipService.CreateRelationship(relationship))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestFindByTwoEmailIdsAndStatus(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 2, SecondEmailId: 3, Status: enum.FRIEND}

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{enum.FRIEND, enum.SUBSCRIBE}
	expectedResult := []entity.Relationship{relationship1, relationship2}

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByTwoEmailIdsAndStatus", firstEmailId, secondEmailId, status).Return(expectedResult)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestFindByEmailIdAndStatus(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 2, SecondEmailId: 3, Status: enum.FRIEND}

	emailId := int64(2)
	status := []int64{enum.FRIEND}
	expectedResult := []entity.Relationship{relationship1, relationship2}

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByEmailIdAndStatus", emailId, status).Return(expectedResult)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.FindByEmailIdAndStatus(emailId, status))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestFindByFirstOrSecondEmailIdAndStatus(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 2, SecondEmailId: 3, Status: enum.FRIEND}
	relationship3 := entity.Relationship{Id: 3, FirstEmailId: 1, SecondEmailId: 4, Status: enum.SUBSCRIBE}

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{enum.FRIEND, enum.SUBSCRIBE}
	expectedResult := []entity.Relationship{relationship1, relationship2, relationship3}

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByFirstOrSecondEmailIdAndStatus", firstEmailId, secondEmailId, status).Return(expectedResult)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, status))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestFindSubscribersByEmailId(t *testing.T) {
	emailId := int64(2)
	expectedResult := []int64{1, 3, 4}

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindSubscribersByEmailId", emailId).Return(expectedResult)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.FindSubscribersByEmailId(emailId))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestIsFriendedOrBlocked(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.FRIEND}
	relationship2 := entity.Relationship{Id: 2, FirstEmailId: 2, SecondEmailId: 3, Status: enum.FRIEND}

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{enum.FRIEND, enum.BLOCK}
	listRelationships := []entity.Relationship{relationship1, relationship2}
	expectedResult := true

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByTwoEmailIdsAndStatus", firstEmailId, secondEmailId, status).Return(listRelationships)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.IsFriendedOrBlocked(firstEmailId, secondEmailId))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestIsSubscribedOrBlocked(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.BLOCK}

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{enum.SUBSCRIBE, enum.BLOCK}
	listRelationships := []entity.Relationship{relationship1}
	expectedResult := true

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByTwoEmailIdsAndStatus", firstEmailId, secondEmailId, status).Return(listRelationships)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.IsSubscribedOrBlocked(firstEmailId, secondEmailId))
	relationshipRepositoryMock.AssertExpectations(t)
}

func TestIsBlocked(t *testing.T) {
	relationship1 := entity.Relationship{Id: 1, FirstEmailId: 1, SecondEmailId: 2, Status: enum.BLOCK}

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{enum.BLOCK}
	listRelationships := []entity.Relationship{relationship1}
	expectedResult := true

	relationshipRepositoryMock := repository.RelationshipRepositoryMock{}
	relationshipRepositoryMock.On("FindByTwoEmailIdsAndStatus", firstEmailId, secondEmailId, status).Return(listRelationships)
	relationshipService := RelationshipServiceImpl{relationshipRepositoryMock}

	assert.Equal(t, expectedResult, relationshipService.IsBlocked(firstEmailId, secondEmailId))
	relationshipRepositoryMock.AssertExpectations(t)
}