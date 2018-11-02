package controllers

import (
	"net"
	"time"
)

type AdminController struct {
	BaseAdmin
}

// @router /admin/index [get]
func (c *AdminController) Index() {
	c.TplName = "index.html"
}

// @router /admin/main [get]
func (c *AdminController) Main() {
	c.Data["date"] = time.Now()
	c.Data["ip"] = getLocalIp()
	c.TplName = "Welcome.html"
}

func getLocalIp() (IpAddr string) {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		IpAddr = "localhost"
		return
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return
			}
		}
	}
	IpAddr = "localhost"
	return
}
