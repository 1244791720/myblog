package models

import "github.com/astaxie/beego"

type Db struct {
	User     string
	Password string
	Url      string
	DbName   string
}

func (db *Db) InitDB() {
	db.User = beego.AppConfig.String("mysqluser")
	db.Password = beego.AppConfig.String("mysqlpass")
	db.Url = beego.AppConfig.String("mysqlurls")
	db.DbName = beego.AppConfig.String("mysqldb")
}
