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

// UpdateStatus 更新报修状态 (派单/完成)
func (s *RepairService) UpdateStatus(id int64, status int, feedback string) error {
	// 状态: 0待处理 1处理中 2已完成
	return global.DB.Model(&model.Repair{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":   status,
		"feedback": feedback, // 管理员的处理反馈
	}).Error
}

// GetAllList 管理员查看所有报修
func (s *RepairService) GetAllList(limit int) ([]model.Repair, error) {
	var list []model.Repair
	err := global.DB.Order("id desc").Limit(limit).Find(&list).Error
	return list, err
}
