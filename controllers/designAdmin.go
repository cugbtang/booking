package controllers

import (
	"bookingSystem/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
)

type DesignAdminController struct {
	BaseAdmin
}

//展示设计列表
// @router /design/show [get]
func (c *DesignAdminController) Show() {
	o := orm.NewOrm()
	var design []*models.Design
	o.QueryTable("b_design").All(&design)
	c.Data["Design"] = design
	c.TplName = "DesignList.html"
}

//添加设计
// @router /design/add [get,post]
func (c *DesignAdminController) Add() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "DesignAdd.html"
	} else {
		o := orm.NewOrm()
		var design models.Design
		design.Name = c.GetString("name")
		design.Desc = c.GetString("desc")
		design.Status = 1
		design.Key = str2md5(design.Name)
		_, err := o.Insert(&design)
		if err != nil {
			c.AjaxErr("添加失败")
		} else {
			c.AjaxOk("添加成功")
		}
	}
}
func str2md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)               //将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
}

//删除设计
// @router /design/delete [post]
func (c *DesignAdminController) Delete() {
	o := orm.NewOrm()
	id, err := c.GetInt("design_id")
	if err == nil {
		if num, err := o.Delete(&models.Design{DesignId: id}); err == nil && num == 1 {
			c.AjaxOk("删除成功")
		}
	}
	c.AjaxErr("删除失败")
}

//修改设计
// @router /design/update/?:id [get,post]
func (c *DesignAdminController) Update() {
	if c.Ctx.Request.Method == "GET" {
		o := orm.NewOrm()
		var design models.Design
		id := c.Ctx.Input.Param(":id")
		o.QueryTable("b_design").Filter("design_id", id).One(&design)
		c.Data["DesignInfo"] = design
		c.TplName = "DesignUpdate.html"
	} else {
		o := orm.NewOrm()
		id, _ := c.GetInt("design_id")
		design := models.Design{DesignId: id}
		if o.Read(&design) == nil {
			if temp := c.GetString("name"); temp != "" {
				design.Name = temp
				design.Key = str2md5(temp)
			}
			if temp := c.GetString("desc"); temp != "" {
				design.Desc = temp
			}
			if temp, _ := c.GetInt("status"); temp != 0 {
				design.Status = temp
			}
			if num, err := o.Update(&design); err == nil && num == 1 {
				c.AjaxOk("修改成功")
			}
		}
		c.AjaxErr("修改失败")
	}
}
