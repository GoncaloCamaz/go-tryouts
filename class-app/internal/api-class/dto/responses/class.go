package responses

import (
	"class-app/internal/api-class/datamodel"
	"time"
)

type Class struct {
	ID      int64     `json:"id"`
	Number  int64     `json:"number"`
	Year    string    `json:"year"`
	Created time.Time `json:"created" format:"date-time"`
	Updated time.Time `json:"updated" format:"date-time"`
}

func ClassSerializer(i any) any {
	data := i.(datamodel.Class)
	return &Class{
		ID:      data.ID,
		Number:  data.Number,
		Year:    data.Year,
		Created: data.Created,
		Updated: data.Updated,
	}
}
