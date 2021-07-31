package main

import (
	"Database/moduls"
	"Database/mydatabase"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	Surname  string
	Age      int
	Sex      string
	Login    string
	Password string
	IsAdmin  bool
	Remove   bool
}

const (
	username = "Sulaimon"
	password = "liverpool19"
	hostname = "127.0.0.1:3306"
	dbname   = "test"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Users(w http.ResponseWriter, r *http.Request){
    // connect to DB
    db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

    //canceling connection if it takes much time
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
    //checking number of affected rows
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return 
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

    //creating tables if not exist
	mydatabase.DBInit(db)	
	
	moduls.AddUser(db)

    //reading data from from table
    users := User{}
	rows, err := db.Query(`select * from users where Id % 2 =0`)
	if err != nil {
		log.Println(err, `users are not selected`)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&users.Id,
			&users.Name,
			&users.Surname,
			&users.Age,
			&users.Sex,
			&users.Login,
			&users.Password,
			&users.IsAdmin,
			&users.Remove,
		)
		if err != nil {
			log.Fatal(err, ` not selected next`)
		}   
        unmarshaled := []User{users}
		Marshaled, _ := json.Marshal(unmarshaled)
		fmt.Println(string(Marshaled))
        
        fmt.Fprintf(w, string(Marshaled))
	}
}

func setupRoutes(){
    http.HandleFunc("/users", Users)
    http.ListenAndServe(":3001", nil)
   return
}

func main() {
	setupRoutes()
}
