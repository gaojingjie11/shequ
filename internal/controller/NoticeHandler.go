package controller

import (
	"smartcommunity/internal/service"
	"smartcommunity/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoticeHandler struct {
	Service service.NoticeService
}

// List 首页公告列表
func (h *NoticeHandler) List(c *gin.Context) {
	list, err := h.Service.GetList(10) // 默认取最新 10 条
	if err != nil {
		response.Fail(c, "获取失败")
		return
	}
	response.Success(c, list)
}

// Detail 公告详情
func (h *NoticeHandler) Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	notice, err := h.Service.GetDetail(id)
	if err != nil {
		response.Fail(c, "公告不存在")
		return
	}
	response.Success(c, notice)
}
