package main

import (
	_ "bookingSystem/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbPassword := beego.AppConfig.String("dbPassword")
	dbName :=beego.AppConfig.String("dbName")
	dsn := dbUser + ":" +dbPassword +"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn)
}
func main() {
	orm.RunCommand()
	orm.Debug = true
	beego.Run()
}

