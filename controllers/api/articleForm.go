package api

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"myblog/models"
	"os"
	"path"
	"time"
)

type ArticleFormController struct {
	beego.Controller
}
type form struct {
	Id int `formL"id"`
	ArticleTitle string `form:"article_title"`
	ArticleSimpleContent string `form:"article_simple_content"`
	ArticleAuthor string `form:"article_author"`
	CoverUrl string
	TypeId int `form:"type_id"`
}

func (c *ArticleFormController) Post() {
	data := modifyArticleByFormData(c)
	c.Data["json"] = data
	c.ServeJSON()
}

func modifyArticleByFormData(c *ArticleFormController) *models.Result {
	result := new(models.Result)
	// 获取表单数据
	form := new(form)
	if err := c.ParseForm(form); err != nil {
		logs.Error("form转换错误" + err.Error())
		return result.Error()
	}
	//uploadResult := formDataUplodaFile(c)
	//if uploadResult.Success == 0 {
	//	logs.Error("上传失败--", uploadResult.Message)
	//	return result.Error()
	//}
	//form.CoverUrl = formDataUplodaFile(c).Url
	article := new(models.Article)
	article.Id = form.Id
	article.ArticleAuthor = form.ArticleAuthor
	article.ArticleSimpleContent = form.ArticleSimpleContent
	article.ArticleTitle = form.ArticleTitle
	article.TypeId = form.TypeId
	logs.Info("article", article)
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(article, "article_author", "article_simple_content", "article_title", "type_id")
	if err != nil {
		logs.Error("更新失败")
		return result.Error()
	}
	return result.Success(nil)
}


func formDataUplodaFile(c *ArticleFormController) *UploadResult {
	file, fileHeader, err := c.GetFile("coverImage")
	if err != nil {
		logs.Error("文件上传失败")
		return new(UploadResult).uploadFailed("文件上传失败"+ err.Error())
	}

	//后缀名 如：.jpg
	ext := path.Ext(fileHeader.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":true,
		".jpeg":true,
		".png":true,
	}
	if _,ok:=AllowExtMap[ext];!ok{
		return new(UploadResult).uploadFailed("后缀名不符合上传要求" + err.Error())
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err = os.MkdirAll( uploadDir , 777)
	if err != nil {
		return new(UploadResult).uploadFailed("文件上传失败"+ err.Error())
	}

	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	fileName := fmt.Sprintf("%x",hashName) + ext

	// 上传
	fpath := uploadDir + fileName
	defer file.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("coverImage", fpath)
	if err != nil {
		logs.Error("saveToFile：", err)
		return new(UploadResult).uploadFailed("文件上传失败" + err.Error())
	}
	return new(UploadResult).uploadSucceed(fpath)
}
