package requests

type StudentCreate struct {
	ClassId int    `json:"class_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type StudentUpdate struct {
	ClassId int    `json:"class_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type StudentId struct {
	ID int `json:"id"`
}
