package service

import (
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
)

type NoticeService struct{}

// GetList 获取公告列表 (Limit限制条数，通常首页取前5条)
func (s *NoticeService) GetList(limit int) ([]model.Notice, error) {
	var list []model.Notice
	// 按时间倒序排列
	err := global.DB.Order("created_at desc").Limit(limit).Find(&list).Error
	return list, err
}

// GetDetail 获取详情并增加浏览量 (可选功能)
func (s *NoticeService) GetDetail(id int64) (*model.Notice, error) {
	var notice model.Notice
	if err := global.DB.First(&notice, id).Error; err != nil {
		return nil, err
	}

	// 浏览量 +1 (这里为了性能可以不加锁，或者用 Redis 计数)
	// 简单起见直接 SQL 更新
	global.DB.Model(&notice).UpdateColumn("view_count", notice.ViewCount+1)

	return &notice, nil
}
