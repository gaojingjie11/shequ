package service

import (
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
)

type ProductService struct{}

// ProductListReq 定义列表查询的参数结构
type ProductListReq struct {
	Page     int    `form:"page"`      // 第几页
	PageSize int    `form:"page_size"` // 每页几条
	Keyword  string `form:"keyword"`   // 搜索关键词
}

// GetList 获取商品列表 (带分页和搜索)
func (s *ProductService) GetList(req ProductListReq) ([]model.Product, int64, error) {
	var list []model.Product
	var total int64

	db := global.DB.Model(&model.Product{}).Where("status = ?", 1) // 只查上架商品

	// 1. 如果有关键词，进行模糊搜索
	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}

	// 2. 计算总数 (用于前端分页显示)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 3. 分页查询
	// Limit: 取多少条, Offset: 跳过多少条
	offset := (req.Page - 1) * req.PageSize
	if err := db.Limit(req.PageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetDetail 获取商品详情
func (s *ProductService) GetDetail(id int64) (*model.Product, error) {
	var product model.Product
	if err := global.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
