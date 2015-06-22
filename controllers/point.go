package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"sfs/models"
)

type PointController struct {
	beego.Controller
}

func (c *PointController) Test() {
	c.TplNames = "points/index.tpl"

	u := models.Points{}
	maps := u.GetPoints()

	for _, m := range maps {
		//fmt.Println(m["Id"], m["Name"])
		fmt.Println(m)
		c.Data["Id"] = m["Id"]
		c.Data["Name"] = m["Name"]
	}
}

func (c *PointController) Index() {
	c.TplNames = "points/index.tpl"
	beego.TemplateLeft = "{{{"
	beego.TemplateRight = "}}}"
	c.Data["Str1"] = "aaa"
}

func (c *PointController) Edit() {
	c.TplNames = "points/edit.tpl"
	beego.TemplateLeft = "{{{"
	beego.TemplateRight = "}}}"
	c.Data["Str1"] = "aaa"
}
