package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myblog/dao"
	"myblog/models"
)

type SearchByKeyWordController struct {
	beego.Controller
}

func (c *SearchByKeyWordController) Get() {
	result := c.searchByKeyWord()
	c.SetData(result)
	c.ServeJSON()
}

func (c *SearchByKeyWordController) searchByKeyWord() *models.Result {
	keyWord := c.GetString("keyWord")
	articles, err := dao.SearchByKeyWord(keyWord)
	result := new(models.Result)
	if err != nil {
		logs.Error(err.Error())
		return result.Error()
	}
	return result.Success(articles)
}
