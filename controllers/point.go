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
	this.Data["Str1"] = model.GetPoints()

	//index_keys, points := model.GetWeekPoints(1)
	//this.Data["Str1"] = index_keys

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
	// p := make([]byte, 10)
	// bytesRead, _ := this.Ctx.Request.Body

	// result["datas"] = this.Ctx.Request.Body
	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	result["datas"] = req.Title

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
