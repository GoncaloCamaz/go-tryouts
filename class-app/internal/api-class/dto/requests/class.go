package requests

type ClassCreate struct {
	Number int64  `json:"number"`
	Year   string `json:"year"`
}

type ClassUpdate struct {
	Number int64  `json:"number"`
	Year   string `json:"year"`
}

type ClassId struct {
	ID int64 `json:"id"`
}
