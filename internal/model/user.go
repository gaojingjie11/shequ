package model

import "time"

// SysUser 对应数据库 sys_user 表 [cite: 39]
type SysUser struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // JSON返回时隐藏密码
	RealName  string    `json:"real_name"`
	Mobile    string    `gorm:"unique" json:"mobile"` // 手机号唯一
	Age       int       `json:"age"`
	Gender    int       `json:"gender"`
	Avatar    string    `json:"avatar"`
	Balance   float64   `json:"balance"`
	Role      string    `json:"role"`   // user, admin
	Status    int       `json:"status"` // 1正常 0冻结
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
