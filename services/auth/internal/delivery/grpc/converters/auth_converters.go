package converters

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
)

func GRPCRegisterRequestToPortsRegisterRequest(req *auth.RegisterRequest) ports.RegisterRequest {
	return ports.RegisterRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	}
}

func PostRegisterResponseToGRPCRegisterResponse(resp ports.RegisterResponse) *auth.RegisterResponse {
	return &auth.RegisterResponse{
		Token: resp.Token,
		User: &auth.User{
			Id:       resp.User.ID,
			Username: resp.User.Username,
			Email:    resp.User.Email,
		},
	}
}
