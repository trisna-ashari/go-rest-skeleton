package services

import (
	"context"
	"fmt"
	"go-rest-skeleton/grpc/proto/v1/user"
	"go-rest-skeleton/infrastructure/persistence"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type User struct {
	DBServices *persistence.Repositories
}

// NewUser is a constructor.
func NewUser(repo *persistence.Repositories) *User {
	return &User{
		DBServices: repo,
	}
}

// Run starts the server
func (u *User) Run(port int) error {
	srv := grpc.NewServer()
	user.RegisterUserServiceServer(srv, u)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4040))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return err
}

func (u *User) GetUser(ctx context.Context, requestUser *user.RequestUser) (*user.User, error) {
	userData, err := u.DBServices.User.GetUser(requestUser.Uuid)
	if err != nil {
		return nil, err
	}

	userResult := &user.User{
		Uuid:  userData.UUID,
		Name:  userData.Name,
		Email: userData.Email,
		Phone: userData.Phone,
	}

	return userResult, nil
}

func (u *User) GetUsers(ctx context.Context, users *user.RequestUsers) (*user.Users, error) {
	panic("implement me")
}
