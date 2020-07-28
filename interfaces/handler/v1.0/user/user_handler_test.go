package userv1point00

import (
	"encoding/json"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/util"
	"go-rest-skeleton/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// GetUser Test.
func TestGetUser_Success(t *testing.T) {
	var user entity.User
	var userApp mock.UserAppInterface
	userHandler := NewUsers(&userApp)
	UUID := uuid.New().String()
	userApp.GetUserFn = func(string) (*entity.User, error) {
		return &entity.User{
			UUID:      UUID,
			FirstName: "Example",
			LastName:  "User",
			Email:     "example@test.com",
		}, nil
	}
	r := gin.Default()
	v1 := r.Group("/api/v1/external/")
	v1.GET("/users/:uuid", userHandler.GetUser)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/external/users/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response := util.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &user)

	assert.Equal(t, w.Code, 200)
	assert.EqualValues(t, user.UUID, UUID)
	assert.EqualValues(t, user.FirstName, "Example")
	assert.EqualValues(t, user.LastName, "User")
	assert.EqualValues(t, user.Email, "example@test.com")
}
