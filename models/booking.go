package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Booking struct {
	BookingId int `orm:"pk;auto"`
	BookingUserName string `orm:"size(100)"`
	BookingUserTel string `orm:"size(100)"`
	BookingContent string
	BookingUseNum int `orm:"size(3)"`
	BookingStatus int `orm:"size(3)"`
	BookingOrg string `orm:"size(100)"`
	BookingTimeType int `orm:"size(3)"`
	BookingExt string
	BookingDate time.Time `orm:"auto_now;type(date)"`
	Date string `orm:"-"`
	RoomName string `orm:"-"`
	Room     *Room   `orm:"rel(fk)"`
}
func init(){
	orm.RegisterModelWithPrefix("b_", new(Booking))
}

