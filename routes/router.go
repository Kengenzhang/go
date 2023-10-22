package router

// 路由
import (
	v1 "aa/api/v1"
	"aa/utils"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {
	gin.SetMode(utils.AppMode)
	e := gin.Default()
	// 创建路由表
	router_1 := e.Group("api/v1")
	{
		// user接口
		router_1.POST("user/add", v1.AddUser)
		router_1.GET("user", v1.GetUsers)
		router_1.PUT("user/:id", v1.Edituser)
		router_1.DELETE("user/:id", v1.DelUser)
		// 分类接口
		router_1.POST("category/add", v1.Addcategory)
		router_1.GET("category", v1.Getcategory)
		router_1.PUT("category/:id", v1.Editcategory)
		router_1.DELETE("category/:id", v1.Delcategory)
		// 文章接口
		router_1.POST("artucle/add", v1.Addartucle)
		router_1.GET("artucle", v1.Getartucle)
		router_1.PUT("artucle/:id", v1.Editartucle)
		router_1.DELETE("artucle/:id", v1.Delartucle)
		router_1.GET("artucle/list", v1.GETcateart)
		router_1.GET("artucle/info/:id", v1.GETartinfo)

	}
	e.Run(utils.HttpPort)
	// 路由在这个端口上跑起来
}
