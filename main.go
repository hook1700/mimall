package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"mimall/models"
	_ "mimall/routers"
)

func init()  {
	gob.Register(models.Manager{})
}

func main() {
	//注册模板函数
	beego.AddFuncMap("unixToDate",models.UnixToDate)
	beego.AddFuncMap("setting", models.GetSettingFromColumn)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("formatAttr", models.FormatAttr)
	beego.AddFuncMap("mul", models.Mul)


	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
}

