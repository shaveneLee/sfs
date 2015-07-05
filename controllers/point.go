package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
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
	model := models.Point{}
	this.Data["Str1"] = model.Test()
}

func (this *PointController) GetTypeList() {
	point := models.Point{}
	this.Data["json"] = point.GetTypeList()
	this.ServeJson()
}

func (this *PointController) SavePoints() {
	result := make(map[string]interface{})
	result["sucess"] = true
	result["message"] = ""
	//var post_data interface{}
	var post_data map[string]interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &post_data); err != nil {
		result["message"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	model := models.Point{}
	model.Id = 1
	model.Name = "CCCC"
	err := model.InsertOrUpdate()
	if err != nil {

	}
	/*
		for _, point := range post_data {
			for _, value := range point{
				beego.Info(point)
			}
			beego.Info(point)
		}
		/*
			err := model.InsertOrUpdate()
			if (err != nil) {
				result["message"] = err.Error()
				this.Data["json"] = result
				this.ServeJson()
				return;
			}
	*/

	result["success"] = true
	result["datas"] = err
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
