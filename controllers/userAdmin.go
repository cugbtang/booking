package controllers

import (
	"bookingSystem/models"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
)

type UserAdminController struct {
	BaseAdmin
}

// @router /user/show [get]
func (c *UserAdminController) Show() {
	o := orm.NewOrm()
	var user []*models.User
	o.QueryTable("b_user").All(&user)
	c.Data["Design"] = user
	c.TplName = "UserList.html"
}

// @router /user/add [get,post]
func (c *UserAdminController) Add() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "UserAdd.html"
	} else {
		o := orm.NewOrm()
		var user models.User
		user.UserName = c.GetString("username")
		rand.Seed(time.Now().Unix())
		user.Salt = GetRandomString(5)
		user.Password = str2md5(c.GetString("password") + user.Salt)
		user.NickName = c.GetString("nickname")
		_, err := o.Insert(&user)
		if err != nil {
			c.AjaxOk("添加失败")
		} else {
			c.AjaxOk("添加成功")
		}
	}
}
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// @router /user/update/?:id [get,post]
func (c *UserAdminController) Update() {
	if c.Ctx.Request.Method == "GET" {
		o := orm.NewOrm()
		var user models.User
		id := c.Ctx.Input.Param(":id")
		o.QueryTable("b_user").Filter("user_id", id).One(&user)
		c.Data["UserInfo"] = user
		c.TplName = "UserUpdate.html"
	} else {
		o := orm.NewOrm()
		id, _ := c.GetInt("user_id")
		user := models.User{UserId: id}
		if o.Read(&user) == nil {
			if temp := c.GetString("nickname"); temp != "" {
				user.NickName = temp
			}
			if temp := c.GetString("password1"); temp != "" {
				user.Salt = GetRandomString(5)
				user.Password = str2md5(temp + user.Salt)
			}
			if num, err := o.Update(&user); err == nil && num == 1 {
				c.AjaxOk("修改成功")
			}
		}
		c.AjaxErr("修改失败")
	}
}

// @router /user/delete [post]
func (c *UserAdminController) Delete() {
	o := orm.NewOrm()
	id, err := c.GetInt("user_id")
	if err == nil {
		if num, err := o.Delete(&models.User{UserId: id}); err == nil && num == 1 {
			c.AjaxOk("删除成功")
		}
	}
	c.AjaxErr("删除失败")
}
