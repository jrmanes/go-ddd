package users

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jrmanes/go-ddd/infrastructure/security"
)

// The Users struct
type User struct {
	IdUser    uint64     `gorm:"primary_key;auto_increment" json:"id_user"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	UserName  string     `gorm:"size:100;not null;" json:"user_name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null;" json:"password"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PublicUser struct {
	IdUser   uint64 `gorm:"primary_key;auto_increment" json:"id_user"`
	Name     string `gorm:"size:100;not null;" json:"name"`
	UserName string `gorm:"size:100;not null;" json:"user_name"`
}

// TODO -> Create a function for password hash before added it to database.
//BeforeSave is a gorm hook
func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

type Users []User

//So that we dont expose the user's email address and password to the world
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

//So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		IdUser:   u.IdUser,
		Name:     u.Name,
		UserName: u.UserName,
	}
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET USER")
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	// var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		//if u.Email != "" {
		//	if err = u.UserName.checkmail.ValidateFormat(u.Email); err != nil {
		//		errorMessages["invalid_email"] = "email email"
		//	}
		//}

	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		//if u.Email != "" {
		//	if err = checkmail.ValidateFormat(u.Email); err != nil {
		//		errorMessages["invalid_email"] = "please provide a valid email"
		//	}
		//}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		//if u.Email != "" {
		//	if err = checkmail.ValidateFormat(u.Email); err != nil {
		//		errorMessages["invalid_email"] = "please provide a valid email"
		//	}
		//}
	default:
		if u.UserName == "" {
			errorMessages["username_required"] = "UserName is required"
		}
		if u.Name == "" {
			errorMessages["name_required"] = " name is required"
		}
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["invalid_password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		//if u.Email != "" {
		//	if err = checkmail.ValidateFormat(u.Email); err != nil {
		//		errorMessages["invalid_email"] = "please provide a valid email"
		//	}
		//}
	}
	return errorMessages
}
