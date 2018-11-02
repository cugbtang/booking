package controllers

import (
	"bookingSystem/models"
	"github.com/astaxie/beego/orm"
)

type BookingAdminController struct {
	BaseAdmin
}

//展示预约列表
// @router /booking/show [get]
func (c *BookingAdminController) Show() {
	c.TplName = "BookingShow.html"
}

//展示单个预约
// @router /booking/detail/:id [get]
func (c *BookingAdminController) Single() {
	o := orm.NewOrm()
	var booking models.Booking
	id := c.Ctx.Input.Param(":id")
	err := o.QueryTable("b_booking").RelatedSel().Filter("booking_id", id).One(&booking)
	if err == nil {
		timeStart := seralizeTime(booking.Room.StartTime)[booking.BookingTimeType]
		timeEnd := seralizeTime(booking.Room.EndTime)[booking.BookingTimeType]
		booking.Date = booking.BookingDate.Format("2006-01-02") + " " + timeStart + "~" + timeEnd
		booking.RoomName = booking.Room.Name
		c.Data["data"] = booking
		c.TplName = "BookingDetail.html"
	}
}

//获取预约列表
func (c *BookingAdminController) GetList() {
	var page, limit int
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&limit, "limit")
	o := orm.NewOrm()
	var bookings []*models.Booking
	o.QueryTable("b_booking").RelatedSel().OrderBy("booking_status").Limit(limit, (page-1)*10).All(&bookings)
	for _, v := range bookings {
		timeStart := seralizeTime(v.Room.StartTime)[v.BookingTimeType]
		timeEnd := seralizeTime(v.Room.EndTime)[v.BookingTimeType]
		v.Date = v.BookingDate.Format("2006-01-02") + " " + timeStart + "~" + timeEnd
		v.RoomName = v.Room.Name
	}
	json := make(map[string]interface{})
	json["data"] = bookings
	json["code"] = 0
	json["msg"] = "获取成功"
	cnt, _ := o.QueryTable("b_booking").Count()
	json["count"] = cnt
	c.Data["json"] = json
	c.ServeJSON()
}

//（通过|不通过）预约
// @router /booking/option [post]
func (c *BookingAdminController) Option() {
	o := orm.NewOrm()
	id, _ := c.GetInt("id")
	status, _ := c.GetInt("status")
	booking := models.Booking{BookingId: id}
	if o.Read(&booking) == nil {
		booking.BookingStatus = status
		if num, err := o.Update(&booking); err == nil && num == 1 {
			c.AjaxOk("修改成功")
		}
	}
	c.AjaxErr("修改失败")
}
