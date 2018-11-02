package controllers

type BaseAdmin struct {
	Base
}

func (c *BaseAdmin) Prepare() {
	if c.GetSession("admin_auth") != 1 {
		c.Redirect("/", 302)
		c.StopRun()
	}
}
