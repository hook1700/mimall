package itying

import (
	"math"
	"mimall/models"
	"strconv"
	"strings"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) CategoryList() {

	//调用公共功能
	c.SuperInit()

	id := c.Ctx.Input.Param(":id")
	cateId, _ := strconv.Atoi(id)
	curretGoodsCate := models.GoodsCate{}
	subGoodsCate := []models.GoodsCate{}
	models.DB.Where("id=?", cateId).Find(&curretGoodsCate)
	//分页
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5

	var tempSlice []int
	if curretGoodsCate.Pid == 0 { //顶级分类
		//二级分类
		models.DB.Where("pid=?", curretGoodsCate.Id).Find(&subGoodsCate)
		for i := 0; i < len(subGoodsCate); i++ {
			tempSlice = append(tempSlice, subGoodsCate[i].Id)
		}
	} else {
		//获取当前二级分类对应的兄弟分类
		models.DB.Where("pid=?", curretGoodsCate.Pid).Find(&subGoodsCate)
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	goods := []models.Goods{}
	models.DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Offset((page - 1) * pageSize).Limit(pageSize).Order("sort desc").Find(&goods)
	//查询goods表里面的数量
	var count int
	models.DB.Where(where, tempSlice).Table("goods").Count(&count)

	c.Data["goodsList"] = goods
	c.Data["subGoodsCate"] = subGoodsCate
	c.Data["curretGoodsCate"] = curretGoodsCate
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page

	//指定分类模板
	tpl := curretGoodsCate.Template
	if tpl == "" {
		tpl = "itying/product/list.html"
	}

	c.TplName = tpl
}

func (c *ProductController) ProductItem() {

	c.SuperInit()
	//获取商品的id
	id := c.Ctx.Input.Param(":id")

	//1、获取商品信息,根据id查询goods表
	goods := models.Goods{}
	models.DB.Where("id=?", id).Find(&goods)
	c.Data["goods"] = goods

	//2、获取关联商品  RelationGoods
	relationGoods := []models.Goods{}
	goods.RelationGoods = strings.ReplaceAll(goods.RelationGoods, "，", ",")
	relationIds := strings.Split(goods.RelationGoods, ",")
	models.DB.Where("id in (?)", relationIds).Select("id,title,price,goods_version").Find(&relationGoods)
	c.Data["relationGoods"] = relationGoods

	//3、获取关联赠品 GoodsGift
	goodsGift := []models.Goods{}
	goods.GoodsGift = strings.ReplaceAll(goods.GoodsGift, "，", ",")
	giftIds := strings.Split(goods.GoodsGift, ",")
	models.DB.Where("id in (?)", giftIds).Select("id,title,price,goods_img").Find(&goodsGift)
	c.Data["goodsGift"] = goodsGift

	//4、获取关联颜色 GoodsColor
	//查询商品中颜色的id列表
	goodsColor := []models.GoodsColor{}
	goods.GoodsColor = strings.ReplaceAll(goods.GoodsColor, "，", ",")
	colorIds := strings.Split(goods.GoodsColor, ",")
	models.DB.Where("id in (?)", colorIds).Find(&goodsColor)
	c.Data["goodsColor"] = goodsColor

	//5、获取关联配件 GoodsFitting
	goodsFitting := []models.Goods{}
	goods.GoodsFitting = strings.ReplaceAll(goods.GoodsFitting, "，", ",")
	fittingIds := strings.Split(goods.GoodsFitting, ",")
	models.DB.Where("id in (?)", fittingIds).Select("id,title,price,goods_img").Find(&goodsFitting)
	c.Data["goodsFitting"] = goodsFitting

	//6、获取商品关联的图片 GoodsImage
	goodsImage := []models.GoodsImage{}
	models.DB.Where("goods_id=?", goods.Id).Limit(6).Find(&goodsImage)
	c.Data["goodsImage"] = goodsImage

	//7、获取规格参数信息 GoodsAttr
	goodsAttr := []models.GoodsAttr{}
	models.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	c.Data["goodsAttr"] = goodsAttr

	//项目没有使用到这部分
	//8、获取商品其他规格参数
	/*
			颜色:红色,白色,黄色 |  尺寸:41,42,43

			[

			      { cate: "颜色", list: ["红色", "白色", "黄色"] },
			      { cate: "尺寸", list: ["41", "42", "43"] }

			 ]

			 颜色:红色,白色,黄色
				tempSlice[0]=颜色
				tempSlice[1]=红色,白色,黄色

				goodsItemAttr[0].Cate=颜色
				goodsItemAttr[0].List=切片

			 尺寸:41,42,43
			 	tempSlice[0]=尺寸
				tempSlice[1]=41,42,43

		var goodsItemAttr []models.GoodsItemAttr
		goodsAttrStr := "颜色:红色,白色,黄色|尺寸:41,42,43|套餐:套餐1,套餐2"
		goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "：", ":")
		goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "，", ",")

		if strings.Contains(goodsAttrStr, ":") {
			goodsAttrSlice := strings.Split(goodsAttrStr, "|")
			//分配存储空间
			goodsItemAttr = make([]models.GoodsItemAttr, len(goodsAttrSlice))
			for i := 0; i < len(goodsAttrSlice); i++ {
				tempSlice := strings.Split(goodsAttrSlice[i], ":")
				//分类
				goodsItemAttr[i].Cate = tempSlice[0]
				//列表
				listSlice := strings.Split(tempSlice[1], ",")
				// goodsItemAttr[i].List = append(goodsItemAttr[i].List, listSlice...)
				goodsItemAttr[i].List = listSlice

			}
		}

		c.Data["json"] = goodsItemAttr
		c.ServeJSON()

	*/

	c.TplName = "itying/product/item.html"
}

func (c *ProductController) GetImgList() {
	colorId, err1 := c.GetInt("color_id")
	goodsId, err2 := c.GetInt("goods_id")
	//查询商品图库信息
	goodsImage := []models.GoodsImage{}
	err3 := models.DB.Where("color_id=? AND goods_id=?", colorId, goodsId).Find(&goodsImage).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "失败",
			"success": false,
		}
		c.ServeJSON()
	} else {
		if len(goodsImage) == 0 {
			models.DB.Where("goods_id=?", goodsId).Find(&goodsImage)
		}
		c.Data["json"] = map[string]interface{}{
			"result":  goodsImage,
			"success": true,
		}
		c.ServeJSON()
	}
}
