package models

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	"github.com/astaxie/beego/utils"
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
	StartTime    string
	EndTime      string
	CreateUserId int
	CreateTime   string
	UpdateUserId int
	UpdateTime   string
	Remark       string
}

func init() {
	orm.RegisterModel(new(Point))
}

func (p *Point) Test() string {
	o := orm.NewOrm()
	point := Point{}
	point.Id = 1
	point.Name = "axxxa"
	//err := o.Read(&point)

	_, err := o.QueryTable(new(Point)).Filter("Id", point.Id).Update(orm.Params{
		"Id":   point.Id,
		"Name": point.Name,
	})
	if nil != err {
		return "aaaa"
		return err.Error()
	}
	return "ok"
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
		o.QueryTable(new(Point)).Filter("start_time__gte", start_time).Filter("end_time__lte", end_time).Filter("status", 10).Values(&points)

		data["point"] = points
		result[end_time] = data

		// index_keys[w-i] = end_time
		index_keys = append(index_keys, end_time)
	}

	return index_keys, result
}

//type list enum
func (p *Point) GetTypeList() map[string]string {
	TypeList := make(map[string]string)
	TypeList["10"] = "Planning"
	TypeList["20"] = "Unexpect"
	return TypeList
}

//insert or update point
func (m *Point) InsertOrUpdate() error {
	o := orm.NewOrm()
	q := Point{Id: m.Id}

	err := o.Read(&q)

	if nil != err { // not exit, insert
		m.CreateTime = beego.Date(time.Now(), "Y-m-d H:i:s")
		m.Status = 10
		if _, err := o.Insert(m); err != nil {
			return err
		}
	} else { // update
		_, err := o.QueryTable(new(Point)).Filter("Id", m.Id).Update(orm.Params{
			"Id":     m.Id,
			"Name":   m.Name,
			"Type":   m.Type,
			"Hours":  m.Hours,
			"Stars":  m.Stars,
			"Points": m.Points,
			//"StartTime":  m.StartTime,
			//"EndTime":    m.EndTime,
			"UpdateTime": beego.Date(time.Now(), "Y-m-d H:i:s"),
		})
		if nil != err {
			return err
		}
	}
	return nil
}
