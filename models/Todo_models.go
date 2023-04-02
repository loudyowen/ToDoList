package models

import (
	"time"
)

type Todo struct {
	Id              int       `json:"id"`
	ActivityGroupId int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type AltTodo struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupId int       `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
