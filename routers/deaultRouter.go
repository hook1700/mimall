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
}