package persistence

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // it has to be here, if not cannot use postgres sql
	"github.com/jrmanes/go-ddd/domain/entity/users"
	"github.com/jrmanes/go-ddd/domain/repository"
)

// Here we will add the repositories that we have created
type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbHost, DbPort, DbName string) (*Repositories, error) {
	// create the uri for psql connection
	dbUri := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This apply migrations for all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&users.User{}).Error
}
