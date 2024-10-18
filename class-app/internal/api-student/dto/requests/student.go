package requests

type StudentCreate struct {
	ClassId int64  `json:"class_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type StudentUpdate struct {
	ClassId int64  `json:"class_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type StudentId struct {
	ID int64 `json:"id"`
}
