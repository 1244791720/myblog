package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
)

func ModifyArticle(id,content string) error {
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	intId, err := strconv.Atoi(id)
	if err != nil {
		logs.Info("id转换错误")
		return err
	}
	article.Id = intId
	article.ArticleContent = content
	if num, err := o.Update(article, "ArticleContent"); err == nil {
		fmt.Println(num)
		return err
	}
	return nil
}

func GetOneArticle(id int) (*models.Article, error) {
	article := new(models.Article)
	o := orm.NewOrm()
	o.Using("default")
	article.Id = id

	// 获取article数据
	err := o.Read(article)
	if err != nil {
		return nil, err
	}
	return article, err
}

func ViewNumIncrease(id, viewNum int) error {
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	article.Id = id
	article.ViewNum = viewNum + 1
	_, err := o.Update(article, "view_num")
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

