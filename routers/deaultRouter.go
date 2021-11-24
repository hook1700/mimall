package routers

import (
	"github.com/astaxie/beego"
	"mimall/controllers/index"
)

func init()  {
	beego.Router("/",&index.IndexController{})
	beego.Router("/login",&index.LoginController{})
}