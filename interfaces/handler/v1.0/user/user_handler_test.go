package userv1point00

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/pkg/encoder"
	"go-rest-skeleton/pkg/util"
	"go-rest-skeleton/tests/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSaveUser_Success Test.
func TestSaveUser_Success(t *testing.T) {
	var userData entity.User
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface
	userHandler := NewUsers(&userApp, &storageApp)
	userJSON := `{
		"name": "Example",
		"email": "example@test.com",
		"phone": "+6285725833220",
		"password": "password"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/users", userHandler.SaveUser)

	userApp.SaveUserFn = func(user *entity.User) (*entity.User, map[string]string, error) {
		return &entity.User{
			UUID:  UUID,
			Name:  "Example",
			Email: "example@test.com",
			Phone: "+6285725833220",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/users", bytes.NewBufferString(userJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &userData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, userData.UUID, UUID)
	assert.EqualValues(t, userData.Name, "Example")
	assert.EqualValues(t, userData.Email, "example@test.com")
	assert.EqualValues(t, userData.Phone, "+6285725833220")
}

func TestSaveUser_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"name": "", "email": "example@test.com","password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": "victor", "email": "example@test.com","password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": "victor", "email": "","password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": "victor", "email": "example@test.com","password": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"email": "example@test","password": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": 1234, "email": "example@test.com","password": "password"}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var userApp mock.UserAppInterface
		var storageApp mock.StorageAppInterface
		userHandler := NewUsers(&userApp, &storageApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/users", userHandler.SaveUser)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		r.ServeHTTP(w, c.Request)

		validationErr := make(map[string]string)
		response := encoder.ResponseDecoder(w.Body)
		data, _ := json.Marshal(response["data"])

		err = json.Unmarshal(data, &validationErr)
		if err != nil {
			t.Errorf("error unmarshalling error %s\n", err)
		}
		assert.Equal(t, w.Code, v.statusCode)
	}
}

// TestUpdateUser_Success Test.
func TestUpdateUser_Success(t *testing.T) {
	var userData entity.User
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface
	userHandler := NewUsers(&userApp, &storageApp)
	userJSON := `{
		"name": "Example",
		"email": "example@test.com",
		"phone": "+6285725833220",
		"password": "password"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/users/:uuid", userHandler.UpdateUser)

	userApp.UpdateUserFn = func(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
		return &entity.User{
			UUID:  UUID,
			Name:  "Example",
			Email: "example@test.com",
			Phone: "+6285725833220",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPut, "/api/v1/external/users/"+UUID, bytes.NewBufferString(userJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &userData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, userData.UUID, UUID)
	assert.EqualValues(t, userData.Name, "Example")
	assert.EqualValues(t, userData.Email, "example@test.com")
	assert.EqualValues(t, userData.Phone, "+6285725833220")
}

// TestGetUser_Success Test.
func TestGetUser_Success(t *testing.T) {
	var userData entity.User
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	userHandler := NewUsers(&userApp, &storageApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/users/:uuid", userHandler.GetUser)

	userApp.GetUserFn = func(string) (*entity.User, error) {
		return &entity.User{
			UUID:       UUID,
			Name:       "Example",
			Email:      "example@test.com",
			AvatarUUID: UUID,
		}, nil
	}

	storageApp.GetFileFn = func(string) (interface{}, error) {
		return UUID, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/users/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &userData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, userData.UUID, UUID)
	assert.EqualValues(t, userData.Name, "Example")
	assert.EqualValues(t, userData.Email, "example@test.com")
}

// TestGetUsers_Success Test.
func TestGetUsers_Success(t *testing.T) {
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface
	var usersData []entity.User
	var metaData repository.Meta
	userHandler := NewUsers(&userApp, &storageApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/users", userHandler.GetUsers)
	userApp.GetUsersFn = func(params *repository.Parameters) ([]entity.User, interface{}, error) {
		users := []entity.User{
			{
				UUID:  UUID,
				Name:  "Example 1",
				Email: "example1@test.com",
			},
			{
				UUID:  UUID,
				Name:  "Example 2",
				Email: "example2@test.com",
			},
		}
		meta := repository.NewMeta(params, int64(len(users)))
		return users, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &usersData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(usersData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteUser_Success Test.
func TestDeleteUser_Success(t *testing.T) {
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface
	userHandler := NewUsers(&userApp, &storageApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/users/:uuid", userHandler.DeleteUser)

	userApp.DeleteUserFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/users/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteUser_Failed_UserNotFound Test.
func TestDeleteUser_Failed_UserNotFound(t *testing.T) {
	var userApp mock.UserAppInterface
	var storageApp mock.StorageAppInterface
	userHandler := NewUsers(&userApp, &storageApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/users/:uuid", userHandler.DeleteUser)

	userApp.DeleteUserFn = func(UUID string) error {
		return exception.ErrorTextUserNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/users/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
