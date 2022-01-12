package index

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/go_image"
	"github.com/skip2/go-qrcode"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController)Get()  {

	//实现图片截切 按宽度进行比例缩放，输入输出都是
	filename := "static/upload/a.jpg"
	savepath := "static/upload/a_800.jpg"
	err1 := go_image.ScaleF2F(filename,savepath,800)
	if err1 != nil{
		beego.Error(err1)
	}

	//按宽度和高度进行比例缩放，输入和输出都是文件
	// filename := "static/upload/b.png"
	// savepath := "static/upload/b_400.png"
	// err := ThumbnailF2F(filename, savepath, 400, 400)
	// if err != nil {
	// 	beego.Error(err)
	// }
	err2 := qrcode.WriteFile("https://www.baidu.com",qrcode.Medium,800,"static/upload/qr.png")
	if err2 != nil{
		beego.Error("生成二维码失败")
	}
	c.TplName = "index/index.html"
}