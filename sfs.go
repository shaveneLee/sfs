package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "sfs/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//dbLink := "ocs:" + "root123@tcp(dbdev.gz.cvte.cn:3306)/sfs?charset=utf8"
	dbLink := "root:" + "root123@tcp(192.168.2.114:3306)/sfs?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dbLink)
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	beego.SetStaticPath("/img", "img")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")
}

func main() {
	beego.Run()
}
