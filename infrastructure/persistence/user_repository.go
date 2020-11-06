package persistence

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/jrmanes/go-ddd/domain/entity/users"
	"github.com/jrmanes/go-ddd/domain/repository"
	"github.com/jrmanes/go-ddd/infrastructure/security"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

// Implement the methods that we have defined in our repository/user_repository.go
// This was made possible using the UserRepo struct which implements the UserRepository interface, as seen in this line:

func (r *UserRepo) SaveUser(user *users.User) (*users.User, map[string]string) {
	log.Println("USER IS:", user)
	//log.Println("USER IS: name=%s,  username=%s, email=%s, pass=%s", user.Name, user.UserName, user.Email, user.Password)
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

func (r *UserRepo) GetUser(idUser uint64) (*users.User, error) {
	var user users.User
	err := r.db.Debug().Where("id_user = ?", idUser).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() ([]users.User, error) {
	var users []users.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *users.User) (*users.User, map[string]string) {
	var user users.User
	dbErr := map[string]string{}
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}
