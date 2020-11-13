package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jrmanes/go-ddd/infrastructure/auth"
	"github.com/jrmanes/go-ddd/infrastructure/persistence"
	"github.com/jrmanes/go-ddd/interfaces"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home path")
}

func init() {
	// Load .env
	//err := godotenv.Load(os.ExpandEnv("/app/.env")) // this is another way in order to look for it in an specific path
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file... ERROR: ", err)
	}
}

func main() {
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	// Connect to DB
	services, err := persistence.NewRepositories(dbdriver, user, password, host, dbport, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate() // execute migrations into db

	tk := auth.NewToken()
	// We define the interfaces for all our services
	users := interfaces.NewUsers(services.User)
	authentication := interfaces.NewAuthenticate(services.User, tk)
	// Load Port from .env for application server
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	//r := mux.NewRouter()
	//r.HandleFunc("/", Home).Methods("GET")

	r := gin.Default()

	//user routes
	r.POST("/users", users.SaveUser)
	r.GET("/users", users.GetUsers)
	r.GET("/users/:id_user", users.GetUser)

	r.POST("/login", authentication.Login)
	r.POST("/logout", authentication.Logout)
	r.POST("/refresh", authentication.Refresh)

	//r.HandleFunc("/users", users.SaveUser).Methods("POST")
	//r.HandleFunc("/users", users.GetUsers).Methods("GET")
	//r.HandleFunc("/users/:user_id", users.GetUser).Methods("GET")

	log.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
