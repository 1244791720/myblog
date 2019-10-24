package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myblog/controllers/api"
	"myblog/dao"
	"strconv"
)

func init() {
	// 过滤器实现 put delete 请求
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = ctx.Input.Query("_method")
		}
	}
	// 过滤器实现token 认证
	var FilterToken = func(ctx *context.Context) {
		token := ctx.Input.Cookie("token")
		userTypeId := ctx.Input.Cookie("user_type_id")
		intId, err := strconv.Atoi(userTypeId)
		if err != nil {
			ctx.Redirect(302, "/adminLogin.html")
		}
		isRight := dao.IsRightToken(token)
		// 用户存在且身份为管理员，除此之外登录不进
		if !(isRight && intId == 2) {
			ctx.Redirect(302, "/adminLogin.html")
		}
	}

	// 过滤器
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
	beego.InsertFilter("/admin.html", beego.BeforeStatic, FilterToken)
	// 路由
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
	beego.Router("/deleteOneArticle", &api.DeleteOneArticleController{})
	// 修改一个文章封面图
	beego.Router("/updateOneArticleCover", &api.UpdateOneArticleCoverController{})
	// 登录
	beego.Router("/login", &api.LoginController{})
	// 文章阅读数加一
	beego.Router("/viewNum", &api.ViewNumController{})
	// 获取文章类型列表
	beego.Router("/articleType", &api.ArticleTypeController{})
	// 按照关键词分页获取文章
	beego.Router("/searchByKeyWord", &api.SearchByKeyWordController{})

}
