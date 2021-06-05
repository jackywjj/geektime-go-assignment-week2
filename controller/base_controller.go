package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
控制器基类
*/
type BaseController struct {
}

/**
无错误返回，统一将下发数据打包成对象
*/
func (controller *BaseController) RenderSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

/**
有错误返回，返回友好错误信息
*/
func (controller *BaseController) RenderError(c *gin.Context, errorStr string) {
	c.JSON(http.StatusOK, gin.H{"error": errorStr})
}
