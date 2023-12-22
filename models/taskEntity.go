package models

type TaskEntity struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
