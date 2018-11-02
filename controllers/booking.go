package controllers

import (
	"bookingSystem/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type BookingController struct {
	Base
}

//新预约
// @router /booking/new [get]
func (c *BookingController) Show() {
	o := orm.NewOrm()
	var rooms []*models.Room
	o.QueryTable("b_room").Filter("status", 1).All(&rooms)
	c.Data["RoomList"] = rooms
	var designs []*models.Design
	o.QueryTable("b_design").Filter("status", 1).All(&designs)
	c.Data["DesignList"] = designs
	c.TplName = "NewBooking.html"
}

//预约列表
// @router /booking/list [get]
func (c *BookingController) List() {
	c.TplName = "ShowBooking.html"
}

//获取可以预约的时间段
// @router /booking/time [post]
func (c *BookingController) Time() {
	o := orm.NewOrm()
	id, _ := c.GetInt("room_id")
	date := c.GetString("date")
	room := models.Room{RoomId: id}
	o.Read(&room)
	if room.Status != 1 {
		c.AjaxErr("信息错误")
	}
	var list orm.ParamsList
	o.QueryTable("b_booking").Filter("room_id", id).Filter("booking_date", date).ValuesFlat(&list, "booking_time_type")
	startList := seralizeTime(room.StartTime)
	endList := seralizeTime(room.EndTime)
	res := make([]interface{}, 0)
	for k, v := range startList {
		if !inArr(k, list) {
			single := make(map[string]interface{})
			single["type"] = k
			single["str"] = v + "~" + endList[k]
			res = append(res, single)
		}
	}
	json := make(map[string]interface{})
	json["status"] = 1
	json["msg"] = "获取成功"
	json["data"] = res
	c.Data["json"] = json
	c.ServeJSON()
}
func inArr(str int, arr []interface{}) bool {
	for _, v := range arr {
		if v == int64(str) {
			return true
		}
	}
	return false
}

//增加
// @router /booking/add [post]
func (c *BookingController) Add() {
	//1、验证借用的教室是否存在，是否处于可借用状态，存在则取出时间段
	o := orm.NewOrm()
	var newBooking models.Booking
	var room models.Room
	id, _ := c.GetInt("room_id")
	rooms := models.Room{RoomId: id}
	o.Read(&rooms)
	newBooking.Room = &rooms
	err := o.QueryTable("b_room").Filter("room_id", id).Filter("status", 1).One(&room)
	if err != nil || room.RoomId == 0 {
		c.AjaxErr("该教室不存在")
	}
	//2、验证时间段是否正确
	timeArr := seralizeTime(room.StartTime)
	useType, err := c.GetInt("use_type")
	newBooking.BookingTimeType = useType
	if err != nil || useType < 0 || useType > len(timeArr)-1 {
		c.AjaxErr("借用时间段错误")
	}
	//3、该时间段是否被借用
	valid := validation.Validation{}
	var booking models.Booking
	useDate := c.GetString("use_date")
	date, _ := time.Parse("2006-01-02", useDate)
	newBooking.BookingDate = date
	//初始化其他变量
	newBooking.BookingContent = c.GetString("content")
	newBooking.BookingOrg = c.GetString("org")
	newBooking.BookingUseNum, _ = c.GetInt("num")
	newBooking.BookingUserName = c.GetString("name")
	newBooking.BookingUserTel = c.GetString("tel")
	valid.Required(newBooking.BookingDate, "type")
	valid.Required(newBooking.BookingContent, "content")
	valid.Required(newBooking.BookingUseNum, "num")
	valid.Required(newBooking.BookingUserName, "name")
	valid.Required(newBooking.BookingUserTel, "tel")
	valid.Required(newBooking.BookingOrg, "org")
	if valid.HasErrors() {
		c.AjaxErr("信息有误，请重新填写！")
	}
	err = o.QueryTable("b_booking").Filter("booking_date", useDate).Filter("booking_time_type", useType).One(&booking)
	if (err != nil) == false {
		c.AjaxErr("该时段已经被借用！")
	}
	//5、格式化添加扩展信息
	var design []*models.Design
	o.QueryTable("b_design").Filter("status", 1).All(&design)
	var ext string
	for _, v := range design {
		ext += v.Name + ":" + c.GetString(v.Key) + "\n"
	}
	newBooking.BookingExt = ext
	//6、insert
	_, err = o.Insert(&newBooking)
	if err != nil {
		c.AjaxErr("借用失败！")
	}
	//7、返回结果
	c.AjaxOk("借用成功,请等待审核")
}

//获取预约列表
func (c *BookingController) GetList() {
	var page, limit int
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&limit, "rows")
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
	json["rows"] = bookings
	json["code"] = 0
	json["msg"] = "获取成功"
	cnt, _ := o.QueryTable("b_booking").Count() // SELECT COUNT(*) FROM USER
	json["total"] = cnt
	c.Data["json"] = json
	c.ServeJSON()
}
