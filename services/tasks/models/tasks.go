package models

type Task struct {
	ID          int    `json:"id"`
	Done        bool   `json:"done"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
