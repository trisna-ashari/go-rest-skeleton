package user_v2_0_0

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/interfaces/middleware"
	"golang.org/x/exp/errors"
	"net/http"
)

//Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
}

//Users constructor
func NewUsers(us application.UserAppInterface, rd authorization.AuthInterface, tk authorization.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

func (s *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New("Unprocessable Entity"))
		return
	}
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New("Unprocessable Entity"))
		return
	}
	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Unable to create user"))
		return
	}
	c.Status(http.StatusCreated)
	middleware.Formatter(c, newUser.DetailUser(), "Successfully create a new user")
}

func (s *Users) GetUsers(c *gin.Context) {
	users := entity.Users{}
	var err error
	users, meta, err := s.us.GetUsers(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	middleware.FormatterWithMeta(c, users.DetailUsers(), "Successfully get user list", meta)
}

func (s *Users) GetUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindUri(&userEntity.UUID); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad Request"))
		return
	}
	UUID := c.Param("uuid")
	user, err := s.us.GetUser(UUID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithError(http.StatusNotFound, errors.New("Record not found"))
			return
		} else {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}
	}
	middleware.Formatter(c, user.DetailUser(), "Successfully get user detail")
}