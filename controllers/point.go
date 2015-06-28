package controllers

import (
	//	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"sfs/models"
	"time"
)

//time format layout
const t_layout = "2006-01-02"

type PointController struct {
	beego.Controller
}

func (this *PointController) Index() {
	this.Data["Str1"] = "aaa"
	this.Data["json"] = "aaadd"
	this.ServeJson()
}

func (this *PointController) Edit() {
	this.TplNames = "point/edit.tpl"
	beego.TemplateLeft = "{{{"
	beego.TemplateRight = "}}}"
	model := models.Point{}
	this.Data["Str1"] = model.GetPoints()

	this.Data["Str1"] = this.GetSunday()

}

func (this *PointController) GetMonday() string {
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

func (this *PointController) GetSunday() string {
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

func (this *PointController) GetTypeList() {
	point := models.Point{}
	this.Data["json"] = point.GetTypeList()
	this.ServeJson()
}

func (this *PointController) SavePoints() {
	result := make(map[string]interface{})
	result["sucess"] = true
	result["message"] = "call success"
	result["datas"] = "xxx"
	json_str, _ := json.Marshal(result)
	this.Ctx.Output.Body([]byte(json_str))
	return
}
