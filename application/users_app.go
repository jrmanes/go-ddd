package application

import (
	"github.com/jrmanes/go-ddd/domain/entity/users"
	"github.com/jrmanes/go-ddd/domain/repository"
)

// We have successfully defined the API business logic in our domain.
// The application connex-special/nautilus-clipboard
// The above have methods to save and retrieve user data.
// The UserApp struct has the UserRepository interface, which made it possible to call the user repository methods.

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*users.User) (*users.User, map[string]string)
	GetUsers() ([]users.User, error)
	GetUser(uint64) (*users.User, error)
	GetUserByEmailAndPassword(*users.User) (*users.User, map[string]string)
}

func (u *userApp) SaveUser(user *users.User) (*users.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(userId uint64) (*users.User, error) {
	return u.us.GetUser(userId)
}

func (u *userApp) GetUsers() ([]users.User, error) {
	return u.us.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *users.User) (*users.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}
