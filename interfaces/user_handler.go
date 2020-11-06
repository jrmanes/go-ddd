package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jrmanes/go-ddd/application"
	"github.com/jrmanes/go-ddd/domain/entity/users"
)

// Users struct defines the dependencies that we will use
type Users struct {
	user application.UserAppInterface
}

// Users constructor
func NewUsers(user application.UserAppInterface) *Users {
	return &Users{user: user}
}

func (s *Users) SaveUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}
	newUser, err := s.user.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser.PublicUser())
}

func (s *Users) GetUsers(c *gin.Context) {
	users := users.Users{} //customize user
	var err error
	//us, err = application.UserApp.GetUsers()
	users, err = s.user.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users.PublicUsers())
}

func (s *Users) GetUser(c *gin.Context) {
	fmt.Println("id_user", c.Param("id_user"))
	userId, err := strconv.ParseUint(c.Param("id_user"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := s.user.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user.PublicUser())
}

// standar library
//func DecodeUserRequest(w http.ResponseWriter, r *http.Request) users.User {
//	var user users.User
//	errs := json.NewDecoder(r.Body).Decode(&user)
//	if errs != nil {
//		http.Error(w, errs.Error(), http.StatusBadRequest)
//	}
//	return user
//}
//
//func (u *Users) SaveUser(w http.ResponseWriter, r *http.Request) {
//	var user = DecodeUserRequest(w, r)
//	_, err := u.user.SaveUser(&user)
//	if err != nil {
//		fmt.Println("creation error:", http.StatusInternalServerError)
//		//return http.StatusInternalServerError
//	} else {
//		fmt.Println("creation Ok:", http.StatusCreated)
//		//return http.StatusCreated
//	}
//}
