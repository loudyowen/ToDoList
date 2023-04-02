package controllers

import (
	"database/sql"
	"net/http"
	"time"
	"todolist/models"
	"todolist/responses"

	"github.com/gin-gonic/gin"
)

func GetAllTodolist(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activityGroupId := ctx.Query("activity_group_id")
		var rows *sql.Rows
		var err error
		if activityGroupId != "" {
			rows, err = db.Query("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo_items WHERE activity_group_id=?", activityGroupId)
			if err != nil {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{"error": err.Error()},
				)
				return
			}
		} else {
			rows, err = db.Query("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo_items")
			if err != nil {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{"error": err.Error()},
				)
				return
			}
		}
		defer rows.Close()

		showAllTodoItems := []models.Todo{}
		for rows.Next() {
			var todo_item models.Todo
			var createdAtStr string
			var updatedAtStr string
			err := rows.Scan(&todo_item.Id, &todo_item.ActivityGroupId, &todo_item.Title, &todo_item.IsActive, &todo_item.Priority, &createdAtStr, &updatedAtStr)
			if err != nil {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{"error": err.Error()},
				)
				return
			}

			// parsing waktu supaya bisa dibaca oleh golang struct time.Time
			createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
			if err != nil {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{"error": err},
				)
				return
			}
			updatedAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
			if err != nil {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{"error": err},
				)
				return
			}
			todo_item.CreatedAt = createdAt
			todo_item.UpdatedAt = updatedAt

			showAllTodoItems = append(showAllTodoItems, todo_item)
		}

		ctx.JSON(
			http.StatusAccepted,
			responses.TodoResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showAllTodoItems,
			},
		)
	}
}
func GetOneTodolist(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todoId := ctx.Param("id")
		rows := db.QueryRow("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo_items WHERE id=?", todoId)

		var todo_item models.Todo
		var createdAtStr string
		var updatedAtStr string
		err := rows.Scan(&todo_item.Id, &todo_item.ActivityGroupId, &todo_item.Title, &todo_item.IsActive, &todo_item.Priority, &createdAtStr, &updatedAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}

		// parsing waktu supaya bisa dibaca oleh golang struct time.Time
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}
		updatedAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}

		showOneTodoItem := models.Todo{
			Id:              todo_item.Id,
			Title:           todo_item.Title,
			ActivityGroupId: todo_item.ActivityGroupId,
			IsActive:        todo_item.IsActive,
			Priority:        todo_item.Priority,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.OneTodoResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneTodoItem,
			},
		)
	}
}
func CreateTodolist(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var todo_item models.AltTodo
		if err := ctx.BindJSON(&todo_item); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := db.Exec("INSERT INTO todo_items (activity_group_id, title, is_active, priority, created_at, updated_at) VALUES (?, ?, ?, 'very-high', NOW(), NOW())", todo_item.ActivityGroupId, todo_item.Title, todo_item.IsActive)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}

		resId, err := res.LastInsertId()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		rows := db.QueryRow("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo_items WHERE id=?", resId)

		// var todo_item models.Todo
		var createdAtStr string
		var updatedAtStr string
		err = rows.Scan(&todo_item.Id, &todo_item.ActivityGroupId, &todo_item.Title, &todo_item.IsActive, &todo_item.Priority, &createdAtStr, &updatedAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}

		// parsing waktu supaya bisa dibaca oleh golang struct time.Time
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}
		updatedAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}

		showOneTodoItem := models.AltTodo{
			Id:              todo_item.Id,
			Title:           todo_item.Title,
			ActivityGroupId: todo_item.ActivityGroupId,
			IsActive:        todo_item.IsActive,
			Priority:        todo_item.Priority,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.AltToDoResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneTodoItem,
			},
		)
	}
}
func UpdateTodolist(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todoId := ctx.Param("id")
		var todo_item models.Todo
		if err := ctx.BindJSON(&todo_item); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE todo_items SET title=?, priority=?, is_active=? WHERE id=?", todo_item.Title, todo_item.Priority, todo_item.IsActive, todoId)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}

		rows := db.QueryRow("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo_items WHERE id=?", todoId)

		// var todo_item models.Todo
		var createdAtStr string
		var updatedAtStr string
		err = rows.Scan(&todo_item.Id, &todo_item.ActivityGroupId, &todo_item.Title, &todo_item.IsActive, &todo_item.Priority, &createdAtStr, &updatedAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}

		// parsing waktu supaya bisa dibaca oleh golang struct time.Time
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}
		updatedAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}

		showOneTodoItem := models.Todo{
			Id:              todo_item.Id,
			Title:           todo_item.Title,
			ActivityGroupId: todo_item.ActivityGroupId,
			IsActive:        todo_item.IsActive,
			Priority:        todo_item.Priority,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.OneTodoResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneTodoItem,
			},
		)
	}
}
func DeleteTodolist(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todoId := ctx.Param("id")
		_, err := db.Exec("DELETE FROM activities WHERE id=?", todoId)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Success",
				"message": "Success",
				"data":    []models.Todo{},
			},
		)

	}
}
