package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

//type Stock struct {
//	BaseModel
//	Name string
//	Address string
//}

type Inventory struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"` // 分布式锁的乐观锁
}

type InventoryNew struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"` //分布式锁的乐观锁
	Freeze  int32 `gorm:"type:int"` //冻结库存
}

//type InventoryHistory struct {
//	user int32
//	goods int32
//	nums int32
//	order int32
//	status int32 //1. 表示库存是预扣减， 幂等性， 2. 表示已经支付
//}
type GoodsDetail struct {
	Goods int32
	Num   int32
}

type GoodsDetailList []GoodsDetail

func (g GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type StockSellDetail struct {
	OrderSn string          `gorm:"type:varchar(200);index:idx_order_sn,unique;"`
	Status  int32           `gorm:"type:varchar(200)"` //1 表示已扣减 2. 表示已归还
	Detail  GoodsDetailList `gorm:"type:varchar(200)"`
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号';not null"`
	Password string     `gorm:"type:varchar(100) comment '密码';not null"`
	Nickname string     `gorm:"type:varchar(20) comment '昵称'"`
	Birthday *time.Time `gorm:"type:datetime comment '生日'"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female:女，male:男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1:注册用户，2:管理员'"`
}
