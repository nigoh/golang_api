package main

import (
	"strconv"

	"local.packages/db"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	db.Init()

	// index
	router.GET("/", func(ctx *gin.Context) {
		todos := db.GetAll()
		ctx.HTML(200, "index.html", gin.H{
			"tados": todos,
		})
	})

	// Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		db.Insert(text, status)
		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := db.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{
			"tado": todo,
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
		db.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	// Delete_Check/:id
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Erorr")
		}
		todo := db.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{
			"tado": todo,
		})
	})

	// Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Erorr")
		}
		db.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
