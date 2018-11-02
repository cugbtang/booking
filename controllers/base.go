package controllers

import (
	"github.com/astaxie/beego"
)

type Base struct {
	beego.Controller
}

func (c *Base) AjaxOk(str string) {
	c.Data["json"] = ajax(str, 1)
	c.ServeJSON()
	c.StopRun()
}
func (c *Base) AjaxErr(str string) {
	c.Data["json"] = ajax(str, 1)
	c.ServeJSON()
	c.StopRun()
}

func ajax(str string, status int) (map[string]interface{}) {
	json := make(map[string]interface{})
	json["status"] = status
	json["msg"] = str
	return json
}
