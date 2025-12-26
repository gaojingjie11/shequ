package service

import (
	"errors"
	"fmt"
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
	"time"

	"gorm.io/gorm"
)

type OrderService struct{}

// CreateOrder 下单逻辑
// cartIDs: 用户勾选的购物车记录ID列表
func (s *OrderService) CreateOrder(userID int64, storeID int64, cartIDs []int64) (*model.Order, error) {
	// 返回的数据
	var order *model.Order

	// 开启数据库事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 查询用户勾选的购物车商品
		var cartList []model.Cart
		if err := tx.Preload("Product").Where("id IN ? AND user_id = ?", cartIDs, userID).Find(&cartList).Error; err != nil {
			return err
		}
		if len(cartList) == 0 {
			return errors.New("请选择要购买的商品")
		}

		// 2. 准备订单数据
		orderNo := fmt.Sprintf("%d%d", time.Now().UnixNano(), userID) // 简单生成唯一订单号
		totalAmount := 0.0
		var orderItems []model.OrderItem

		// 3. 遍历购物车，处理库存和计算金额
		for _, cart := range cartList {
			// 校验库存并扣减 (乐观锁：利用 SQL 的原子性)
			// UPDATE pms_product SET stock = stock - ? WHERE id = ? AND stock >= ?
			result := tx.Model(&model.Product{}).
				Where("id = ? AND stock >= ?", cart.ProductID, cart.Quantity).
				UpdateColumn("stock", gorm.Expr("stock - ?", cart.Quantity))

			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("商品[%s]库存不足", cart.Product.Name)
			}

			// 累加总金额
			itemTotal := cart.Product.Price * float64(cart.Quantity)
			totalAmount += itemTotal

			// 构建订单明细
			orderItems = append(orderItems, model.OrderItem{
				ProductID: cart.ProductID,
				Price:     cart.Product.Price, // 锁定当前价格
				Quantity:  cart.Quantity,
			})
		}

		// 4. 创建订单主表记录
		order = &model.Order{
			OrderNo:     orderNo,
			UserID:      userID,
			TotalAmount: totalAmount,
			Status:      0, // 待付款

			Items:     orderItems, // GORM 会自动创建关联的 Items
			CreatedAt: time.Now(),
		}
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// 5. 清除购物车中已购买的商品
		if err := tx.Delete(&model.Cart{}, cartIDs).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return order, err
}

// GetList 获取用户订单列表
func (s *OrderService) GetList(userID int64) ([]model.Order, error) {
	var list []model.Order
	// Preload("Items.Product") 会自动加载：订单 -> 明细 -> 商品信息
	// 这样前端可以直接显示买的是什么
	err := global.DB.Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at desc"). // 按时间倒序
		Find(&list).Error
	return list, err
}

// PayOrder 支付订单 (核心逻辑)
func (s *OrderService) PayOrder(userID int64, orderID int64) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 查询订单
		var order model.Order
		if err := tx.First(&order, orderID).Error; err != nil {
			return errors.New("订单不存在")
		}

		// 2. 校验状态与归属
		if order.UserID != userID {
			return errors.New("无权操作此订单")
		}
		if order.Status != 0 { // 假设 0 是待支付
			return errors.New("订单状态不正确，无法支付")
		}

		// 3. 查询用户余额
		var user model.SysUser
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		// 4. 余额判断
		if user.Balance < order.TotalAmount {
			return errors.New("余额不足，请充值")
		}

		// 5. 扣减余额 (更新 sys_user)
		// UPDATE sys_user SET balance = balance - ? WHERE id = ?
		if err := tx.Model(&user).Update("balance", gorm.Expr("balance - ?", order.TotalAmount)).Error; err != nil {
			return err
		}

		// 6. 更新订单状态为已完成 (假设 1 代表已支付/已完成)
		// 如果你之前的表结构里删了 pay_time，这里就不更新 pay_time 了
		if err := tx.Model(&order).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}
