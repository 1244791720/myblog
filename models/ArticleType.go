package models

type ArticleType struct {
	Id          int
	ArticleType string
}

func (articleType *ArticleType) TableName() string {
	return "article_type"
}
