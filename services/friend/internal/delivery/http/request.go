package http

import "github.com/popeskul/awesome-messanger/services/friend/internal/core/models"

type FriendRequest struct {
	UserId   string `json:"userId" validate:"required"`
	FriendId string `json:"friendId" validate:"required"`
}

type Friend struct {
	UserId   string `json:"userId"`
	FriendId string `json:"friendId"`
}

type PostRespondFriendRequestJSONRequestBody struct {
	UserId   string `json:"userId" validate:"required"`
	FriendId string `json:"friendId" validate:"required"`
	Accept   bool   `json:"accept" validate:"required"`
}

func (f FriendRequest) ConvertToModel() *models.Friend {
	return &models.Friend{
		UserId:   f.UserId,
		FriendId: f.FriendId,
	}
}

func (f PostRespondFriendRequestJSONRequestBody) ConvertToModel() *models.Friend {
	return &models.Friend{
		UserId:   f.UserId,
		FriendId: f.FriendId,
	}
}
