package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Room struct {
	RoomId     int `orm:"pk;auto"`
	Name       string
	PosNum     int `orm:"size(10)"`
	Desc       string
	StartTime  string
	EndTime    string
	Status     int        `orm:"default(1)"`
	UpdateTime time.Time  `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("b_", new(Room))
}
