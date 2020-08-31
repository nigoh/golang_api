package database

import (
	"github.com/jinzhu/gorm"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"

	"github.com/up-app/app/domain"
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
	db.AutoMigrate(&domain.Todo{})
	defer db.Close()
}

//Insert DB追加
func Insert(text string, status string) {
	db := Open()
	db.Create(&domain.Todo{Text: text, Status: status})
	defer db.Close()
}

//Update DB更新
func Update(id int, text string, status string) {
	db := Open()
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//Delete DB削除
func Delete(id int) {
	db := Open()
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//GetAll DB全取得
func GetAll() []domain.Todo {
	db := Open()
	var todos []domain.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//GetOne DB一つ取得
func GetOne(id int) domain.Todo {
	db := Open()
	var todo domain.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
