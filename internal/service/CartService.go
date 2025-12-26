package service

import (
	"errors"
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
)

type CartService struct{}

// AddCart 添加商品到购物车
func (s *CartService) AddCart(userID, productID int64, num int) error {
	// 1. 检查商品是否存在
	var product model.Product
	if err := global.DB.First(&product, productID).Error; err != nil {
		return errors.New("商品不存在")
	}

	// 2. 检查购物车里是否已经有这个商品
	var cart model.Cart
	err := global.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error

	if err == nil {
		// 情况A: 已经有了 -> 更新数量
		cart.Quantity += num
		return global.DB.Save(&cart).Error
	} else {
		// 情况B: 没有 -> 新增记录
		newCart := model.Cart{
			UserID:    userID,
			ProductID: productID,
			Quantity:  num,
		}
		return global.DB.Create(&newCart).Error
	}
}

// GetCartList 获取用户的购物车列表
func (s *CartService) GetCartList(userID int64) ([]model.Cart, error) {
	var list []model.Cart
	// Preload("Product") 会自动把关联的商品信息查出来填充进去
	err := global.DB.Preload("Product").Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

// DeleteCart 删除购物车项
func (s *CartService) DeleteCart(userID, cartID int64) error {
	return global.DB.Where("user_id = ? AND id = ?", userID, cartID).Delete(&model.Cart{}).Error
}
