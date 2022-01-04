package models


type Role struct {
	Id       int
	Status int
	AddTime  int
	Title string
	Description string


}

func (Role) TableName() string {
	return "role"
}