package routes

import (
	"database/sql"
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func Todo(r *gin.Engine, db *sql.DB) {
	r.GET("/todo-items", controllers.GetAllTodolist(db))
	r.GET("/todo-items/:id", controllers.GetOneTodolist(db))
	r.POST("/todo-items", controllers.CreateTodolist(db))
	r.PATCH("/todo-items/:id", controllers.UpdateTodolist(db))
	r.DELETE("/todo-items/:id", controllers.DeleteTodolist(db))
}
