package repository

import (
	"friend/entity"
	"github.com/stretchr/testify/mock"
)

type RelationshipRepositoryMock struct {
	mock.Mock
}

func (r RelationshipRepositoryMock) CreateRelationship(relationship entity.Relationship) bool {
	returnVals := r.Called(relationship)
	return returnVals.Get(0).(bool)
}

func (r RelationshipRepositoryMock) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(firstEmailId, secondEmailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipRepositoryMock) FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(emailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipRepositoryMock) FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(firstEmailId, secondEmailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipRepositoryMock) FindSubscribersByEmailId(emailId int64) []int64 {
	returnVals := r.Called(emailId)
	return returnVals.Get(0).([]int64)
}
