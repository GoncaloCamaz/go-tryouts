package responses

import (
	"time"
)

type Student struct {
	ID          int       `json:"id"`
	ClassId     int       `json:"class_id"`
	ClassNumber int       `json:"class_number"`
	ClassYear   string    `json:"class_year"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
}
