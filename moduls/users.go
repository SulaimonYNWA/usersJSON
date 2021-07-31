package moduls

import (
	"database/sql"
	"fmt"
	"log"
)


func AddUser(db *sql.DB){
	_, err := db.Query("INSERT INTO users(name,surname,age, sex, login, password, IsAdmin, remove) VALUES ('Dovud', 'Inomov', 22, 'female', 'dovkudidnomov', '28.92', false, false)")
	if err != nil {
		log.Fatal("could not insert a user", err)
	}
	fmt.Printf("inserted successfully\n")
}

