package controller

import (
	"smartcommunity/internal/service"
	"smartcommunity/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service service.ProductService
}

// List 商品列表接口
func (h *ProductHandler) List(c *gin.Context) {
	var req service.ProductListReq
	// ShouldBindQuery 用于获取 URL 中的参数 (?page=1&keyword=苹果)
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, "参数错误")
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := h.Service.GetList(req)
	if err != nil {
		response.Fail(c, "获取商品列表失败")
		return
	}

	// 返回列表和总数
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Detail 商品详情接口
func (h *ProductHandler) Detail(c *gin.Context) {
	// 获取 URL 路径参数 /product/:id
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	product, err := h.Service.GetDetail(id)
	if err != nil {
		response.Fail(c, "商品不存在")
		return
	}

	response.Success(c, product)
}
