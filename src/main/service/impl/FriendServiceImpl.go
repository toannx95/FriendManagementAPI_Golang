package impl

import (
	"encoding/json"
	"main/dto"
	"main/entity"
	"main/enum"
	ex "main/exception"
	"main/service"
	"main/utils"
	"net/http"
)

type FriendServiceImpl struct {
	RelationshipService service.RelationshipService
	UserService service.UserService
}

func (f FriendServiceImpl) CreateFriend(r *http.Request) (bool, *ex.Exception) {
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]
	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	firstEmailId, err1 := f.UserService.FindUserIdByEmail(firstEmail)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	secondEmailId, err2 := f.UserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// check friend and blocked
	if f.RelationshipService.IsFriendedOrBlocked(firstEmailId, secondEmailId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not make friend!"}
	}

	relationship := entity.Relationship{FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: enum.FRIEND}
	result := f.RelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when create friend!"}
	}
	return true, nil
}

func (f FriendServiceImpl) CreateSubscribe(r *http.Request) (bool, *ex.Exception) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	requestor := requestDto.Requestor
	target := requestDto.Target
	if !utils.IsFormatEmail(requestor) || !utils.IsFormatEmail(target) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	requestorId, err1 := f.UserService.FindUserIdByEmail(requestor)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	targetId, err2 := f.UserService.FindUserIdByEmail(target)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// check requestor has not subscribed target yet
	subscriberEmailIds := f.RelationshipService.FindSubscribersByEmailId(targetId)
	subscriberEmailIds = utils.RemoveItemFromList(subscriberEmailIds, targetId)

	if subscriberEmailIds != nil && len(subscriberEmailIds) > 0 && utils.Contains(subscriberEmailIds, requestorId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	// check both emails have not blocked each other
	if f.RelationshipService.IsBlocked(requestorId, targetId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.SUBSCRIBE}
	result := f.RelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when subscribe!"}
	}
	return true, nil
}

func (f FriendServiceImpl) CreateBlock(r *http.Request) (bool, *ex.Exception) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	requestor := requestDto.Requestor
	target := requestDto.Target
	if !utils.IsFormatEmail(requestor) || !utils.IsFormatEmail(target) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	requestorId, err1 := f.UserService.FindUserIdByEmail(requestor)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	targetId, err2 := f.UserService.FindUserIdByEmail(target)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	/// check blocked
	if f.RelationshipService.IsBlocked(requestorId, targetId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.BLOCK}
	result := f.RelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when subscribe!"}
	}
	return true, nil
}

func (f FriendServiceImpl) GetFriendsListByEmail(r *http.Request) ([]string, *ex.Exception) {
	var emailDto dto.EmailDto
	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	emails := []string{}
	if !utils.IsFormatEmail(emailDto.Email) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	emailId, err1 := f.UserService.FindUserIdByEmail(emailDto.Email)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	// find list relationship by an email and friend status
	relationships := f.RelationshipService.FindByEmailIdAndStatus(emailId, []int64{enum.FRIEND})
	if relationships != nil {
		// get list email_ids
		emailIds := getEmailIdsFromListRelationships(relationships)
		emailIds = utils.RemoveItemFromList(emailIds, emailId)

		// get list emails by list email_ids
		if emailIds != nil && len(emailIds) > 0 {
			emails = f.UserService.FindByIds(emailIds)
		}
	}
	return emails, nil
}

