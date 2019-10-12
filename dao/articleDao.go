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
	if num, err := o.Update(&article); err == nil {
		fmt.Println(num)
		return err
	}
	return nil
}
