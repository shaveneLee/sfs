package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "sfs/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//dbLink := "ocs:" + "root123@tcp(dbdev.gz.cvte.cn:3306)/sfs?charset=utf8"
	//dbLink := "root:" + "root123@tcp(192.168.2.114:3306)/sfs?charset=utf8"
	dbLink := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname") + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dbLink)
	//orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	beego.SetStaticPath("/img", "img")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")
}

func main() {
	beego.Run()
}
