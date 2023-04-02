package routes

import (
	"database/sql"
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func Activity(r *gin.Engine, db *sql.DB) {
	r.GET("/activity-groups", controllers.GetAllActivity(db))
	r.GET("/activity-groups/:id", controllers.GetOneActivity(db))
	r.POST("/activity-groups", controllers.CreateActivity(db))
	r.PATCH("/activity-groups/:id", controllers.UpdateActivity(db))
	r.DELETE("/activity-groups/:id", controllers.DeleteActivity(db))
}
