package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
)

func ModifyArticle(id, content string) error {
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	intId, err := strconv.Atoi(id)
	if err != nil {
		logs.Info("id转换错误")
		return err
	}
	article.Id = intId
	article.ArticleContent = content
	if num, err := o.Update(article, "ArticleContent"); err == nil {
		fmt.Println(num)
		return err
	}
	return nil
}

func GetOneArticle(id int) (*models.Article, error) {
	article := new(models.Article)
	o := orm.NewOrm()
	o.Using("default")
	article.Id = id

	// 获取article数据
	err := o.Read(article)
	if err != nil {
		return nil, err
	}
	return article, err
}

func ViewNumIncrease(id, viewNum int) error {
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	article.Id = id
	article.ViewNum = viewNum + 1
	_, err := o.Update(article, "view_num")
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func SearchByKeyWord(keyWord string) (*[]models.ArticleVO, error) {
	articles := new([]models.Article)
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	_, err = o.QueryTable("article").Filter("article_title__icontains", keyWord).All(articles)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	// article to articleVo
	articleVos := make([]models.ArticleVO, 0)
	for _, article := range *articles {
		var articleVo models.ArticleVO
		articleVo.Id = article.Id
		articleType, err := GetArticleTypeById(article.TypeId)
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
		articleVo.ArticleType = articleType
		articleVo.ViewNum = article.ViewNum
		articleVo.CoverUrl = article.CoverUrl
		articleVo.ArticleAuthor = article.ArticleAuthor
		articleVo.ArticleSimpleContent = article.ArticleSimpleContent
		articleVo.CommentNum = article.CommentNum
		articleVo.LikeNum = article.LikeNum
		articleVo.CreateTime = article.CreateTime.Format("2006-01-02 15:04:05")
		articleVo.ArticleTitle = article.ArticleTitle
		articleVos = append(articleVos, articleVo)

	}

	return &articleVos, err
}

func GetAllArticleByType(typeId int) (*[]models.ArticleVO, error) {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	articles := new([] models.Article)
	_, err = o.QueryTable("article").Filter("type_id", typeId).Filter("is_del", 0).All(articles)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	var articleVos []models.ArticleVO
	for _, article :=range *articles {
		var articleVo models.ArticleVO
		articleVo.ViewNum = article.ViewNum
		articleVo.CoverUrl = article.CoverUrl
		articleVo.ArticleAuthor = article.ArticleAuthor
		articleVo.ArticleSimpleContent = article.ArticleSimpleContent
		articleVo.CommentNum = article.CommentNum
		articleVo.LikeNum = article.LikeNum
		articleVo.CreateTime = article.CreateTime.Format("2006-01-02 15:04:05")
		articleVo.ArticleTitle = article.ArticleTitle
		articleVos = append(articleVos, articleVo)
	}

	return &articleVos, nil
}