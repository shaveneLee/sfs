package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Point struct {
	Id           int
	Name         string
	Status       int16
	Type         int16
	PreHours     float64 `orm:"digits(10);decimals(1)"`
	PreStars     float64 `orm:"digits(10);decimals(1)"`
	Hours        float64 `orm:"digits(10);decimals(1)"`
	Stars        float64 `orm:"digits(10);decimals(1)"`
	Points       float64 `orm:"digits(10);decimals(1)"`
	StartTime    time.Time
	EndTime      time.Time
	CreateUserId int
	CreateTime   time.Time
	UpdateUserId int
	UpdateTime   time.Time
	Remark       string
}

func init() {
	orm.RegisterModel(new(Point))
}

func (p *Point) GetPoints() []orm.Params {
	var maps []orm.Params

	o := orm.NewOrm()

	o.QueryTable(new(Point)).Filter("status", 1).Values(&maps)

	return maps
}

//任务类型枚举
func (p *Point) GetTypeList() map[string]string {
	TypeList := make(map[string]string)
	TypeList["10"] = "Planning"
	TypeList["20"] = "Unexpect"
	return TypeList
}
