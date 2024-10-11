package requests

type ClassCreate struct {
	Number int    `json:"number"`
	Year   string `json:"year"`
}

type ClassUpdate struct {
	Number int    `json:"number"`
	Year   string `json:"year"`
}

type ClassId struct {
	ID int `json:"id"`
}
