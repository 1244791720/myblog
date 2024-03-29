package api

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myblog/dao"
	"myblog/models"
	"strconv"
)

func init() {
}

type ArticleController struct {
	beego.Controller
}

type ArticleVO struct {
	Id int
	// typeid
    TypeId               int `json:"type_id"`
	ArticleTitle         string `json:"article_title"`
	ArticleSimpleContent string `json:"article_simple_content"`
	ArticleContent       string `json:"article_content"`
	CreateTime           string `json:"create_time"`
	IsDel                int`json:"is_del"`
	CoverUrl             string `json:"cover_url"`
	// 浏览数
	ViewNum int `json:"view_num"`
	// 评论数
	CommentNum int `json:"comment_num"`
	// 喜欢
	LikeNum int `json:"like_num"`
	// 作者
	ArticleAuthor string `json:"article_author"`
}

type Param struct {
	Id string `json:"id"`
	Content string `json:"content"`
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

	// to vo
	vo := new(ArticleVO)
	vo.CreateTime = article.CreateTime.Format("2006-01-02 15:04:05")
	vo.ArticleContent = article.ArticleContent
	vo.ArticleSimpleContent = article.ArticleSimpleContent
	vo.ArticleAuthor = article.ArticleAuthor
	vo.TypeId = article.TypeId
	vo.ArticleTitle = article.ArticleTitle
	vo.ViewNum = article.ViewNum
	// 返回结果
	this.Data["json"] = result.Success(vo)
	this.ServeJSON()
}

// 修改单个文章内容
func (this *ArticleController) Post() {
	data := modifyArticle(this)
	this.Data["json"] = data
	this.ServeJSON()
}

func modifyArticle(this *ArticleController) *models.Result{
	var param Param
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	if  err != nil {
		logs.Error("保存文章失败", err.Error())
	}
	logs.Info(param)
	err = dao.ModifyArticle(param.Id, param.Content)
	if err != nil {
		return new(models.Result).Error()
	}
	return new(models.Result).Success(param)
}