package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Email     string
	Username  string
	Password  string
	Address   string
	Telephone string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
