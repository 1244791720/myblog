package models

import "time"

type Article struct {
	Id int
	// typeid
	TypeId               int
	ArticleTitle         string
	ArticleContent       string
	ArticleSimpleContent string
	// 作者
	ArticleAuthor string
	CreateTime    time.Time
	UpdateTime    time.Time
	IsDel         int
	CoverUrl      string
	// 浏览数
	ViewNum int
	// 评论数
	CommentNum int
	// 喜欢
	LikeNum int
}

type ArticleVO struct {
	Id int
	// typeid
	ArticleType          string
	ArticleTitle         string
	ArticleSimpleContent string
	ArticleContent       string
	CreateTime           string
	IsDel                int
	CoverUrl             string
	// 浏览数
	ViewNum int
	// 评论数
	CommentNum int
	// 喜欢
	LikeNum int
	// 作者
	ArticleAuthor string
}

func (article *Article) TableName() string {
	return "article"
}
