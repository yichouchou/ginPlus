package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	//base := annotation.New()
	//base.Dev(false)
	//base.Register(engine, new(Hello), new(Example), new(examples.UserRest))
	engine.Run(":8088")
}
