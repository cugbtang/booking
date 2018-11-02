package routers

import (
	"bookingSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BookingController{}, "*:Show")
	beego.Router(`/booking/get`, &controllers.BookingAdminController{}, "*:GetList")
	beego.Router(`/booking/get_list`, &controllers.BookingController{}, "*:GetList")
	//注册 namespace
	beego.Include(&controllers.AdminController{})
	beego.Include(&controllers.ClassAdminController{})
	beego.Include(&controllers.BookingAdminController{})
	beego.Include(&controllers.DesignAdminController{})
	beego.Include(&controllers.BookingController{})
	beego.Include(&controllers.AuthAdminController{})
	beego.Include(&controllers.UserAdminController{})
}
