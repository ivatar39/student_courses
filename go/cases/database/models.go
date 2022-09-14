package database

import "gorm.io/gorm"

type TODO struct {
	gorm.Model        // хранит в себе структуру, которая содержит id, created_at, updated_at, deleted_at
	Name       string `gorm:"not null;"`
	Desc       string `gorm:"not null;"`
	UserID     uint
}

type User struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	TODOs []TODO `gorm:"foreignKey:UserID"`
}
