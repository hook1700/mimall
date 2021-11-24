package routers

import (
	"github.com/astaxie/beego"
	"mimall/controllers/api"
)

func init()  {
	ns :=
		beego.NewNamespace("api",
			beego.NSRouter("login",&api.LoginController{}))
	beego.AddNamespace(ns)
}
