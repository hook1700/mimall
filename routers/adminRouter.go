package routers

import (
	"mimall/controllers/admin"
	"mimall/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns :=
		beego.NewNamespace("/admin",
			//中间件:匹配路由前会执,可以用于权限验证
			//注意引入的包： github.com/astaxie/beego/context
			beego.NSBefore(func(ctx *context.Context) {
				// userinfo := c.GetSession("userinfo")
				pathname := ctx.Request.URL.String()
				userinfo, ok := ctx.Input.Session("userinfo").(models.Manager) //类型断言
				if !(ok && userinfo.Username != "") {
					if pathname != "/admin/login" && pathname != "/admin/login/doLogin" {
						ctx.Redirect(302, "/admin/login")
					}
				}

			}),
			beego.NSRouter("/", &admin.MainController{}),
			beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),

			beego.NSRouter("/manager", &admin.ManagerController{}),
			beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
			beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),

			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),

			beego.NSRouter("/focus", &admin.FocusController{}),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
