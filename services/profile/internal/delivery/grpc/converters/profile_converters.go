package converters

import (
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

func ProtoCreateProfileRequestToPortsRegisterRequest(req *profile.CreateProfileRequest) *ports.CreateProfileRequest {
	return &ports.CreateProfileRequest{
		UserId:    req.GetUserId(),
		Nickname:  req.GetNickname(),
		Bio:       req.GetBio(),
		AvatarUrl: req.GetAvatarUrl(),
	}
}

func ProtoUpdateProfileRequestToPortsUpdateRequest(req *profile.UpdateProfileRequest) *ports.UpdateProfileRequest {
	return &ports.UpdateProfileRequest{
		UserId:    req.GetUserId(),
		Nickname:  req.GetNickname(),
		Bio:       req.GetBio(),
		AvatarUrl: req.GetAvatarUrl(),
	}
}

func ProtoGetProfileRequestToPortsGetRequest(req *profile.GetProfileRequest) *ports.GetProfileRequest {
	return &ports.GetProfileRequest{
		UserId: req.GetUserId(),
	}
}

func DomainGetProfileResponseToProtoGetProfileResponse(resp *domain.Profile) *profile.GetProfileResponse {
	return &profile.GetProfileResponse{
		Nickname:  resp.Nickname,
		Bio:       resp.Bio,
		AvatarUrl: resp.AvatarUrl,
	}
}
