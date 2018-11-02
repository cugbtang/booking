package controllers

import (
	"bookingSystem/models"
	"github.com/astaxie/beego/orm"
	"strings"
)

type ClassAdminController struct {
	BaseAdmin
}

//展示教室列表
// @router /room/show [get,post]
func (c *ClassAdminController) Show() {
	if c.Ctx.Request.Method == "GET" {
		o := orm.NewOrm()
		var rooms []*models.Room
		o.QueryTable("b_room").All(&rooms)
		c.Data["Room"] = rooms
		c.TplName = "RoomList.html"
	}
}

//添加教室
// @router /room/add [get,post]
func (c *ClassAdminController) Add() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "RoomAdd.html"
	} else {
		o := orm.NewOrm()
		var room models.Room
		room.Name = c.GetString("name")
		room.StartTime = c.GetString("start_time")
		room.EndTime = c.GetString("end_time")
		room.EndTime = c.GetString("end_time")
		room.PosNum, _ = c.GetInt("pos_num")
		room.Desc = c.GetString("desc")
		room.Status = 1
		_, err := o.Insert(&room)
		if err != nil {
			c.AjaxErr("添加失败")
		} else {
			c.AjaxOk("添加成功")
		}

	}
}

//删除教室
// @router /room/delete [post]
func (c *ClassAdminController) Delete() {
	o := orm.NewOrm()
	id, err := c.GetInt("room_id")
	if err == nil {
		if num, err := o.Delete(&models.Room{RoomId: id}); err == nil && num == 1 {
			c.AjaxOk("删除成功")
		}
	}
	c.AjaxErr("删除失败")
}

//修改教室信息
// @router /room/update/?:id [get,post]
func (c *ClassAdminController) Update() {
	if c.Ctx.Request.Method == "GET" {
		o := orm.NewOrm()
		var room models.Room
		id := c.Ctx.Input.Param(":id")
		o.QueryTable("b_room").Filter("room_id", id).One(&room)
		c.Data["RoomInfo"] = room
		c.TplName = "RoomUpdate.html"
	} else {
		o := orm.NewOrm()
		id, _ := c.GetInt("room_id")
		room := models.Room{RoomId: id}
		if o.Read(&room) == nil {
			if temp := c.GetString("name"); temp != "" {
				room.Name = temp
			}
			if temp := c.GetString("start_time"); temp != "" {
				room.StartTime = temp
			}
			if temp := c.GetString("end_time"); temp != "" {
				room.EndTime = temp
			}
			if temp, _ := c.GetInt("pos_num"); temp != 0 {
				room.PosNum = temp
			}
			if temp := c.GetString("desc"); temp != "" {
				room.Desc = temp
			}
			if temp, _ := c.GetInt("status"); temp != 0 {
				room.Status = temp
			}
			if num, err := o.Update(&room); err == nil && num == 1 {
				c.AjaxOk("修改成功")
			}
		}
		c.AjaxErr("修改失败")
	}
}

func seralizeTime(str string) []string {
	return strings.Split(str, ",")
}
