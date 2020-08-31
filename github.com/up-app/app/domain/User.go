package domain

import "gorm.io/gorm"

// User Userエンティティ
type User struct {
	gorm.Model
	UserName string
	TimeCard []TimeCard
}
