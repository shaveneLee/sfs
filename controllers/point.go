package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"sfs/models"
)

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

	index_keys, points := model.GetWeekPoints(2)
	for k, v := range index_keys {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(points)
	this.Data["Str1"] = index_keys

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

//get specify week points by param.
//@ m param
func (this *PointController) GetWeekPoints(w int) interface{} {
	//point := models.Point{}
	m := make(map[string]interface{})
	m["aa"] = "bbb"
	x := make(map[string]string)
	x["xxx"] = "yy"
	m["Point"] = x
	return m
}
