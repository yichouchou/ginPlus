package annotation

import (
	"context"
	"fmt"
	"ginPlus/utils"
	"github.com/xxjwxc/public/mybigcamel"
	"go/ast"
	"reflect"
	"strings"
)

type MvcType int32

const (
	GIN   MvcType = 0
	ECHO  MvcType = 1
	BEEGO MvcType = 2
	IRIS  MvcType = 3
)

type DocType int32

const (
	SWAGGER DocType = 0
	DEFAULT DocType = 1
)

type annoProcess struct {
	MvcType        MvcType             //选择哪个web容器
	docType        DocType             //选择哪个api接口文档
	checkHandler   checkHandlerService //通过反射对于参数与方法的检查--还可以包括其他检查，只涉及检查不绑定的话其实会使得后面参数绑定重复--或许可以抽象一些检查逻辑，再解析的时候，从这里先检查再绑定
	checkAnno      checkAnnoService    //ast语法树 注解，注释检查，其实也可以包括语法 命名等的检查
	parserComments parserComments      //注解解析逻辑
	handlerFuncObj handlerFuncObj      //核心：参数绑定逻辑
	errorHander    errorHandler        //各种异常的处理方法

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error
	Context     context.Context
}

type checkHandlerService interface {
}

type checkAnnoService interface {
}

type parserComments interface {
	//根据注释内容信息，装填GenComment 路由配置实体，和doc文档用的请求和响应原始数据
	parserComments(f *ast.FuncDecl, objName, objFunc string, imports map[string]string, objPkg string, num int, t reflect.Type) ([]*utils.GenComment, *utils.ParmInfo, *utils.ParmInfo)

	//格式化参数的方法  从注释获取路由等
	getDefaultComments(objName, objFunc string, num int) (routerPath string, methods []string)
}

type defaultParserComment struct {
	isBigCamel bool // big camel style.大驼峰命名规则

}

