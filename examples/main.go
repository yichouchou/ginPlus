package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/examples/controller"
	_ "github.com/yichouchou/ginPlus/examples/routers"
)

func main() {
	engine := gin.Default()
	base := annotation.New()
	base.Dev(true)
	// 使用默认路径，路由文件将生成在项目根目录的 routers 目录下
	base.Register(engine, new(controller.UserRest))
	engine.Run(":8088")
}
