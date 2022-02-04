package itying

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"mimall/models"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) SuperInit() {
	//获取顶部导航
	topNav := []models.Nav{}
	if hasTopNav := models.CacheDb.Get("topNav", &topNav); hasTopNav == true {
		c.Data["topNavList"] = topNav
	} else {
		models.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topNav)
		c.Data["topNavList"] = topNav
		models.CacheDb.Set("topNav", topNav)
	}
	//左侧分类  https://gorm.io/zh_CN/docs/preload.html
	goodsCate := []models.GoodsCate{}

	if hasGoodsCate := models.CacheDb.Get("goodsCate", &goodsCate); hasGoodsCate == true {
		c.Data["goodsCateList"] = goodsCate
	} else {
		models.DB.Preload("GoodsCateItem", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Where("pid=0 AND status=1").Order("sort desc", true).Find(&goodsCate)
		c.Data["goodsCateList"] = goodsCate
		models.CacheDb.Set("goodsCate", goodsCate)
	}

	//获取中间导航的数据
	middleNav := []models.Nav{}
	if hasMiddleNav := models.CacheDb.Get("middleNav", &middleNav); hasMiddleNav == true {
		c.Data["middleNavList"] = middleNav
	} else {
		models.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleNav)

		for i := 0; i < len(middleNav); i++ {
			//获取关联商品
			// middleNav[i].Relation  19,20,21
			middleNav[i].Relation = strings.ReplaceAll(middleNav[i].Relation, "，", ",")
			relation := strings.Split(middleNav[i].Relation, ",")
			goods := []models.Goods{}
			models.DB.Where("id in (?)", relation).Select("id,title,goods_img,price").Find(&goods)
			middleNav[i].GoodsItem = goods
		}
		c.Data["middleNavList"] = middleNav
		models.CacheDb.Set("middleNav", middleNav)
	}
}