func (f FriendServiceImpl) GetCommonFriends(r *http.Request) ([]string, *ex.Exception) {
	commonEmails := []string{}
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]
	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	firstEmailId, err1 := f.UserService.FindUserIdByEmail(firstEmail)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	secondEmailId, err2 := f.UserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// find list relationship between two email and friend status
	relationships := f.RelationshipService.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, []int64{enum.FRIEND})
	if relationships != nil && len(relationships) > 0 {
		mapFirst := make(map[int64]bool)
		friendsOfFirstEmail := []int64{}

		mapSecond := make(map[int64]bool)
		friendsOfSecondEmail := []int64{}

		for _, rela := range relationships {
			// get friends of firstEmailId
			if firstEmailId == rela.FirstEmailId || firstEmailId == rela.SecondEmailId {
				if _, ok := mapFirst[rela.FirstEmailId]; !ok {
					mapFirst[rela.FirstEmailId] = true
					friendsOfFirstEmail = append(friendsOfFirstEmail, rela.FirstEmailId)
				}

				if _, ok := mapFirst[rela.SecondEmailId]; !ok {
					mapFirst[rela.SecondEmailId] = true
					friendsOfFirstEmail = append(friendsOfFirstEmail, rela.SecondEmailId)
				}
			}

			// get friends of secondEmailId
			if secondEmailId == rela.FirstEmailId || secondEmailId == rela.SecondEmailId {
				if _, ok := mapSecond[rela.FirstEmailId]; !ok {
					mapSecond[rela.FirstEmailId] = true
					friendsOfSecondEmail = append(friendsOfSecondEmail, rela.FirstEmailId)
				}

				if _, ok := mapSecond[rela.SecondEmailId]; !ok {
					mapSecond[rela.SecondEmailId] = true
					friendsOfSecondEmail = append(friendsOfSecondEmail, rela.SecondEmailId)
				}
			}
		}
		// remove emails in request
		friendsOfFirstEmail = utils.RemoveItemFromList(friendsOfFirstEmail, firstEmailId)
		friendsOfSecondEmail = utils.RemoveItemFromList(friendsOfSecondEmail, secondEmailId)

		// get common emailIds
		commonEmailIds := []int64{}
		for _, mailId := range friendsOfFirstEmail {
			if utils.Contains(friendsOfSecondEmail, mailId) {
				commonEmailIds = append(commonEmailIds, mailId)
			}
		}

		// get string emails by emailId
		if commonEmailIds != nil && len(commonEmailIds) > 0 {
			commonEmails = f.UserService.FindByIds(commonEmailIds)
		}
	}
	return commonEmails, nil
}

func (f FriendServiceImpl) GetReceiversList(r *http.Request) ([]string, *ex.Exception) {
	receiverEmails := []string{}
	var senderDto dto.SenderDto
	if err := json.NewDecoder(r.Body).Decode(&senderDto); err != nil {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Invalid request body!"}
	}

	senderEmail := senderDto.Sender
	if !utils.IsFormatEmail(senderEmail) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	senderEmailId, err1 := f.UserService.FindUserIdByEmail(senderEmail)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	// get list friend_ids by emailId
	friendRelationships := f.RelationshipService.FindByEmailIdAndStatus(senderEmailId, []int64{enum.FRIEND})
	friendEmailIds := getEmailIdsFromListRelationships(friendRelationships)
	friendEmailIds = utils.RemoveItemFromList(friendEmailIds, senderEmailId)

	// get subscribers by emailId
	subscriberEmailIds := f.RelationshipService.FindSubscribersByEmailId(senderEmailId)
	subscriberEmailIds = utils.RemoveItemFromList(subscriberEmailIds, senderEmailId)

	// list receiver emailIds
	receiverEmailIds := friendEmailIds
	for _, mailId := range subscriberEmailIds {
		if !utils.Contains(receiverEmailIds, mailId) {
			receiverEmailIds = append(receiverEmailIds, mailId)
		}
	}

	// get list blocked_ids
	blockedRelationships := f.RelationshipService.FindByEmailIdAndStatus(senderEmailId, []int64{enum.BLOCK})
	blockedMailIds := getEmailIdsFromListRelationships(blockedRelationships)
	blockedMailIds = utils.RemoveItemFromList(blockedMailIds, senderEmailId)

	// remove blocked in list receiverEmailIds
	for _, mailId := range blockedMailIds {
		receiverEmailIds = utils.RemoveItemFromList(receiverEmailIds, mailId)
	}

	// find email by list receiverEmailIds
	receiverEmails = f.UserService.FindByIds(receiverEmailIds)

	// get mentionedEmail from text of sender
	if &senderDto.Text != nil {
		mentionEmails := utils.GetEmailFromText(senderDto.Text)
		if &mentionEmails != nil && len(mentionEmails) > 0 {
			for _, email := range mentionEmails {
				receiverEmails = append(receiverEmails, email)
			}
		}
	}
	return receiverEmails, nil
}

func getEmailIdsFromListRelationships(relationships []entity.Relationship) []int64 {
	keys := make(map[int64]bool)
	set := []int64{}
	if relationships != nil && len(relationships) > 0 {
		for _, rela := range relationships {
			if _, ok := keys[rela.FirstEmailId]; !ok {
				keys[rela.FirstEmailId] = true
				set = append(set, rela.FirstEmailId)
			}

			if _, ok := keys[rela.SecondEmailId]; !ok {
				keys[rela.SecondEmailId] = true
				set = append(set, rela.SecondEmailId)
			}
		}
	}
	return set
}