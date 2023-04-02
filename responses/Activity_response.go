package responses

import (
	"todolist/models"
)

type ActivityResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []models.Activity `json:"data"`
}

type OneActivityResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    models.Activity `json:"data"`
}
