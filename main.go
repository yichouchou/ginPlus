package main

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/bind"

	"github.com/gin-gonic/gin"
	_ "github.com/yichouchou/ginPlus/routers" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	base := annotation.New()
	base.Dev(true)
	base.Register(engine, new(Hello), new(Example), new(bind.UserRest222))
	engine.Run(":8088")
}
