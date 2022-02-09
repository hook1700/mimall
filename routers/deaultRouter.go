package routers

import (
	"github.com/astaxie/beego"
	"mimall/controllers/itying"
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

}