package models


//GoodsAttr 储存商品的规格和包装
type GoodsAttr struct {
	Id              int
	GoodsId         int
	AttributeCateId int
	AttributeId     int
	AttributeTitle  string
	AttributeType   int
	AttributeValue  string
	Sort            int
	AddTime         int
	Status          int
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
