package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"todolist/models"
	"todolist/responses"

	"github.com/gin-gonic/gin"
)

func GetAllActivity(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, title, email, created_at, updated_at FROM activities")
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
		defer rows.Close()

		// fmt.Println("rows:", rows)

		showAllActivity := []models.Activity{}
		for rows.Next() {
			var activity models.Activity
			var createdAtStr string
			var updatedAtStr string
			err := rows.Scan(&activity.Id, &activity.Title, &activity.Email, &createdAtStr, &updatedAtStr)
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
			activity.CreatedAt = createdAt
			activity.UpdatedAt = updatedAt

			showAllActivity = append(showAllActivity, activity)
		}

		ctx.JSON(
			http.StatusAccepted,
			responses.ActivityResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showAllActivity,
			},
		)
	}
}

func GetOneActivity(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activiyId := ctx.Param("id")
		rows := db.QueryRow("SELECT id, title, email, created_at, updated_at FROM activities WHERE id=?", activiyId)

		var activity models.Activity
		var createdAtStr string
		var updatedAtStr string
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Email, &createdAtStr, &updatedAtStr)
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

		showOneActivity := models.Activity{
			Id:        activity.Id,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.OneActivityResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneActivity,
			},
		)
	}
}

func CreateActivity(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var activity models.Activity
		if err := ctx.BindJSON(&activity); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := db.Exec("INSERT INTO activities (title, email, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", activity.Title, activity.Email)
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

		rows := db.QueryRow("SELECT id, title, email, created_at, updated_at FROM activities WHERE id=?", resId)

		var createdAtStr string
		var updatedAtStr string
		err = rows.Scan(&activity.Id, &activity.Title, &activity.Email, &createdAtStr, &updatedAtStr)
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

		showOneActivity := models.Activity{
			Id:        activity.Id,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.OneActivityResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneActivity,
			},
		)

	}
}

func UpdateActivity(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activiyId := ctx.Param("id")
		var activity models.Activity
		if err := ctx.BindJSON(&activity); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE activities SET title=?, email=? WHERE id=?", activity.Title, activity.Email, activiyId)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}

		rows := db.QueryRow("SELECT id, title, email, created_at, updated_at FROM activities WHERE id=?", activiyId)

		var createdAtStr string
		var updatedAtStr string
		err = rows.Scan(&activity.Id, &activity.Title, &activity.Email, &createdAtStr, &updatedAtStr)
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

		showOneActivity := models.Activity{
			Id:        activity.Id,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		ctx.JSON(
			http.StatusAccepted,
			responses.OneActivityResponse{
				Status:  "Success",
				Message: "Success",
				Data:    showOneActivity,
			},
		)

	}
}

func DeleteActivity(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activiyId := ctx.Param("id")
		_, err := db.Exec("DELETE FROM activities WHERE id=?", activiyId)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
		msg := fmt.Sprintf("Activity with ID %s Not Found", activiyId)
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": msg,
			},
		)

	}
}
