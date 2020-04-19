package impl

import (
	"entity"
	"enum"
	"repository"
)

type RelationshipServiceImpl struct {
	RelationshipRepository repository.RelationshipRepository
}

func (r RelationshipServiceImpl) CreateRelationship(relationship entity.Relationship) bool {
	return r.RelationshipRepository.CreateRelationship(relationship)
}

func (r RelationshipServiceImpl) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	return r.RelationshipRepository.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status)
}

func (r RelationshipServiceImpl) FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship {
	return r.RelationshipRepository.FindByEmailIdAndStatus(emailId, status)
}

func (r RelationshipServiceImpl) FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	return  r.RelationshipRepository.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, status)
}

func (r RelationshipServiceImpl) FindSubscribersByEmailId(emailId int64) []int64 {
	return r.RelationshipRepository.FindSubscribersByEmailId(emailId)
}

func (r RelationshipServiceImpl) IsFriendedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.FRIEND, enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}

func (r RelationshipServiceImpl) IsSubscribedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.SUBSCRIBE, enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}

func (r RelationshipServiceImpl) IsBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}


