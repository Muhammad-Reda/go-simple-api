package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:admin@/api1?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("can not connect to database: %v", err)
	}

	fmt.Println("Access to database")

	return db, nil
}
