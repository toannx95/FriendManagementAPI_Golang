package service

import "main/entity"

type RelationshipService interface {
	CreateRelationship(relationship entity.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship
	FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindSubscribersByEmailId(emailId int64) []int64
	IsFriendedOrBlocked(firstEmailId int64, secondEmailId int64) bool
	IsSubscribedOrBlocked(firstEmailId int64, secondEmailId int64) bool
	IsBlocked(firstEmailId int64, secondEmailId int64) bool
}
