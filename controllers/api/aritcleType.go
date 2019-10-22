package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myblog/dao"
	"myblog/models"
)

type ArticleTypeController struct {
	beego.Controller
}

func (c *ArticleTypeController) Get() {
	types := c.getArticleTypeList()
	c.SetData(types)
	c.ServeJSON()
}

func (c *ArticleTypeController) getArticleTypeList() *models.Result{
	result := new(models.Result)
	types, err := dao.GetArticleTypeList()
	if err != nil {
		logs.Error("获取类型列表失败")
		return result.Error()
	}
	return result.Success(types)
}