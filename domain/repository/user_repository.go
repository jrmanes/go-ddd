package repository

import (
	"github.com/jrmanes/go-ddd/domain/entity/users"
)

// Repository: The UserRepository defines a collection of methods that the infrastructure implements.
// The methods are defined in an interface. These methods will later be implemented in the infrastructure layer.

type UserRepository interface {
	SaveUser(*users.User) (*users.User, map[string]string)
	GetUser(uint64) (*users.User, error)
	GetUsers() ([]users.User, error)
	GetUserByEmailAndPassword(*users.User) (*users.User, map[string]string)
}
