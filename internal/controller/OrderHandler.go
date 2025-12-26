package controller

import (
	"smartcommunity/internal/service"
	"smartcommunity/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service service.OrderService
}

// Create 创建订单接口
func (h *OrderHandler) Create(c *gin.Context) {
	userID, _ := c.Get("userID")

	// 定义请求参数
	var req struct {
		CartIDs []int64 `json:"cart_ids"` // 购物车ID数组 [1, 2, 5]
		StoreID int64   `json:"store_id"` // 取货门店ID
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数错误")
		return
	}

	if len(req.CartIDs) == 0 {
		response.Fail(c, "未选择商品")
		return
	}

	// 调用 Service
	order, err := h.Service.CreateOrder(userID.(int64), req.StoreID, req.CartIDs)
	if err != nil {
		// 可能是库存不足，也可能是DB错误
		response.Fail(c, "下单失败: "+err.Error())
		return
	}

	// 返回订单号和总金额，方便前端跳转支付
	response.Success(c, gin.H{
		"order_no":     order.OrderNo,
		"total_amount": order.TotalAmount,
		"order_id":     order.ID,
	})
}

// List 订单列表接口
func (h *OrderHandler) List(c *gin.Context) {
	userID, _ := c.Get("userID")

	list, err := h.Service.GetList(userID.(int64))
	if err != nil {
		response.Fail(c, "获取失败")
		return
	}

	response.Success(c, list)
}

// Pay 支付接口
func (h *OrderHandler) Pay(c *gin.Context) {
	userID, _ := c.Get("userID")

	// 获取 URL 参数中的订单ID: /order/pay/:id
	idStr := c.Param("id")
	orderID, _ := strconv.ParseInt(idStr, 10, 64)

	// 或者你也可以用 JSON Body 传 ID，这里演示用 POST Body 传更加规范
	// 如果想简单点，直接用 URL 参数也行。这里为了极简，假设用 POST /order/pay 且带 JSON
	var req struct {
		OrderID int64 `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		// 兼容一下，如果 Body 没传，试试读 URL
		if orderID == 0 {
			response.Fail(c, "参数错误")
			return
		}
		req.OrderID = orderID
	}

	// 调用支付业务
	if err := h.Service.PayOrder(userID.(int64), req.OrderID); err != nil {
		response.Fail(c, "支付失败: "+err.Error())
		return
	}

	response.Success(c, gin.H{"msg": "支付成功"})
}
