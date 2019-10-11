package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"goblog/myblog/controllers/api"
)

func init() {
	// 过滤器实现 put delete 请求
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = ctx.Input.Query("_method")
		}
	}

	// 路由
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
	beego.Router("/article/?:id", &api.ArticleController{})
	beego.Router("/articles", &api.ArticlesController{})
	beego.Router("/pageArticles", &api.ArticlesPageController{})
	beego.Router("/countArticles", &api.ArticlesCountController{})
	beego.Router("/download", &api.FileDownLoadController{})
	beego.Router("/upload", &api.FileUpLoadController{})
}
