package responses

import (
	"class-app/internal/api-student/datamodel"
	"time"
)

type Student struct {
	ID      int       `json:"id"`
	ClassId int       `json:"class_id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created" format:"date-time"`
	Updated time.Time `json:"updated" format:"date-time"`
}

func StudentSerializer(i any) any {
	studentData := i.(datamodel.Student)

	return &Student{
		ID:      studentData.ID,
		ClassId: studentData.ClassId,
		Name:    studentData.Name,
		Email:   studentData.Email,
		Created: studentData.Created,
		Updated: studentData.Updated,
	}
}
