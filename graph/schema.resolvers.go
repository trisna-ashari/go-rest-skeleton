package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/graph/model"
	"go-rest-skeleton/graph/resolver"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*entity.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*entity.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	c, err := ginContextFromContext(ctx)
	if err != nil {
		return false, newError(c, err)
	}

	err = r.DBServices.User.DeleteUser(uuid)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserNotFound) {
			return false, newError(c, err)
		}
		return false, newError(c, err)
	}

	return true, nil
}

func (r *queryResolver) Users(
	ctx context.Context,
	search *model.SearchUserInput,
	orderBy model.UserOrderFields,
	pagination model.PaginationInput) (*model.UserConnection, error) {
	c, err := ginContextFromContext(ctx)
	if err != nil {
		return nil, newError(c, err)
	}

	parameters := repository.NewGqlParameters(c, newGqlParameters(search, orderBy, pagination))
	users, meta, err := r.DBServices.User.GetUsers(parameters)
	if err != nil {
		return nil, newError(c, err)
	}

	userList := model.UserConnection{
		Pagination: newPageInfo(meta),
		List:       users,
	}

	return &userList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() resolver.MutationResolver { return &mutationResolver{r} }

// Query returns resolver.QueryResolver implementation.
func (r *Resolver) Query() resolver.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func newKeywords(search *model.SearchUserInput) map[string]interface{} {
	var keywords = make(map[string]interface{})
	searchRec, _ := json.Marshal(search.Keywords)
	err := json.Unmarshal(searchRec, &keywords)
	if err != nil {
		return keywords
	}

	return keywords
}

func newGqlParameters(
	search *model.SearchUserInput,
	orderBy model.UserOrderFields,
	pagination model.PaginationInput) *repository.GqlParameters {
	return &repository.GqlParameters{
		SearchMethod:    search.Rules.Method.String(),
		SearchCondition: search.Rules.Condition.String(),
		SearchKeywords:  newKeywords(search),
		Order:           orderBy.String(),
		Page:            pagination.Page,
		PerPage:         pagination.PerPage,
	}
}

func newPageInfo(meta *repository.Meta) *model.Pagination {
	return &model.Pagination{
		Page:    meta.Page,
		PerPage: meta.PerPage,
		Total:   int(meta.Total),
	}
}

func newError(c *gin.Context, err error) error {
	message := response.NewError(c, err.Error()).Message

	return errors.New(message)
}
