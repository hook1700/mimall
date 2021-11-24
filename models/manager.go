package models

import (
_ "github.com/jinzhu/gorm"
)

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   int
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
}

func (Manager) TableName() string {
	return "manager"
}

