package dao

import "github.com/astaxie/beego/orm"

// 查询token是否存在
func IsRightToken(token string) bool{
	o := orm.NewOrm()
	o.Using("default")
	return o.QueryTable("token").Filter("token_value", token).Exist()
}
