package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primary"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
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
