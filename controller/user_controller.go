package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

type UserController struct {
	BaseController
}

func (controller UserController) GetUserList(c *gin.Context) {
	users, err := models.GetUserList()
	if err != nil {
		controller.RenderError(c, "请求失败")
	}
	controller.RenderSuccess(c, users)
}

func (controller UserController) GetUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	var user, err = models.GetUser(id)
	if err != nil {
		// 打印错误日志，本质上需要第三方库输出到日志服务器
		fmt.Println("error detail for developer : ", err.Error())
		// 输出到终端，切换为友好提示
		controller.RenderError(c, "请求失败")
	} else {
		controller.RenderSuccess(c, user)
	}
}
