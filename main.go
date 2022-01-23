package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/bind"
	_ "github.com/yichouchou/ginPlus/routers" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
	_ "github.com/yichouchou/ginPlus/swagger" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	base := annotation.New()
	base.Dev(true)
	base.Register(engine, new(Hello), new(Example), new(bind.UserRest222))

	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	engine.Run(":8088")
}
