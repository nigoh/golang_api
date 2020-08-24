package db

import (
	"local.packages/tables"

	"github.com/jinzhu/gorm"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//Open DB
func Open() *gorm.DB {
	driver := "sqlite3"
	dbname := "test.sqlite3"
	db, err := gorm.Open(driver, dbname)
	if err != nil {
		panic("DataBase Open Erorr（dbInit）")
	}
	return db
}

//Init DB初期化
func Init() {
	db := Open()
	db.AutoMigrate(&tables.Todo{})
	defer db.Close()
}

//Insert DB追加
func Insert(text string, status string) {
	db := Open()
	db.Create(&tables.Todo{Text: text, Status: status})
	defer db.Close()
}

//Update DB更新
func Update(id int, text string, status string) {
	db := Open()
	var todo tables.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//Delete DB削除
func Delete(id int) {
	db := Open()
	var todo tables.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//GetAll DB全取得
func GetAll() []tables.Todo {
	db := Open()
	var todos []tables.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//GetOne DB一つ取得
func GetOne(id int) tables.Todo {
	db := Open()
	var todo tables.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
