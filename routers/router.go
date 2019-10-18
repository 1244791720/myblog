package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myblog/controllers/api"
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
	// 文章编辑器保存文章
	beego.Router("/submitArticle", &api.SubmitArticleController{})
	// 表单修改文章内容
	beego.Router("/articleForm", &api.ArticleFormController{})
	// 删除一篇文章
	beego.Router("deleteOneArticle", &api.DeleteOneArticleController{})
	// 修改一个文章封面图
	beego.Router("updateOneArticleCover", &api.UpdateOneArticleCoverController{})


}
