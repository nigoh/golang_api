package tables

import "github.com/jinzhu/gorm"

// Todo TODOリスト
type Todo struct {
	gorm.Model
	Text   string
	Status string
}
