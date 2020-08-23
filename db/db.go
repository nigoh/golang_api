package db

import (
	"golang_api/db/todo"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr（dbInit）")
	}
	db.AutoMigrate(&todo.Todo{})
	defer db.Close()
}

//DB追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr（dbInsert)")
	}
	db.Create(&todo.Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr（dbUpdate)")
	}
	todo := todo.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr（dbDelete)")
	}
	todo := todo.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func dbGetAll() []todo.Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr(dbGetAll())")
	}
	todos := []todo.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func dbGetOne(id int) todo.Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DataBase Open Erorr(dbGetOne())")
	}
	todo := todo.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
