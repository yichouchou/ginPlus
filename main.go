package main

import (
	"fmt"

	"ginPlus/annotation"
	"ginPlus/bind"
	"ginPlus/examples"

	//_ "ginPlus/routers" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	base := annotation.New()
	//base.Dev(false)
	base.Register(engine, new(Hello), new(Example), new(examples.UserRest))
	engine.Run(":8088")
}

// ginPlus的逻辑，读取ast树，然后把注释的关键信息放入到 genInfo 中，然后根据代码生成 gen_router.go文件，其中内容包括路由以及入参和出参以及参数校验规则

//在生产的环境，根据gen_router里面的init加载路由和关键注释信息

//todo 在生成的路由注册文件中，import 额外的内容是会报错的，1。可能 go fmt 未生效 2。go fmt 也无法除去，需要在生成前控制

//todo 生成前控制的话会很麻烦了，需要便利参数，import内容的 key是否在参数中contains，然后没有的话，再检查value是否有/,有的话找到最后的/的内容，去检查是否contains

//todo 如果 没有/ 那么检查"" 内的内容是否contains，没有的话就移除

//todo 随着功能的拓展，导入的内容可能会越来越多，需要尽可能的限制（目前先ide格式化一下）

func main2() {
	engine := gin.Default() //todo 考虑兼容 beego iris 等的的注解路由
	engine.POST("hello.hi4", func(ctx *gin.Context) {
		var temp []bind.ReqTest
		ctx.ShouldBind(&temp)
		fmt.Println("这是post请求接受结构体数组")

		for i := range temp {
			fmt.Println(temp[i])
		}
		ctx.JSON(200, 4)
	})
	engine.GET("hello.hi5", func(ctx *gin.Context) {
		var temp bind.ReqTest
		ctx.ShouldBind(&temp)
		fmt.Println("这是post请求接受结构体数组")

		ctx.JSON(200, 5)
	})
	engine.POST("hello.hi5")
	//base := annotation.New()
	//base.Dev(false)
	//base.Register(engine, new(Hello))
	engine.Run(":8088")
}
