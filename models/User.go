package models

type User struct {
	Id int64
	UserName string
	UserTypeId int64
	UserPwd string
	TokenId int64
}


func (user *User) TableName() string {
	return "user"
}
