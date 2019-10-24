package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myblog/models"
)

type LoginController struct {
	beego.Controller
}

type ResultData struct {
	TokenValue string
	UserTypeId int64
}

func (c *LoginController) Post() {
	result := c.login()
	c.SetData(result)
	c.ServeJSON()
}

func (c *LoginController) login() *models.Result {
	var result models.Result
	username := c.GetString("username")
	password := c.GetString("password")
	user := new(models.User)
	user.UserName = username
	user.UserPwd = password
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		logs.Error(err.Error())
		return result.Error()
	}

	err = o.Read(user, "user_name", "user_pwd")
	if err != nil {
		logs.Error("根据用户名密码查询User:" + err.Error())
	}

	// 如果查询id为0 则用户密码不存在 返回403
	if user.Id == 0 {
		return &models.Result{
			Code: 403,
			Data: nil,
			Msg:  "用户名或密码不正确",
		}
	}

	// 如果id有值，根据tokenId获取tokenValue
	token := new(models.Token)
	token.Id = user.TokenId
	err = o.Read(token, "id")
	if err != nil {
		logs.Error("获取token值" + err.Error())
		return result.Error()
	}
	resultData := new(ResultData)
	resultData.TokenValue = token.TokenValue
	resultData.UserTypeId = user.UserTypeId
	return result.Success(resultData)
}
