package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"log"
	"myblog/models"
	_ "myblog/routers"
)

func init() {
	// 静态资源路径设置
	beego.SetStaticPath("/", "static/html")
	beego.SetStaticPath("/images", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/upload", "static/upload")
	beego.SetStaticPath("/lib", "static/lib")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/plugins", "static/plugins")
	// datasource init
	db := new(models.Db)
	db.InitDB()

	// set default database
	dataSource := fmt.Sprintf(`%v:%v@tcp(%v)/%v?charset=utf8`, db.User, db.Password, db.Url, db.DbName)
	log.Println(dataSource)
	orm.RegisterDataBase("default", "mysql", dataSource, 30)

	// register model
	orm.RegisterModel(new(models.Article))
	orm.RegisterModel(new(models.ArticleType))
	orm.RegisterModel(new(models.Token))
	orm.RegisterModel(new(models.User))
	// create table
	//orm.RunSyncdb("default", true, true)

	// 初始化日志
	logs.SetLogger("console")

	// 输出行号
	logs.EnableFuncCallDepth(true)

	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

}
func main() {
	beego.Run()
}
