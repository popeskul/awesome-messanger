package ports

type CreateProfileRequest struct {
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type CreateProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateProfileRequest struct {
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type GetProfileRequest struct {
	UserId string `json:"user_id"`
}

type GetProfileResponse struct {
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}
