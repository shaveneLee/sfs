package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Points struct {
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
	orm.RegisterModel(new(Points))
}

func (u *Points) GetPoints() []orm.Params {
	var maps []orm.Params

	o := orm.NewOrm()

	o.QueryTable(new(Points)).Limit(10).Values(&maps)

	return maps
}
