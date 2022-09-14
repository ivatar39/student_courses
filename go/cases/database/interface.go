package database

import (
	"errors"
	"strconv"
)

type CRUD interface {
	Create(data ...string) error
	Get() []interface{}
	Update(interface{}) error
	Delete(id uint) error
}

type UserDB struct {
	conn *DataBase
}

type TODODB struct {
	conn *DataBase
}

func (c *UserDB) Create(data ...string) error {
	user := User{
		Name: data[0],
	}

	return c.conn.db.Create(&user).Error
}

func (c *UserDB) Get() []interface{} {
	var users []User
	c.conn.db.Preload("TODOs").Find(&users)
	userI := make([]interface{}, len(users))
	for i, user := range users {
		userI[i] = user
	}
	return userI
}

func (c *UserDB) Update(userT interface{}) error {
	user, ok := userT.(User)
	if !ok {
		return errors.New("not user type")
	}
	if err := c.conn.db.First(&user, user.ID).Error; err != nil {
		return err
	}

	return c.conn.db.Save(&user).Error
}

func (c *UserDB) Delete(id uint) error {
	var user User
	if err := c.conn.db.First(&user, id).Error; err != nil {
		return err
	}
	return c.conn.db.Delete(&user).Error
}

func (c *TODODB) Create(data ...string) error {
	userID, err := strconv.Atoi(data[2])
	if err != nil {
		return errors.New("can't get userID")
	}

	todo := TODO{
		Name:   data[0],
		Desc:   data[1],
		UserID: uint(userID),
	}

	return c.conn.db.Create(&todo).Error
}

func (c *TODODB) Get() []interface{} {
	var todos []TODO
	c.conn.db.Find(&todos)
	todoI := make([]interface{}, len(todos))
	for i, todo := range todos {
		todoI[i] = todo
	}
	return todoI
}

func (c *TODODB) Update(todoT interface{}) error {
	todo, ok := todoT.(TODO)
	if !ok {
		return errors.New("not user type")
	}
	if err := c.conn.db.First(&todo, todo.ID).Error; err != nil {
		return err
	}

	return c.conn.db.Save(&todo).Error
}

func (c *TODODB) Delete(id uint) error {
	var todo TODO
	if err := c.conn.db.First(&todo, id).Error; err != nil {
		return err
	}
	return c.conn.db.Delete(&todo).Error
}
