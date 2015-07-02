package models

import (
	// "fmt"
	"github.com/astaxie/beego"
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

func (this *Point) GetMonday(t time.Time) string {
	//t := time.Now()
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

func (this *Point) GetSunday(t time.Time) string {
	//t := time.Now()
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
	//monday = t.Format(t_layout)
	monday = beego.Date(t, "Y-m-d")
	return monday
}

//get specify week points.
//@ w param specify how many weeks data to get (not include this week)
func (this *Point) GetWeekPoints(w int) ([]string, interface{}) {
	var points []orm.Params
	o := orm.NewOrm()
	index_keys := []string{}
	result := make(map[string]interface{})

	t := time.Now()
	sunday_str := this.GetSunday(t)
	t, _ = beego.DateParse(sunday_str, "Y-m-d")
	for i := w; i >= 0; i-- {
		last_sunday := t.AddDate(0, 0, -7*i)
		last_monday := last_sunday.AddDate(0, 0, -6)
		last_sunday = last_sunday.Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))
		end_time := beego.Date(last_sunday, "Y-m-d H:i:s")
		end_date := beego.Date(last_sunday, "Y-m-d")
		start_time := beego.Date(last_monday, "Y-m-d")

		data := make(map[string]interface{})
		data["start_time"] = start_time
		data["end_date"] = end_date
		data["end_time"] = end_time
		data["time_str"] = start_time + "_" + end_time
		o.QueryTable(new(Point)).Filter("create_time__gte", start_time).Filter("create_time__lte", end_time).Values(&points)

		data["point"] = points
		result[end_time] = data

		// index_keys[w-i] = end_time
		index_keys = append(index_keys, end_time)
	}

	return index_keys, result
}

//任务类型枚举
func (p *Point) GetTypeList() map[string]string {
	TypeList := make(map[string]string)
	TypeList["10"] = "Planning"
	TypeList["20"] = "Unexpect"
	return TypeList
}
