package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["bookingSystem/controllers:AdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/admin/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:AdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Main",
            Router: `/admin/main`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:AuthAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:AuthAdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/auth/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:AuthAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:AuthAdminController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/auth/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"],
        beego.ControllerComments{
            Method: "Single",
            Router: `/booking/detail/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"],
        beego.ControllerComments{
            Method: "Option",
            Router: `/booking/option`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingAdminController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/booking/show`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/booking/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/booking/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/booking/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:BookingController"],
        beego.ControllerComments{
            Method: "Time",
            Router: `/booking/time`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/room/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/room/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/room/show`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:ClassAdminController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/room/update/?:id`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/design/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/design/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/design/show`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:DesignAdminController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/design/update/?:id`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/user/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/user/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/user/show`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["bookingSystem/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/user/update/?:id`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
