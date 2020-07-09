package authorization

import (
	"go-rest-skeleton/infrastructure/exception"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// Token is a struct.
type Token struct{}

// NewToken is a constructor will initialize token.
func NewToken() *Token {
	return &Token{}
}

// TokenInterface is an interface.
type TokenInterface interface {
	CreateToken(UUID string) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

// Token implements the TokenInterface.
var _ TokenInterface = &Token{}

// CreateToken is a function uses to create a new token.
func (t *Token) CreateToken(UUID string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.TokenUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = td.TokenUUID + "++" + UUID

	var err error

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUUID
	atClaims["uuid"] = UUID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("APP_PRIVATE_KEY")))
	if err != nil {
		return nil, err
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["uuid"] = UUID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("APP_PRIVATE_KEY")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// TokenValid is a function uses to validate token.
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken is a function to verify token.
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
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
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractTokenMetadata is a function to extract meta data from the token.
func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		UUID, _ := claims["uuid"].(string)
		return &AccessDetails{
			TokenUUID: accessUUID,
			UUID:      UUID,
		}, nil
	}
	return nil, err
}
