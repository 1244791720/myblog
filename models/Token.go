package models

import "time"

type Token struct {
	Id int64
	TokenValue string
	CreateTime time.Time
	UpdateTime time.Time
}

func (token *Token) TableName() string {
	return "token"
}
