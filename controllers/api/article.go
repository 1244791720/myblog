package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/myblog/models"
	"strconv"
)

func init() {
}

type ArticleController struct {
	beego.Controller
}

// 获取单个文章
func (this *ArticleController) Get() {
	apiName := "获取单个文章"
	var result models.Result

	// 从 url 获取 id
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(apiName+"id转换失败", err.Error())
		this.Data["json"] = result.Error()
		this.ServeJSON()
	}
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	article.Id = intId

	// 获取article数据
	err = o.Read(article)
	if err != nil {
		logs.Error(apiName+":获取单个文章数据失败", err.Error())
		this.Data["json"] = result.Error()
		this.ServeJSON()
	}

	// 返回结果
	this.Data["json"] = result.Success(article)
	this.ServeJSON()
}
