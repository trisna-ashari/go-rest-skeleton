package authorization

import (
	"fmt"
	"go-rest-skeleton/infrastructure/exception"
	"time"

	"github.com/go-redis/redis"
)

// AuthInterface is an interface.
type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

// ClientData is a struct to store redis client.
type ClientData struct {
	client *redis.Client
}

var _ AuthInterface = &ClientData{}

// NewAuth will initialize redis client.
func NewAuth(client *redis.Client) *ClientData {
	return &ClientData{client: client}
}

// AccessDetails is struct to store access detail.
type AccessDetails struct {
	TokenUUID string
	UUID      string
}

// TokenDetails is struct to store token detail.
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUUID    string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// CreateAuth will save token metadata to Redis.
func (tk *ClientData) CreateAuth(UUID string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(tk.client.Context(), td.TokenUUID, UUID, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(tk.client.Context(), td.RefreshUUID, UUID, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return exception.ErrorTextNoRecordInsertedToRedis
	}
	return nil
}

// FetchAuth will check the metadata saved.
func (tk *ClientData) FetchAuth(tokenUUID string) (string, error) {
	UUID, err := tk.client.Get(tk.client.Context(), tokenUUID).Result()
	if err != nil {
		return "", err
	}
	return UUID, nil
}

// DeleteTokens will delete token once a user row in the token table.
func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	// Get the refresh uuid
	refreshUUID := fmt.Sprintf("%s++%s", authD.TokenUUID, authD.UUID)
	// Delete access token
	deletedAt, err := tk.client.Del(tk.client.Context(), authD.TokenUUID).Result()
	if err != nil {
		return err
	}
	// Delete refresh token
	deletedRt, err := tk.client.Del(tk.client.Context(), refreshUUID).Result()
	if err != nil {
		return err
	}
	// When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return exception.ErrorTextUnauthorized
	}
	return nil
}

// DeleteRefresh will delete refresh token from redis.
func (tk *ClientData) DeleteRefresh(refreshUUID string) error {
	// Delete refresh token
	deleted, err := tk.client.Del(tk.client.Context(), refreshUUID).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
