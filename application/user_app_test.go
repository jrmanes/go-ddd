package application

import (
	"testing"

	"github.com/jrmanes/go-ddd/domain/entity/users"
	"github.com/stretchr/testify/assert"
)

//IF YOU HAVE TIME, YOU CAN TEST ALL THE METHODS FAILURES

var (
	saveUserRepo                func(*users.User) (*users.User, map[string]string)
	getUserRepo                 func(userId uint64) (*users.User, error)
	getUsersRepo                func() ([]users.User, error)
	getUserEmailAndPasswordRepo func(*users.User) (*users.User, map[string]string)
)

type fakeUserRepo struct{}

func (u *fakeUserRepo) SaveUser(user *users.User) (*users.User, map[string]string) {
	return saveUserRepo(user)
}
func (u *fakeUserRepo) GetUser(userId uint64) (*users.User, error) {
	return getUserRepo(userId)
}
func (u *fakeUserRepo) GetUsers() ([]users.User, error) {
	return getUsersRepo()
}
func (u *fakeUserRepo) GetUserByEmailAndPassword(user *users.User) (*users.User, map[string]string) {
	return getUserEmailAndPasswordRepo(user)
}

var userAppFake UserAppInterface = &fakeUserRepo{} //this is where the real implementation is swap with our fake implementation

func TestSaveUser_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	saveUserRepo = func(user *users.User) (*users.User, map[string]string) {
		return &users.User{
			IdUser:   1,
			Name:     "joseramon",
			UserName: "manes",
			Email:    "manes@example.com",
			Password: "password",
		}, nil
	}
	user := &users.User{
		IdUser:   1,
		Name:     "joseramon",
		UserName: "manes",
		Email:    "manes@example.com",
		Password: "password",
	}
	u, err := userAppFake.SaveUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "joseramon")
	assert.EqualValues(t, u.UserName, "manes")
	assert.EqualValues(t, u.Email, "manes@example.com")
}

func TestGetUser_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUserRepo = func(userId uint64) (*users.User, error) {
		return &users.User{
			IdUser:   1,
			Name:     "joseramon",
			UserName: "manes",
			Email:    "manes@example.com",
			Password: "password",
		}, nil
	}
	userId := uint64(1)
	u, err := userAppFake.GetUser(userId)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "joseramon")
	assert.EqualValues(t, u.UserName, "manes")
	assert.EqualValues(t, u.Email, "manes@example.com")
}

func TestGetUsers_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUsersRepo = func() ([]users.User, error) {
		return []users.User{
			{
				IdUser:   1,
				Name:     "joseramon",
				UserName: "manes",
				Email:    "manes@example.com",
				Password: "password",
			},
			{
				IdUser:   2,
				Name:     "kobe",
				UserName: "bryant",
				Email:    "kobe@example.com",
				Password: "password",
			},
		}, nil
	}
	users, err := userAppFake.GetUsers()
	assert.Nil(t, err)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUserEmailAndPasswordRepo = func(user *users.User) (*users.User, map[string]string) {
		return &users.User{
			IdUser:   1,
			Name:     "joseramon",
			UserName: "manes",
			Email:    "manes@example.com",
			Password: "password",
		}, nil
	}
	user := &users.User{
		IdUser:   1,
		Name:     "joseramon",
		UserName: "manes",
		Email:    "manes@example.com",
		Password: "password",
	}
	u, err := userAppFake.GetUserByEmailAndPassword(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "joseramon")
	assert.EqualValues(t, u.UserName, "manes")
	assert.EqualValues(t, u.Email, "manes@example.com")
}
