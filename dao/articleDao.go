package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/controllers/api"
	"myblog/models"
	"strconv"
)

func ModifyArticle(param api.Param) error {
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	id, err := strconv.Atoi(param.Id)
	if err != nil {
		logs.Info("id转换错误")
		return err
	}
	article.Id = id
	article.ArticleContent = param.Content
	if num, err := o.Update(&article); err == nil {
		fmt.Println(num)
		return err
	}
	return nil
}
