package responses

import (
	"todolist/models"
)

type TodoResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []models.Todo `json:"data"`
}

type OneTodoResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    models.Todo `json:"data"`
}

type AltToDoResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    models.AltTodo `json:"data"`
}
