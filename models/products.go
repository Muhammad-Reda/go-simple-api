package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id        int
	Code      string
	Name      string
	Category  string
	Price     int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
