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
	"strconv"
	"time"
)

type UpdateOneArticleCoverController struct {
	beego.Controller
}
func (c *UpdateOneArticleCoverController) Post() {
	var result *models.Result
	uploadResult := uploadOneArticelCoverFile(c)
	logs.Info("c.INput", c.Input())
	id := c.Input().Get("id")
	o := orm.NewOrm()
	o.Using("default")
	article := new(models.Article)
	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		logs.Error(err.Error())
		result = result.Error()
		c.SetData(result)
		c.ServeJSON()
	}
	article.CoverUrl = uploadResult.Url
	row, err :=o.Update(article, "cover_url")
	if err!=nil || row == 0 {
		logs.Error(err.Error())
		result = result.Error()
		c.SetData(result)
		c.ServeJSON()
	}
	result = result.Success(nil)
	c.SetData(result)
	c.ServeJSON()
}

func uploadOneArticelCoverFile(c *UpdateOneArticleCoverController) *UploadResult {
	file, fileHeader, err := c.GetFile("article_cover")
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
	err = c.SaveToFile("article_cover", fpath)
	if err != nil {
		logs.Error("saveToFile：", err)
		return new(UploadResult).uploadFailed("文件上传失败" + err.Error())
	}
	return new(UploadResult).uploadSucceed(fpath)
}
