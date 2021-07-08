package main

import (
	"ginPlus/annotation"
	_ "ginPlus/routers" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	base := annotation.New()
	base.Dev(false)
	base.Register(engine, new(Hello))
	engine.Run(":8088")
}

//todo ginPlus的逻辑，读取ast树，然后把注释的关键信息放入到 genInfo 中，然后根据代码生成 gen_router.go文件，其中内容包括路由以及入参和出参以及参数校验规则

//在生产的环境，根据gen_router里面的init加载路由和关键注释信息

//todo 在register的时候，registerHandlerObjTemp 在里面的方法中进行参数绑定和controller方法的生成，其中根据 genInfo 信息内的参数，获取string / josn /object对象等等 然后在返回参数也渲染上

//todo 支持【】传参  比如 []Hello{}
