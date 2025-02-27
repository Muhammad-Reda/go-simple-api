package models

import (
	"time"
)

type Transactions struct {
	Code        string
	Date        time.Time
	TotalPrice  int64
	Discount    int
	Status      string
	Payment     string
	ProductCode string
	UserId      int64
	Quantity    int
}
