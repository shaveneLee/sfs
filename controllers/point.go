package controllers

import (
	//	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"sfs/models"
)

type PointController struct {
	beego.Controller
}

func (c *PointController) Index() {
	c.Data["Str1"] = "aaa"
	c.Data["json"] = "aaadd" // map[string]string{"ObjectId": "hello"}
	c.ServeJson()
}

func (c *PointController) Edit() {
	c.TplNames = "points/edit.tpl"
	beego.TemplateLeft = "{{{"
	beego.TemplateRight = "}}}"
	model := models.Points{}
	c.Data["Str1"] = model.GetPoints()
	c.Data["TypeList"] = model.GetTypeList()
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
