package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
)



type DeleteOneArticleController struct {
	beego.Controller
}

func (c *DeleteOneArticleController) Get() {
	result := deleteOneArticle(c)
	c.SetData(result)
	c.ServeJSON()
}

func deleteOneArticle(c *DeleteOneArticleController) *models.Result {
	result := new(models.Result)
	id := c.Input().Get("id")
	logs.Info("c.Input()", c.Input())
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		return result.Error()
	}
	logs.Info("id", id)
	article.IsDel = 1
	row, err := o.Update(article, "is_del")
	logs.Info("row", row)
	if err != nil || row == 0 {
		return result.Error()
	}
	return result.Success(nil)
}
