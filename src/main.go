package main

import (
	"github.com/gin-gonic/gin"
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
