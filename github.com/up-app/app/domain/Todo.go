package domain

import "gorm.io/gorm"

// Todo Todoエンティティ
type Todo struct {
	gorm.Model
	Text   string
	Status string
}
