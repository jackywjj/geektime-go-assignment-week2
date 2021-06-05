package main

import (
	"fmt"
	"go-demo/dao"
	"go-demo/models"
	"go-demo/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出 关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.User{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:5000")); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
