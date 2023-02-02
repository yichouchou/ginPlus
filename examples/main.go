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
	// todo fix this outpath
	base.OutPath("D:\\goProject\\ginPlus\\examples\\routers\\")
	base.Register(engine, new(controller.UserRest))
	engine.Run(":8088")
}
