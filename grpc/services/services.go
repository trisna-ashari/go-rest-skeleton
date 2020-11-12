package services

import (
	"encoding/json"
	"fmt"
	"go-rest-skeleton/grpc/client"
	"go-rest-skeleton/grpc/proto/v1/user"
	"net/http"

	"google.golang.org/grpc"
)

type Services struct {
	userClient user.UserServiceClient
}

func (s *Services) Run(port int) error {
	mux := client.NewServeMux()
	mux.Handle("/user", http.HandlerFunc(s.GetUserHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func NewServices(userClient *grpc.ClientConn) *Services {
	return &Services{
		userClient: user.NewUserServiceClient(userClient),
	}
}

func (s *Services) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := r.Context()

	uuid := r.URL.Query().Get("uuid")
	if uuid == "" {
		http.Error(w, "Please specify uuid", http.StatusBadRequest)
		return
	}

	userResult, errGet := s.userClient.GetUser(ctx, &user.RequestUser{Uuid: uuid})
	if errGet != nil {
		http.Error(w, errGet.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(userResult)
}
