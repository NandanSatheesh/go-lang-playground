package models

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type ErrorMessage struct {
	ErrorCode    uint   `json:"code"`
	ErrorMessage string `json:"message"`
}
