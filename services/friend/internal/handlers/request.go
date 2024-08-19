package handlers

import "github.com/popeskul/awesome-messanger/services/friend/internal/models"

type PostRespondFriendRequestJSONRequestBody struct {
	FriendId string `json:"friend_id"`
	Response string `json:"response"`
}

func (r *PostRespondFriendRequestJSONRequestBody) ConvertToModel() *models.Friend {
	return &models.Friend{
		FriendId: r.FriendId,
	}
}

type FriendRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	FriendId string `json:"friend_id" validate:"required"`
}

func (r *FriendRequest) ConvertToModel() *models.Friend {
	return &models.Friend{
		UserId:   r.UserId,
		FriendId: r.FriendId,
	}
}
