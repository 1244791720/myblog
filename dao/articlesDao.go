package dao

import (
	"github.com/astaxie/beego/orm"
	"goblog/myblog/models"
)

func GetAllArticles() ([]models.Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	articleList := make([]models.Article, 0)
	_, err := o.QueryTable("article").Filter("is_del", 0).All(&articleList)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

// 分页查询
// limit 每页多少个
// offset 当前第几页
func GetArticlesByPage(limit, offset int) ([]models.Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	articleList := make([]models.Article, 0)
	_, err := o.QueryTable("article").Filter("is_del", 0).OrderBy("-id").Limit(limit, offset).All(&articleList)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func GetArticleCount() (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	cnt, err := o.QueryTable("article").Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
