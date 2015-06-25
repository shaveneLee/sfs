package controllers

import (
	//	"fmt"
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
