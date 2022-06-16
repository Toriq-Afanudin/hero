package model

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "toriq:Ayu@1999@tcp(toriq1999.database.windows.net)/capstone")
	if err != nil {
		return nil, err
	}

	return db, nil
}
