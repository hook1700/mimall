package routers

import (
	"github.com/astaxie/beego"
	"mimall/controllers/admin"
	"mimall/middleware"
)

func init() {
	ns :=
		beego.NewNamespace("/"+ beego.AppConfig.String("adminPath"),
			//中间件:匹配路由前会执,可以用于权限验证
			//注意引入的包： github.com/astaxie/beego/context
			beego.NSBefore(middleware.AdminAuth),
			beego.NSRouter("/", &admin.MainController{}),
			beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),
			beego.NSRouter("/main/changeStatus", &admin.MainController{}, "get:ChangeStatus"),
			beego.NSRouter("/main/editNum", &admin.MainController{}, "get:EditNum"),


			//登录
			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),
			beego.NSRouter("/login/loginOut", &admin.LoginController{}, "get:LoginOut"),

			beego.NSRouter("/focus", &admin.FocusController{}),
			//角色管理
			beego.NSRouter("/role", &admin.RoleController{}),
			beego.NSRouter("/role/add", &admin.RoleController{}, `get:Add`),
			beego.NSRouter("/role/edit", &admin.RoleController{}, `get:Edit`),
			beego.NSRouter("/role/doAdd", &admin.RoleController{}, `post:DoAdd`),
			beego.NSRouter("/role/doEdit", &admin.RoleController{}, `post:DoEdit`),
			beego.NSRouter("/role/delete", &admin.RoleController{}, `get:Delete`),
			beego.NSRouter("/role/auth", &admin.RoleController{}, `get:Auth`),
			beego.NSRouter("/role/doAuth", &admin.RoleController{}, `post:DoAuth`),


			//管理员管理
			beego.NSRouter("/manager", &admin.ManagerController{}),
			beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
			beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
			beego.NSRouter("/manager/doAdd", &admin.ManagerController{}, `post:DoAdd`),
			beego.NSRouter("/manager/doEdit", &admin.ManagerController{}, `post:DoEdit`),
			beego.NSRouter("/manager/delete", &admin.ManagerController{}, `get:Delete`),

			//权限管理
			beego.NSRouter("/access", &admin.AccessController{}),
			beego.NSRouter("/access/add", &admin.AccessController{}, "get:Add"),
			beego.NSRouter("/access/edit", &admin.AccessController{}, "get:Edit"),
			beego.NSRouter("/access/doAdd", &admin.AccessController{}, `post:DoAdd`),
			beego.NSRouter("/access/doEdit", &admin.AccessController{}, `post:DoEdit`),
			beego.NSRouter("/access/delete", &admin.AccessController{}, `get:Delete`),

			//轮播图管理
			beego.NSRouter("/focus", &admin.FocusController{}),
			beego.NSRouter("/focus/add", &admin.FocusController{}, `get:Add`),
			beego.NSRouter("/focus/edit", &admin.FocusController{}, `get:Edit`),
			beego.NSRouter("/focus/doAdd", &admin.FocusController{}, `post:DoAdd`),
			beego.NSRouter("/focus/doEdit", &admin.FocusController{}, `post:DoEdit`),
			beego.NSRouter("/focus/delete", &admin.FocusController{}, `get:Delete`),

			//商品分类管理
			beego.NSRouter("/goodsCate", &admin.GoodsCateController{}),
			beego.NSRouter("/goodsCate/add", &admin.GoodsCateController{}, `get:Add`),
			beego.NSRouter("/goodsCate/edit", &admin.GoodsCateController{}, `get:Edit`),
			beego.NSRouter("/goodsCate/doAdd", &admin.GoodsCateController{}, `post:DoAdd`),
			beego.NSRouter("/goodsCate/doEdit", &admin.GoodsCateController{}, `post:DoEdit`),
			beego.NSRouter("/goodsCate/delete", &admin.GoodsCateController{}, `get:Delete`),

			//商品类型管理
			beego.NSRouter("/goodsType", &admin.GoodsTypeController{}),
			beego.NSRouter("/goodsType/add", &admin.GoodsTypeController{}, `get:Add`),
			beego.NSRouter("/goodsType/edit", &admin.GoodsTypeController{}, `get:Edit`),
			beego.NSRouter("/goodsType/doAdd", &admin.GoodsTypeController{}, `post:DoAdd`),
			beego.NSRouter("/goodsType/doEdit", &admin.GoodsTypeController{}, `post:DoEdit`),
			beego.NSRouter("/goodsType/delete", &admin.GoodsTypeController{}, `get:Delete`),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
