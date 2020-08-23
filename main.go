package main

import (
	"golang_api/db"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	router.GET("/", func(ctx *gin.Context) {
		todos := db.dbGetAll()
		ctx.HTML(200, "index.html",gin.H{
			"tados": todos
		})
	})

	data := "in data"

	router.GET("/data", func(ctx *gin.Context) {
		ctx.HTML(200, "data.html", gin.H{"data": data})
	})

	router.Run()
}
