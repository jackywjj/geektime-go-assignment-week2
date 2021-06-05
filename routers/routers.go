package routers

import (
	"github.com/gin-gonic/gin"
	"go-demo/controller"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/", controller.IndexHandler)
	router := r.Group("user")
	{
		// 获取所有用户列表
		router.GET("/", controller.UserController{}.GetUserList)
		// 通过用户ID,获取某个用户信息
		router.GET("/:id", controller.UserController{}.GetUser)
	}
	return r
}
																																				