package model

//Values for status of todo
const (
	TodoValue uint = iota
	OngoingValue
	CompletedValue
)

//Todo model for todo type
type Todo struct {
	ID           uint64 `json:"id"`
	Description  string `json:"description"`
	Status       uint   `json:"status"`
	CompleteDate string `json:"end-date"`
}

//TodoResponse model for todo type
type TodoResponse struct {
	ID           uint64 `json:"id"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	CompleteDate string `json:"end-date"`
}
