package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myblog/dao"
	"myblog/models"
)

type GetAllArticlesByTypeController struct {
	beego.Controller
}

func (c *GetAllArticlesByTypeController) Get() {
	result := c.getAllArticles()
	c.SetData(result)
	c.ServeJSON()
}

func (c *GetAllArticlesByTypeController) getAllArticles() *models.Result {
	result := new(models.Result)
	articleTypeId, err :=  c.GetInt("articleTypeId")
	if err !=nil {
		logs.Error("获取articleType失败")
		return result.Error()
	}

	contents, err := dao.GetAllArticleByType(articleTypeId)
	if err != nil {
		logs.Error("获取文章失败")
	}
	return result.Success(contents)

}
