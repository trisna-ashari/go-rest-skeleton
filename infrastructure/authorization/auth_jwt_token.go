package authorization

import (
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/infrastructure/message/exception"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v8"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

const expiredTime = 60
const totalHoursInADay = 24
const totalDaysInAWeek = 7
const authorizationLen = 2

// Token is a struct.
type Token struct {
	kc config.KeyConfig
	rd *redis.Client
}

// NewToken is a constructor will initialize token.
func NewToken(config config.KeyConfig, rd *redis.Client) *Token {
	return &Token{
		kc: config,
		rd: rd,
	}
}

// TokenInterface is an interface.
type TokenInterface interface {
	CreateToken(UUID string, AUD string) (*TokenDetails, error)
	ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error)
}

// Token implements the TokenInterface.
var _ TokenInterface = &Token{}

// CreateToken is a function uses to create a new token.
func (t *Token) CreateToken(UUID string, AUD string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * expiredTime).Unix()
	td.TokenUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * totalHoursInADay * totalDaysInAWeek).Unix()
	td.RefreshUUID = td.TokenUUID + "++" + UUID

	var err error

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUUID
	atClaims["sub"] = UUID
	atClaims["exp"] = td.AtExpires
	if AUD != "" {
		atClaims["aud"] = AUD
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(t.kc.AppPrivateKey))
	if err != nil {
		return nil, err
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["sub"] = UUID
	rtClaims["exp"] = td.RtExpires
	if AUD != "" {
		rtClaims["aud"] = AUD
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(t.kc.AppPrivateKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// TokenValid is a function uses to validate token.
func TokenValid(c *gin.Context, t *Token) (*AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	metadata, errExtract := t.ExtractTokenMetadata(c)
	if errExtract != nil {
		return nil, errExtract
	}

	_, err = t.rd.Get(t.rd.Context(), metadata.TokenUUID).Result()
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	return metadata, nil
}

// VerifyToken is a function to verify token.
func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, exception.ErrorTextAnErrorOccurred
		}
		return []byte(os.Getenv("APP_PRIVATE_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken is a function uses to get the token from the request body.
func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == authorizationLen {
		return strArr[1]
	}

	accessToken := c.DefaultQuery("access_token", "")
	if accessToken != "" {
		return accessToken
	}

	return ""
}

// ExtractTokenMetadata is a function to extract meta data from the token.
func (t *Token) ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		UUID, _ := claims["sub"].(string)
		return &AccessDetails{
			TokenUUID: accessUUID,
			UUID:      UUID,
		}, nil
	}
	return nil, err
}
