package main

import (
	"net/http"
	"todolist/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	var r *gin.Engine
	r = gin.Default()
	routes.Activity(r, db)
	routes.Todo(r, db)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": "devcode updated",
			},
		)
	})
	r.Run(":3030")
	defer db.Close()

}
