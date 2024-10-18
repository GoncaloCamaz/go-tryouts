package responses

import (
	"time"
)

type Student struct {
	ID          int64     `json:"id"`
	ClassId     int64     `json:"class_id"`
	ClassNumber int64     `json:"class_number"`
	ClassYear   string    `json:"class_year"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
}
