package resolver

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/search/internal/graph/generated"
	"github.com/popeskul/awesome-messanger/services/search/internal/graph/model"
	"github.com/popeskul/awesome-messanger/services/search/internal/service"
)

type IQueryResolver interface {
	Search(ctx context.Context, nickname string) ([]*model.User, error)
}

type Resolver struct {
	Service service.IService
}

func NewResolver(service service.IService) generated.ResolverRoot {
	return &Resolver{
		Service: service,
	}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return nil
}
