package router

import (
	"smartcommunity/internal/controller"
	"smartcommunity/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 初始化 Handler
	userHandler := controller.UserHandler{}
	productHandler := controller.ProductHandler{}
	cartHandler := controller.CartHandler{}
	orderHandler := controller.OrderHandler{} // 新增
	noticeHandler := controller.NoticeHandler{}
	repairHandler := controller.RepairHandler{}
	// 开放接口 (不需要登录)
	apiGroup := r.Group("/api/v1")
	{ //用户
		apiGroup.POST("/register", userHandler.Register)
		apiGroup.POST("/login", userHandler.Login)

		// 商品 (通常游客也可以看商品，所以放在开放接口里)
		apiGroup.GET("/products", productHandler.List)      // 列表
		apiGroup.GET("/product/:id", productHandler.Detail) // 详情

	}

	// 2. 私有接口 (Private - 需要登录)
	// 使用 middleware.JWTAuth() 保护这个组
	private := r.Group("/api/v1")
	private.Use(middleware.JWTAuth())
	{
		// 购物车相关
		private.POST("/cart/add", cartHandler.Add)      // 添加
		private.GET("/cart/list", cartHandler.List)     // 列表
		private.DELETE("/cart/:id", cartHandler.Delete) // 删除
		// 订单相关
		private.POST("/order/create", orderHandler.Create) // 下单
		private.GET("/order/list", orderHandler.List)      // 订单列表
		private.POST("/order/pay", orderHandler.Pay)       // 支付订单
		// 报事报修
		private.POST("/repair/create", repairHandler.Create) // 提交
		private.GET("/repair/list", repairHandler.List)      // 查看历史// 之后可以在这里加: 订单下单、物业报修...
	}

	public := r.Group("/api/v1")
	{
		// ... 其他公开接口 ...
		public.GET("/notices", noticeHandler.List)      // 列表
		public.GET("/notice/:id", noticeHandler.Detail) // 详情
	}
}
