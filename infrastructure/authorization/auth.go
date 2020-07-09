package authorization

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type ClientData struct {
	client *redis.Client
}

var _ AuthInterface = &ClientData{}

func NewAuth(client *redis.Client) *ClientData {
	return &ClientData{client: client}
}

type AccessDetails struct {
	TokenUuid string
	UUID    string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// Save token metadata to Redis
func (tk *ClientData) CreateAuth(UUID string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(tk.client.Context(), td.TokenUuid, UUID, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(tk.client.Context(), td.RefreshUuid, UUID, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

// Check the metadata saved
func (tk *ClientData) FetchAuth(tokenUuid string) (string, error) {
	UUID, err := tk.client.Get(tk.client.Context(), tokenUuid).Result()
	if err != nil {
		return "", err
	}
	return UUID, nil
}

// Once a user row in the token table
func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	// Get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UUID)
	// Delete access token
	deletedAt, err := tk.client.Del(tk.client.Context(), authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	// Delete refresh token
	deletedRt, err := tk.client.Del(tk.client.Context(), refreshUuid).Result()
	if err != nil {
		return err
	}
	// When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("api.msg.error.an_error_occurred")
	}
	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	// Delete refresh token
	deleted, err := tk.client.Del(tk.client.Context(), refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}