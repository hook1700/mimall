package models


type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 定义结构体操作的数据库表
func (User) TableName() string {
	return "user"
}
