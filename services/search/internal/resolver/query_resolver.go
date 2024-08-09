package resolver

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/search/internal/graph/model"
)

type QueryResolver struct {
	*Resolver
}

func (r *QueryResolver) Search(ctx context.Context, nickname string) ([]*model.User, error) {
	user, err := r.Service.UserService().FindUserByNickname(ctx, nickname)
	if err != nil {
		return nil, err
	}
	return []*model.User{user}, nil
}
