package routers

import (
	"github.com/astaxie/beego"
	"mimall/controllers/itying"
	"mimall/middleware"
)

func init()  {
	beego.Router("/", &itying.IndexController{})
	beego.Router("/category_:id([0-9]+).html", &itying.ProductController{}, "get:CategoryList")
	beego.Router("/item_:id([0-9]+).html", &itying.ProductController{}, "get:ProductItem")
	beego.Router("/product/getImgList", &itying.ProductController{}, "get:GetImgList")
	beego.Router("/user", &itying.UserController{})

	beego.Router("/cart", &itying.CartController{})
	beego.Router("/cart/addCart", &itying.CartController{}, "get:AddCart")
	beego.Router("/cart/decCart", &itying.CartController{}, "get:DecCart")
	beego.Router("/cart/incCart", &itying.CartController{}, "get:IncCart")
	beego.Router("/cart/changeOneCart", &itying.CartController{}, "get:ChangeOneCart")
	beego.Router("/cart/changeAllCart", &itying.CartController{}, "get:ChangeAllCart")
	beego.Router("/cart/delCart", &itying.CartController{}, "get:DelCart")

	beego.Router("/pass/login", &itying.PassController{}, "get:Login")
	beego.Router("/pass/doLogin", &itying.PassController{}, "post:DoLogin")
	beego.Router("/pass/loginOut", &itying.PassController{}, "get:LoginOut")

	beego.Router("/pass/registerStep1", &itying.PassController{}, "get:RegisterStep1")
	beego.Router("/pass/registerStep2", &itying.PassController{}, "get:RegisterStep2")
	beego.Router("/pass/registerStep3", &itying.PassController{}, "get:RegisterStep3")

	beego.Router("/pass/sendCode", &itying.PassController{}, "get:SendCode")
	beego.Router("/pass/validateSmsCode", &itying.PassController{}, "get:ValidateSmsCode")
	beego.Router("/pass/doRegister", &itying.PassController{}, "post:DoRegister")

	//配置中间件判断权限
	//配置中间件判断权限
	beego.InsertFilter("/buy/*", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/buy/checkout", &itying.BuyController{}, "get:Checkout")
	beego.Router("/buy/doOrder", &itying.BuyController{}, "post:DoOrder")
	beego.Router("/buy/confirm", &itying.BuyController{}, "get:Confirm")

	//配置中间件判断权限
	beego.InsertFilter("/address/*", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/address/addAddress", &itying.AddressController{}, "post:AddAddress")
	beego.Router("/address/getOneAddressList", &itying.AddressController{}, "get:GetOneAddressList")

	beego.Router("/address/doEditAddressList", &itying.AddressController{}, "post:DoEditAddress")

	beego.Router("/address/changeDefaultAddress", &itying.AddressController{}, "get:ChangeDefaultAddress")

	beego.Router("/user", &itying.UserController{})

}