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

type SubmitArticleController struct {
	beego.Controller
}
type Form struct {
	ArticleTitle string `form:"article_title"`
	ArticleSimpleContent string `form:"article_simple_content"`
	ArticleAuthor string `form:"article_author"`
	CoverUrl string
	ArticleContent string `form:"article_content"`
	TypeId int `form:"type_id"`
}

func (c *SubmitArticleController) Post() {
	data := submitArticle(c)
	c.Data["json"] = data
	c.ServeJSON()
}


func uploadFile(c *SubmitArticleController) *UploadResult {
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

func submitArticle(c *SubmitArticleController) *models.Result{
	var result models.Result

	// 获取表单数据
	form := new(Form)
	if err := c.ParseForm(form); err != nil {
		logs.Error("form转换错误" + err.Error())
		return result.Error()
	}
	uploadResult := uploadFile(c)
	if uploadResult.Success == 0 {
		logs.Error("上传失败--", uploadResult.Message)
		return result.Error()
	}
	form.CoverUrl = uploadFile(c).Url
	fmt.Println("form", form)
	// 创建一条数据
	//insertArticleDao(form)
	o := orm.NewOrm()
	o.Using("default")

	var cstZone = time.FixedZone("CST", 8*3600)
	currentTime := time.Now().In(cstZone)
	article := models.Article{
		TypeId:               form.TypeId,
		ArticleTitle:         form.ArticleTitle,
		ArticleContent:       form.ArticleContent,
		ArticleSimpleContent: form.ArticleSimpleContent,
		ArticleAuthor:        form.ArticleAuthor,
		IsDel:                0,
		CoverUrl:             form.CoverUrl,
		CreateTime:           currentTime,
		UpdateTime:           currentTime,
	}
	_, err := o.Insert(&article)
	if err != nil {
		logs.Error("新建文章失败" + err.Error())
		return result.Error()
	}
	return result.Success(nil)
}