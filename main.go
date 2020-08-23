package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	data := "in data"

	router.GET("/data", func(ctx *gin.Context) {
		ctx.HTML(200, "data.html", gin.H{"data": data})
	})

	router.Run()
}
