package middleware

import (


	"github.com/astaxie/beego/context"
	"mimall/models"
)

func DefaultAuth(ctx *context.Context) {
	//判断前端用户有没有登录
	user := models.User{}
	models.Cookie.Get(ctx, "userinfo", &user)
	if len(user.Phone) != 11 {
		ctx.Redirect(302, "/pass/login")
	}
}
