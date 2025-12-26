package service

import (
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
)

type RepairService struct{}

// Create 提交报修/投诉
func (s *RepairService) Create(repair *model.Repair) error {
	// 默认状态 0 (待处理)
	repair.Status = 0
	return global.DB.Create(repair).Error
}

// GetUserList 获取我的报修记录
func (s *RepairService) GetUserList(userID int64) ([]model.Repair, error) {
	var list []model.Repair
	err := global.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&list).Error
	return list, err
}
