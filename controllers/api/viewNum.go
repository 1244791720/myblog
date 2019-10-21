package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myblog/dao"
	"myblog/models"
)

type ViewNumController struct {
	beego.Controller
}

func (c *ViewNumController) Post() {
	result := c.ViewNumIncrease()
	c.SetData(result)
	c.ServeJSON()
}

func (c *ViewNumController) ViewNumIncrease() *models.Result {
	result := new(models.Result)
	// 根据文章id获取当前ViewNum
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err.Error())
		return result.Error()
	}
	article, err := dao.GetOneArticle(id)
	if err != nil {
		logs.Error(err.Error())
		return result.Error()
	}
	// veiwNum+1
	err = dao.ViewNumIncrease(article.Id, article.ViewNum)
	if err != nil {
		logs.Error("增加阅读数失败")
		return result.Error()
	}
	return result.Success(nil)
}
