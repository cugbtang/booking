package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	UserId     int       `orm:"pk;auto"`
	UserName   string    `orm:"unique"`
	NickName   string    `orm:"size(100)"`
	Password   string    `orm:"size(100)"`
	Salt       string    `orm:"size(10)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("b_", new(User))
}
