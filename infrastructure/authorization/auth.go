package authorization

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (uint64, error)
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

//Save token metadata to Redis
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

//Check the metadata saved
func (tk *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := tk.client.Get(tk.client.Context(), tokenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

//Once a user row in the token table
func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UUID)
	//delete access token
	deletedAt, err := tk.client.Del(tk.client.Context(), authD.TokenUuid).Result()
	//fmt.Println(err)
	if err != nil {
		return err
	}
	//delete refresh token
	fmt.Println(refreshUuid)
	deletedRt, err := tk.client.Del(tk.client.Context(), refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	deleted, err := tk.client.Del(tk.client.Context(), refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}