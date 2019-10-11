package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"goblog/myblog/dao"
	"goblog/myblog/models"
	"strconv"
)

func init() {
}

type ArticlesController struct {
	beego.Controller
}

// 分页控制
type ArticlesPageController struct {
	beego.Controller
}

// 获取文章总数目
type ArticlesCountController struct {
	beego.Controller
}

func getAllArticles(articlesController *ArticlesController) *models.Result {
	apiName := "获取所有文章"
	var result models.Result

	// 获取article表所有数据
	articleList, err := dao.GetAllArticles()
	if err != nil {
		logs.Error(apiName+":获取所有文章数据失败", err.Error())
		var result models.Result
		return result.Error()
	}

	// 封装vo
	articleVOList := make([]models.ArticleVO, 0)
	for _, article := range articleList {
		// 获取type
		articleType, err := dao.GetArticleTypeById(article.Id)
		if err != nil {
			return result.Error()
		}
		articleVO := models.ArticleVO{
			ArticleContent: article.ArticleContent,
			ArticleType:    articleType,
			Id:             article.Id,
			CreateTime:     article.CreateTime.Format("2006-01-02 15:04:05"),
			ArticleTitle:   article.ArticleTitle,
			CoverUrl:       article.CoverUrl,
			ViewNum:        article.ViewNum,
			CommentNum:     article.CommentNum,
			LikeNum:        article.LikeNum,
			ArticleAuthor:  article.ArticleAuthor,
		}
		articleVOList = append(articleVOList, articleVO)
	}

	// 返回结果
	return result.Success(articleVOList)
}

func getArticlesByPage(articlesPageController ArticlesPageController) *models.Result {
	apiName := "获取分页文章"
	var result models.Result
	limit, err := strconv.Atoi(articlesPageController.GetString("limit"))
	if err != nil {
		logs.Error(apiName + ":limit转换错误" + err.Error())
		return result.UnknownParam()
	}
	offset, err := strconv.Atoi(articlesPageController.GetString("offset"))
	if err != nil {
		logs.Error(apiName + ":offset转换错误" + err.Error())
		return result.UnknownParam()
	}
	// 分页获取article表所有数据
	articleList, err := dao.GetArticlesByPage(limit, offset)
	if err != nil {
		logs.Error(apiName+":获取分页文章数据失败", err.Error())
		var result models.Result
		return result.Error()
	}

	// 封装vo
	articleVOList := make([]models.ArticleVO, 0)
	for _, article := range articleList {
		// 获取type
		articleType, err := dao.GetArticleTypeById(article.Id)
		if err != nil {
			return result.Error()
		}

		articleVO := models.ArticleVO{
			ArticleContent:       article.ArticleContent,
			ArticleType:          articleType,
			Id:                   article.Id,
			CreateTime:           article.CreateTime.Format("2006-01-02 15:04:05"),
			ArticleTitle:         article.ArticleTitle,
			CoverUrl:             article.CoverUrl,
			ViewNum:              article.ViewNum,
			CommentNum:           article.CommentNum,
			LikeNum:              article.LikeNum,
			ArticleAuthor:        article.ArticleAuthor,
			ArticleSimpleContent: article.ArticleSimpleContent,
		}
		articleVOList = append(articleVOList, articleVO)
	}

	// 返回结果
	return result.Success(articleVOList)
}

func getArticleCount(articlesCountController ArticlesCountController) *models.Result {
	var result models.Result
	count, err := dao.GetArticleCount()
	if err != nil {
		logs.Error("获取文章总数失败", err.Error())
		return result.Error()
	}
	return result.Success(count)
}

// 获取所有文章
func (articlesController *ArticlesController) Get() {
	data := getAllArticles(articlesController)
	articlesController.Data["json"] = data
	articlesController.ServeJSON()
}

// 分页获取文章
func (articlesPageController ArticlesPageController) Get() {
	data := getArticlesByPage(articlesPageController)
	articlesPageController.Data["json"] = data
	articlesPageController.ServeJSON()
}

// 获取文章数量
func (articlesCountController ArticlesCountController) Get() {
	data := getArticleCount(articlesCountController)
	articlesCountController.Data["json"] = data
	articlesCountController.ServeJSON()
}
