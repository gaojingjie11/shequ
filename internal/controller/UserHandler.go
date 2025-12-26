package controller

import (
	"smartcommunity/internal/model"
	"smartcommunity/internal/service"
	"smartcommunity/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

// 这里必须要用 json:"password"，否则读不到前端传的值
type RegisterRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// Register 处理注册请求
func (h *UserHandler) Register(c *gin.Context) {

	// 使用专门的请求结构体来接收参数
	var req RegisterRequest

	// 1. 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数解析失败: "+err.Error())
		return
	}

	// 2. 基础校验
	if req.Mobile == "" || req.Password == "" {
		response.Fail(c, "手机号和密码不能为空")
		return
	}

	// 3. 将请求结构体(DTO) 转换为 数据库模型(Model)
	user := model.SysUser{
		Mobile:   req.Mobile,
		Password: req.Password, // 这里接收到了明文密码，Service层会加密
		RealName: req.RealName,
		Age:      req.Age,
		Gender:   req.Gender,
		Username: req.Username,
		Avatar:   req.Avatar,
	}

	// 4. 调用业务逻辑
	if err := h.Service.Register(&user); err != nil {
		response.Fail(c, err.Error())
		return
	}

	// 4. 返回成功 (Data 为 nil 或 返回部分用户信息)
	response.Success(c, gin.H{"uid": user.ID})
}

// Login 处理登录请求
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}

	// 1. 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数错误")
		return
	}

	// 2. 调用业务逻辑
	token, user, err := h.Service.Login(req.Mobile, req.Password)
	if err != nil {
		// 登录失败通常报 400 或 401
		response.Fail(c, err.Error())
		return
	}

	// 3. 返回成功数据
	response.Success(c, gin.H{
		"token":     token,
		"user_info": user, // 包含头像、余额等信息
	})
}
