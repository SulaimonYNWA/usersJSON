package mydatabase

import (
	"database/sql"
	"log"
)

func DBInit(db *sql.DB) {
	DDLs := []string{CreateUsersTable}
	for _, ddl:= range DDLs{
		_, err := db.Exec(ddl)
		if err!= nil{
			log.Fatalf("cant init... err is %e", err)
		}
	}
}
