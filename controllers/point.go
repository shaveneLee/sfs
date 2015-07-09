package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/cache"
	//	"reflect"
	"sfs/common"
	"sfs/models"
	"strconv"
)

type PointController struct {
	beego.Controller
}

type Result struct {
	success bool
	message string
	datas   interface{}
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
}

func (this *PointController) GetTypeList() {
	point := models.Point{}
	this.Data["json"] = point.GetTypeList()
	this.ServeJson()
}

func (this *PointController) SavePoints() {
	result := make(map[string]interface{})
	result["sucess"] = false
	result["message"] = ""
	//var post_data map[string]interface{}
	post_data := make(map[string]map[string]interface{})
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &post_data); err != nil {
		result["message"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	//beego.Info(post_data)
	for _, point := range post_data {
		model := models.Point{}
		model.Id = common.GetInt(point["Id"])
		model.Name = common.GetString(point["Name"])
		model.Type = common.GetInt16(point["Type"])
		model.Hours = common.GetFloat64(point["Hours"])
		model.Stars = common.GetFloat64(point["Stars"])
		model.Points = common.GetFloat64(point["Points"])
		model.StartTime = point["StartTime"].(string)
		model.EndTime = point["EndTime"].(string)

		beego.Info(model)
		err := model.InsertOrUpdate()
		if err != nil {
			result["message"] = err.Error()
			continue
		}
	}

	result["success"] = true
	result["datas"] = post_data
	this.Data["json"] = result
	this.ServeJson()
}

//get specify week points by param.
//@ m param
func (this *PointController) GetWeekPointsJson() {
	id := this.Ctx.Input.Param(":id")
	w, _ := strconv.ParseInt(id, 10, 64)
	init_id := int(w)

	model := models.Point{}

	result := make(map[string]interface{})
	index_keys, points := model.GetWeekPoints(init_id)
	result["index_keys"] = index_keys
	result["points"] = points

	this.Data["json"] = result
	this.ServeJson()
}
