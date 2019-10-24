package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myblog/dao"
	"myblog/models"
)

type UserSubmitController struct {
	beego.Controller
}

func (c *UserSubmitController) Post() {
	result := userSubmit(c)
	c.SetData(result)
	c.ServeJSON()
}

func userSubmit(c *UserSubmitController) *models.Result {
	result := new(models.Result)
	username := c.GetString("username")
	password := c.GetString("password")
	// 查询用户名是否已存在
	hasUser, err := dao.HasUser(username)
	if err != nil {
		logs.Error(err.Error())
		return result.Error()
	}
	if hasUser == true {
		return &models.Result{
			Code: 401,
			Msg:  "用户已存在",
			Data: nil,
		}
	}

	err = dao.UserSubmit(username, password)
	if err != nil {
		logs.Error("注册失败")
		return result.Error()
	}
	return result.Success(nil)
}
