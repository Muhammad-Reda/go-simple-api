package models

import (
	"time"
)

type Transactions struct {
	Code       string
	Date       time.Time
	Quantity   int
	TotalPrice int64
	Discount   int
	Status     string
	Payment    string
	UserId     int64
	ProductId  int
}
