package controller

import (
	"smartcommunity/internal/model"
	"smartcommunity/internal/service"
	"smartcommunity/pkg/response"

	"github.com/gin-gonic/gin"
)

type RepairHandler struct {
	Service service.RepairService
}

// Create 提交接口
func (h *RepairHandler) Create(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req struct {
		Type     int    `json:"type"`     // 1报修 2投诉
		Category string `json:"category"` // 分类
		Content  string `json:"content"`  // 内容
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数错误")
		return
	}

	repair := model.Repair{
		UserID:   userID.(int64),
		Type:     req.Type,
		Category: req.Category,
		Content:  req.Content,
	}

	if err := h.Service.Create(&repair); err != nil {
		response.Fail(c, "提交失败")
		return
	}
	response.Success(c, nil)
}

// List 我的记录接口
func (h *RepairHandler) List(c *gin.Context) {
	userID, _ := c.Get("userID")

	list, err := h.Service.GetUserList(userID.(int64))
	if err != nil {
		response.Fail(c, "获取失败")
		return
	}
	response.Success(c, list)
}
