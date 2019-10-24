package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"myblog/models"
)

func UserSubmit(username, password string) error {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	err = o.Begin()
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	// 生成token放入token表
	token := new(models.Token)
	token.TokenValue = randToken(10)
	tokenId, err := o.Insert(token)
	if err != nil {
		o.Rollback()
		logs.Error(err.Error())
		return err
	}
	user := new(models.User)
	user.UserTypeId = 1
	user.UserName = username
	user.UserPwd = password
	user.TokenId = tokenId
	_, err = o.Insert(user)
	if err != nil {
		o.Rollback()
		logs.Error(err.Error())
		return err
	}
	o.Commit()
	return nil
}

// 生成num*2位的字符串
func randToken(num int) string {
	b := make([]byte, num)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func HasUser(username string) (bool, error) {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		logs.Error(err.Error())
		return true, err
	}
	user := new(models.User)
	user.UserName = username
	err = o.Read(user, "user_name")
	if err != nil {
		logs.Error("根据用户名密码查询User:" + err.Error())
		return false, nil
	}

	// 如果查询id为0 则用户密码不存在 返回403
	if user.Id == 0 {
		return false, nil
	}
	return true, nil
}
