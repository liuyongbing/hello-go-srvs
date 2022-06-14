package model

// 商品分类品牌关系
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

// 重新定义表名
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}
