package repository

import "main/entity"

type RelationshipRepository interface {
	CreateRelationship(relationship entity.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship
	FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindSubscribersByEmailId(emailId int64) []int64
}