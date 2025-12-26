package model

import "gorm.io/gorm"

// Finance 资金流水表
type Finance struct {
	gorm.Model
	Type      int    `json:"type" gorm:"type:tinyint(1);not null;comment:1收入 -1支出"`
	Amount    int64  `json:"amount" gorm:"type:bigint;not null;comment:金额(分)"`
	Category  string `json:"category" gorm:"type:varchar(50);comment:分类"`
	RelatedID string `json:"relatedId" gorm:"type:varchar(50);index;comment:关联单号"`
	Remark    string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Operator  string `json:"operator" gorm:"type:varchar(50);comment:操作人"`
}

// TableName 覆盖表名
func (Finance) TableName() string {
	return "finance"
}
