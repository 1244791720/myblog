package api

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"os"
	"path"
	"time"
)

// 文件上传返回结果
type UploadResult struct {
	Success int    `json:"success"` // 0 表示上传失败，1 表示上传成功
	Message string `json:"message"` //提示的信息，上传成功或上传失败及错误信息等
	Url     string `json:"url"`     // 上传成功时才返回
}

func (result *UploadResult)uploadFailed(message string) *UploadResult {
	result.Success =  0
	result.Message =  message
	result.Url = ""
	return result
}
func (result *UploadResult)uploadSucceed(url string) *UploadResult {
	result.Success =  1
	result.Message =  "上传成功"
	result.Url = url
	return result
}

type FileUpLoadController struct {
	beego.Controller
}

func (c *FileUpLoadController) Post() {
	data := uploadDao(c)
	c.Data["json"] = data
	c.ServeJSON()
}

func uploadDao(c *FileUpLoadController) *UploadResult {
	file, fileHeader, err := c.GetFile("editormd-image-file")
	if err != nil {
		logs.Error("文件上传失败")
		return new(UploadResult).uploadFailed("文件上传失败")
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
		return new(UploadResult).uploadFailed("后缀名不符合上传要求")
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err = os.MkdirAll( uploadDir , 777)
	if err != nil {
		return new(UploadResult).uploadFailed("文件上传失败")
	}

	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	fileName := fmt.Sprintf("%x",hashName) + ext

	// 上传
	fpath := uploadDir + fileName
	defer file.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("editormd-image-file", fpath)
	if err != nil {
		logs.Error("saveToFile：", err)
		return new(UploadResult).uploadFailed("文件上传失败")
	}
	return new(UploadResult).uploadSucceed(c.Ctx.Request.RemoteAddr + fpath)
}