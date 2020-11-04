package routers

import (
	"context"
	"fmt"
	"go-rest-skeleton/graph"
	"go-rest-skeleton/graph/resolver"
	"go-rest-skeleton/interfaces/middleware"
	"go-rest-skeleton/pkg/response"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func graphqlRoute(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	guard := middleware.Guard(rg.authGateway)

	e.Use(ginContextToContextMiddleware())
	e.POST("/query", guard.Authenticate(), graphqlHandler(r))
}

func graphqlHandler(r *Router) gin.HandlerFunc {

	h := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{
		Resolvers: &graph.Resolver{DBServices: r.dbService},
	}))

	h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		c, err := ginContextFromContext(ctx)
		if err != nil {
			log.Fatal(err)
		}

		errPresenter := graphql.DefaultErrorPresenter(ctx, e)
		errPresenter.Message = response.NewError(c, e.Error()).Message

		return errPresenter
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

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
