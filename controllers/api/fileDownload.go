package api

import (
	"github.com/astaxie/beego"
)

type FileDownLoadController struct {
	beego.Controller
}

func (c *FileDownLoadController) Get() {
	//filename := c.Ctx.Input.Param("filename")
	file := `upload/1.ts`
	c.Ctx.Output.Download(file, "1.ts")
	c.Ctx.WriteString("sssssss")
}
