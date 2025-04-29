package domain

import (
	"time"
)

type User struct {
	ID         int64
	Lang       string
	ExpiryDate time.Time
}
