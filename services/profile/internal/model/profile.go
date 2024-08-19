package model

type Profile struct {
	UserId    string `json:"user_id,omitempty" validate:"required"`
	Nickname  string `json:"nickname,omitempty" validate:"required"`
	Bio       string `json:"bio,omitempty" validate:"required"`
	AvatarUrl string `json:"avatar_url,omitempty" validate:"required"`
}

type UpdateProfileRequest struct {
	UserId    string `json:"user_id,omitempty" validate:"required"`
	Nickname  string `json:"nickname,omitempty" validate:"required"`
	Bio       string `json:"bio,omitempty" validate:"required"`
	AvatarUrl string `json:"avatar_url,omitempty" validate:"required"`
}

type GetProfileRequest struct {
	UserId string `json:"user_id,omitempty" validate:"required"`
}
