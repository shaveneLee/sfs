package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//time format layout
const t_layout = "2006-01-02"

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

	o.QueryTable(new(Point)).Filter("create_time__gte", "2015-06-29").Filter("create_time__lte", "2015-07-05").Values(&maps)

	return maps
}

func (this *Point) GetMonday() string {
	t := time.Now()
	week := t.Weekday().String()
	var monday string
	switch week {
	case "Sunday":
		t = t.AddDate(0, 0, -6)
	case "Monday":
	case "Tuesday":
		t = t.AddDate(0, 0, -1)
	case "Wednesday":
		t = t.AddDate(0, 0, -2)
	case "Thursday":
		t = t.AddDate(0, 0, -3)
	case "Friday":
		t = t.AddDate(0, 0, -4)
	case "Saturday":
		t = t.AddDate(0, 0, -5)
	}
	monday = t.Format(t_layout)
	return monday
}

func (this *Point) GetSunday() string {
	t := time.Now()
	week := t.Weekday().String()
	var monday string
	switch week {
	case "Sunday":
	case "Monday":
		t = t.AddDate(0, 0, 6)
	case "Tuesday":
		t = t.AddDate(0, 0, 5)
	case "Wednesday":
		t = t.AddDate(0, 0, 4)
	case "Thursday":
		t = t.AddDate(0, 0, 3)
	case "Friday":
		t = t.AddDate(0, 0, 2)
	case "Saturday":
		t = t.AddDate(0, 0, 1)
	}
	monday = t.Format(t_layout)
	return monday
}

//get specify week points.
//@ w param specify how many weeks data to get
func (this *Point) GetWeekPoints(w int) interface{} {
	//point := models.Point{}
	m := make(map[string]interface{})
	m["aa"] = "bbb"
	x := make(map[string]string)
	x["xxx"] = "yy"
	m["Point"] = x
	return m
}

//任务类型枚举
func (p *Point) GetTypeList() map[string]string {
	TypeList := make(map[string]string)
	TypeList["10"] = "Planning"
	TypeList["20"] = "Unexpect"
	return TypeList
}
