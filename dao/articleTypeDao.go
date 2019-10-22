package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/models"
)

func GetArticleTypeById(id int) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	var articleType models.ArticleType
	err := o.QueryTable("article_type").One(&articleType)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return articleType.ArticleType, nil
}

func GetArticleTypeList() ([]models.ArticleType, error) {
	o := orm.NewOrm()
	o.Using("default")
	var articleType []models.ArticleType
	_, err := o.QueryTable("article_type").All(&articleType)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	return articleType, nil
}
