// server.go initializes the database connection, the routes and starts the server on a specified port.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controllers "github.com/dovudwkt/playground/mvc_server/controllers"
	user_model "github.com/dovudwkt/playground/mvc_server/models/user"
	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	username, password, host, dbName string
	port                             int
}
type serverConfig struct {
	host string
	port int
}
type Config struct {
	Server serverConfig
	DB     dbConfig
}

// set up configuration values
var cfg = Config{
	Server: serverConfig{
		host: "127.0.0.1",
		port: 3001,
	},
	DB: dbConfig{
		username: "root",
		password: "password",
		host:     "127.0.0.1",
		dbName:   "playground",
		port:     3306,
	},
}

func dsn(cfg dbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.username, cfg.password, cfg.host, cfg.port, cfg.dbName)
}

func main() {

	// init database connection
	log.Print("Starting mysql server at port ", cfg.DB.port)
	db, err := sql.Open("mysql", dsn(cfg.DB))
	if err != nil {
		log.Printf("Open sql: %s\n", err)
		return
	}
	defer db.Close()

	// init models
	userModel := user_model.NewModel(user_model.Options{DB: db})

	// -------------------------------------------------------------------
	// HTTP router
	r := mux.NewRouter()

	// Users
	r.HandleFunc("/users", (&controllers.UserController{UserModel: userModel}).AddUser).Methods("POST")
	r.HandleFunc("/users", (&controllers.UserController{UserModel: userModel}).GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", (&controllers.UserController{UserModel: userModel}).GetByID).Methods("GET")

	// Health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// --------------------------------------------------------------------

	var addr = fmt.Sprintf(":%d", cfg.Server.port)
	log.Print("Server running on port ", cfg.Server.port)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Print("http listen and serve: ", err.Error())
		return
	}

}