func (receiver defaultParserComment) parserComments(f *ast.FuncDecl, objName, objFunc string, imports map[string]string, objPkg string, num int, t reflect.Type) ([]*utils.GenComment, *utils.ParmInfo, *utils.ParmInfo) {
	var note string
	var gcs []*utils.GenComment
	req := analysisParm(f.Type.Params, imports, objPkg, 0)
	resp := analysisParm(f.Type.Results, imports, objPkg, 0)
	ignore := false

	//最好的方式不是从注释中取，而是从方法本身，但是由于注释/配置大于默认，所以还是从注释中拿，如果没有的话就从它方法本身去获取
	//解析 f.Type的内容，里面包含上述内容

	// 方法上所有的注解都会检查一遍,
	if f.Doc != nil {
		gc := &utils.GenComment{}
		gc.Parms = make([]*utils.Parm, f.Type.Params.NumFields())
		for _, c := range f.Doc.List {
			t := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			//在这里，查找到注解上的@GET 类似注解，为GenComment的Methods 赋值 todo 如果支持多请求方式的话还需要优化
			httpMethod, has := utils.ContainsHttpMethod(t)
			if has {
				gc.Methods = []string{httpMethod}
				//判断是否存在路由信息，有的话提取出来，
				router, isContains := utils.ContainsHttpRouter(t)
				if isContains {
					gc.RouterPath = router
				} else {
					// 判断是否以大驼峰命名风格，为GenComment的RouterPath赋值，这个是默认的写法，就是不存在路由信息注释的时候
					if receiver.isBigCamel { // big camel style.大驼峰
						gc.RouterPath = objName + "." + objFunc
					} else {
						gc.RouterPath = mybigcamel.UnMarshal(objName) + "." + mybigcamel.UnMarshal(objFunc)
					}
				}

				ignore = true
			}
			//解析入参与出参相关的注释，然后填充到gc中
			utils.ContainsParmsOrResults(t, gc)
		}

		gcs = append(gcs, gc)

	}

	//defalt  --上面的条件都不匹配的话，也还是会创建一个GenComment；添加RouterPath 和Methods 其中Methods 为any；如果是默认的话，请求头也不限制
	//默认的话，就是不写注解，请求方式也不给，那么可以支持多请求方式，会根据里面的参数数量和类型自动分析 todo 迫切性不高，代码量不小
	if len(gcs) == 0 && !ignore {
		gc := &utils.GenComment{}
		gc.RouterPath, gc.Methods = receiver.getDefaultComments(objName, objFunc, num)
		gcs = append(gcs, gc)
	}

	// add note 添加注释
	for i := 0; i < len(gcs); i++ {
		gcs[i].Note = note
	}
	// todo 如果用户未添加参数注释，则自动根据参数名称，自动绑定，当前是自动绑定
	// 根据objFunc 来检出在 f.Type.Params.List 内的入参参数名称，和返回参数名称（type不方便获取，注意存在 name, password string 它会把name放到一起去）
	for i := 0; i < len(gcs); i++ {
		if f.Type != nil {
			var temp = 0
			for index, field := range f.Type.Params.List {
				for fieldName := range field.Names {
					//lenParms := len(gcs[i].Parms)
					//如果经过注解内的参数装填完成之后，参数的parmname还是空的话，那么就通过默认的参数name去绑定
					if gcs[i].Parms[temp] == nil || gcs[i].Parms[temp].ParmName == "" {
						gcs[i].Parms[temp] = &utils.Parm{
							//为gcs下的所有的parms 赋ParmName
							ParmName:       f.Type.Params.List[index].Names[fieldName].Name,
							IsHeaderOrBody: utils.Default,
						}
					}
					temp++
				}
			}
			//todo 下方如果也出现类似：name, password string, age, year int 的返回参数，result部分不完整
			for _, fieldResult := range f.Type.Results.List {
				for resultNameIndex := range fieldResult.Names {
					gcs[i].Result = append(gcs[i].Result, &utils.Parm{
						ParmName:       fieldResult.Names[resultNameIndex].Name,
						IsHeaderOrBody: utils.Default,
					})
				}
			}
		}

		//在这里为gcs里面的GenComment的入参与出参的type 赋值
		// 这里需要谨慎，可能有误，t是方法还是 结构体参数 -如果是结构体 方法对应很多的
		for _, gc := range gcs {
			//需要考虑gc里面，部分参数有标注请求头/请求体，如果部分标准，余下部分则反之，另外：如果是get请求，则所有参数都是请求头内。（这样的话就无法支持多请求方式了）
			//todo 伪代码 如果未标注请求头，则默认赋值，另外部分标注的也都给赋值，get请求也全部赋值，不再支持多请求方式，为了统一规范
			utils.ReplenishParmsOrResults(gc)

			// 从1开始，因为调用者也会在其中的第1个位置
			for i := 1; i < t.NumIn(); i++ {
				fmt.Println(t.In(i), "--入参")
				// 在这里，遍历为gcs内的gc的Parms的入参的ParmType 赋值
				gc.Parms[i-1].ParmType = t.In(i)
				gc.Parms[i-1].ParmKind = t.In(i).Kind()
			}

			// 在这里，遍历为gcs内的gc result内的出参的ParmType 赋值 这里注意从0开始，返回参数
			for i := 0; i < t.NumOut(); i++ {
				fmt.Println(t.Out(i), "--出参")
				//todo 在这里，遍历为gcs内的gc的Parms的入参的ParmType 赋值 注意：有时候单返回值，是不存在返回值对应的name的，需要兼容 issues#6
				gc.Result[i].ParmType = t.Out(i)
				gc.Result[i].ParmKind = t.Out(i).Kind()
			}

		}
	}
	return gcs, req, resp

}

func (receiver defaultParserComment) getDefaultComments(objName, objFunc string, num int) (routerPath string, methods []string) {
	methods = []string{"ANY"}
	if num == 2 { // parm 2 , post default
		methods = []string{"post"}
	}

	if receiver.isBigCamel { // big camel style.大驼峰
		routerPath = objName + "." + objFunc
	} else {
		routerPath = mybigcamel.UnMarshal(objName) + "." + mybigcamel.UnMarshal(objFunc)
	}

	return
}

type Parseparms interface {
}

type handlerFuncObj interface {
}

type errorHandler interface {
}
