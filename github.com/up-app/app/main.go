package main

import (
	"strconv"

	"github.com/up-app/app/database"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	database.Init()

	// index
	router.GET("/", func(ctx *gin.Context) {
		todos := database.GetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		database.Insert(text, status)
		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := database.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{
			"todo": todo,
		})
	})

	// Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Erorr")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		database.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	// Delete_Check/:id
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Erorr")
		}
		todo := database.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{
			"todo": todo,
		})
	})

	// Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Erorr")
		}
		database.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
