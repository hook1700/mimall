package admin

import (
	"github.com/astaxie/beego"
	"mimall/models"
)

type GoodsController struct {
	BaseController
}

func (c *GoodsController) Get() {
	c.TplName = "admin/goods/index.html"
}

func (c *GoodsController) Add() {
	//获取商品分类
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=?", 0).Preload("GoodsCateItem").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	beego.Info(goodsCate)

	//获取颜色信息
	goodsColor := []models.GoodsColor{}
	models.DB.Find(&goodsColor)
	c.Data["goodsColor"] = goodsColor

	//获取商品类型信息
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType


	c.TplName = "admin/goods/add.html"
}

func (c *GoodsController) DoAdd() {
	c.Ctx.WriteString("执行增加")
}
func (c *GoodsController) Edit() {

	c.Ctx.WriteString("修改")
}

func (c *GoodsController) DoEdit() {
	c.Ctx.WriteString("执行修改")
}
func (c *GoodsController) Delete() {

	c.Ctx.WriteString("删除")
}

func (c *GoodsController) DoUpload() {

	savePath, err := c.UploadImg("file")
	if err != nil {
		beego.Error("失败")
		c.Data["json"] = map[string]interface{}{
			"link": "",
		}
		c.ServeJSON()
	} else {
		//返回json数据 {link: 'path/to/image.jpg'}
		c.Data["json"] = map[string]interface{}{
			"link": "/" + savePath,
		}
		c.ServeJSON()
	}
}


//获取商品类型属性
func (c *GoodsController) GetGoodsTypeAttribute() {
	cate_id, err1 := c.GetInt("cate_id")
	GoodsTypeAttribute := []models.GoodsTypeAttribute{}
	err2 := models.DB.Where("cate_id=?", cate_id).Find(&GoodsTypeAttribute).Error
	if err1 != nil || err2 != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "",
			"success": false,
		}
		c.ServeJSON()

	} else {
		c.Data["json"] = map[string]interface{}{
			"result":  GoodsTypeAttribute,
			"success": true,
		}
		c.ServeJSON()
	}

}