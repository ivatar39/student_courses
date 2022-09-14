package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	name string
	db   *gorm.DB
	User CRUD
	TODO CRUD
}

func NewDB() (*DataBase, error) { // подключение к бд
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		host, port, user, pass, sslmode, dbname)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn.AutoMigrate(&User{}, &TODO{})

	db := &DataBase{
		name: dbname,
		db:   conn,
	}

	db.User = &UserDB{
		conn: db,
	}
	db.TODO = &TODODB{
		conn: db,
	}
	return db, nil
}
