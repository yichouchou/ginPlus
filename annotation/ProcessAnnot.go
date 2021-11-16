package annotation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"net/http"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xxjwxc/public/errors"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/myast"
	"github.com/xxjwxc/public/mybigcamel"
	"github.com/xxjwxc/public/mydoc"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/myreflect"
	"github.com/xxjwxc/public/serializing"
	"github.com/xxjwxc/public/tools"
	"github.com/yichouchou/ginPlus/utils"
)

// BaseGin  运行时存储结构体
type BaseGin struct {
	isBigCamel       bool // big camel style.大驼峰命名规则
	isDev            bool // if is development
	apiFun           utils.NewAPIFunc
	apiType          reflect.Type
	outPath          string         // output path.输出目录
	beforeAfter      GinBeforeAfter // todo ..调用的前置方法拓展用，比如没有注视 请求方式，根据参数个数/类型自动推导
	isOutDoc         bool
	recoverErrorFunc utils.RecoverErrorFunc
}

type RestGenerator interface {

	//根据注释内容信息，装填GenComment 路由配置实体，和doc文档用的请求和响应原始数据
	parserComments(f *ast.FuncDecl, objName, objFunc string, imports map[string]string, objPkg string, num int, t reflect.Type) ([]*utils.GenComment, *utils.ParmInfo, *utils.ParmInfo)

	//根据req resp文档信息，填充请求与响应文档
	parserStruct(req, resp *utils.ParmInfo, astPkg *ast.Package, modPkg, modFile string) (r, p *mydoc.StructInfo)
}

// Option overrides behavior of Connect.  特有的不定方法参数使用..
type Option interface {
	apply(*BaseGin)
}

type optionFunc func(*BaseGin)

func (f optionFunc) apply(o *BaseGin) {
	f(o)
}

// Model use custom context //使用经典的context 指gin.context 在rest路由中传入这个  其实我用不到
func (b *BaseGin) Model(middleware utils.NewAPIFunc) *BaseGin {
	if middleware == nil { // default middleware
		middleware = NewApiFunc
	}

	b.apiFun = middleware // save callback

	rt := reflect.TypeOf(middleware(&gin.Context{}))
	if rt == nil || rt.Kind() != reflect.Ptr {
		panic("need pointer")
	}
	b.apiType = rt

	return b
}

// Dev set build is development 设置为dev模式
func (b *BaseGin) Dev(isDev bool) {
	b.isDev = isDev
}

// SetRecover set recover err call  设置下线重启？
func (b *BaseGin) SetRecover(errfun func(interface{})) {
	b.recoverErrorFunc = errfun
}

// Default new op obj
func Default() *BaseGin {
	b := new(BaseGin)
	b.Model(NewApiFunc)
	b.Dev(true)
	b.SetRecover(func(err interface{}) {
		mylog.Error(err)
	})

	return b
}

// New new customized base
func New(opts ...Option) *BaseGin {
	b := Default() // default option

	for _, o := range opts {
		o.apply(b)
	}

	b.SetRecover(func(err interface{}) {
		mylog.Error(err)
	})
	return b
}

var serviceMapMu sync.Mutex // protects the serviceMap //保护serviceMap安全

var consolePrint sync.Once //一次性输出到控制台 把rest和func名称

var _genInfo utils.GenInfo //存储路由规则信息的结构体

// SetVersion user timestamp to replace version

var _mmu sync.Mutex

func SetVersion(tm int64) {
	_mmu.Lock()
	defer _mmu.Unlock()
	_genInfo.Tm = tm
}

//处理自动路由和参数绑定的入口
func (b *BaseGin) tryGenRegister(router gin.IRoutes, cList ...interface{}) bool {
	//获取当前运行时条件
	modPkg, modFile, isFind := myast.GetModuleInfo(2)
	if !isFind {
		return false
	}

	groupPath := b.BasePath(router)
	doc := mydoc.NewDoc(groupPath)

	for _, c := range cList {
		refVal := reflect.ValueOf(c)
		t := reflect.Indirect(refVal).Type()
		objPkg := t.PkgPath()
		objName := t.Name()
		// fmt.Println(objPkg, objName)

		// find path
		objFile := myast.EvalSymlinks(modPkg, modFile, objPkg)
		// fmt.Println(objFile)

		astPkgs, _b := myast.GetAstPkgs(objPkg, objFile) // get ast trees.
		if _b {
			//获得astPkgs 之后，去除掉里面的main
			for s := range astPkgs.Files {
				if strings.Contains(s, "main.go") {
					delete(astPkgs.Files, s)
				}
			}
			imports := myast.AnalysisImport(astPkgs)
			//由于当前的imports 还存在对应controller里面其他以来pkg，所以需要剔除，必须依靠参数的 关键字信息进行剔除
			//由于有人可能会写别名，所以还需要特别考虑- 操蛋啊
			//todo 需要纠正下，路由的结构体未必都放在一起扽，这里直接赋值会导致a盖掉b
			_genInfo.PkgImportList = utils.MapMergeMost(_genInfo.PkgImportList, imports)
			funMp := myast.GetObjFunMp(astPkgs, objName)
			// ast.Print(token.NewFileSet(), astPkgs)
			// fmt.Println(b)

			refTyp := reflect.TypeOf(c)
			fmt.Println(refTyp.NumMethod(), "---有多少rest方法")
			// Install the methods
			for m := 0; m < refTyp.NumMethod(); m++ {
				method := refTyp.Method(m)
				num, _b := b.checkHandlerFunc(method.Type /*.Interface()*/, true)
				if _b {
					if sdl, ok := funMp[method.Name]; ok {
						gcs, req, resp := b.parserComments(sdl, objName, method.Name, imports, objPkg, num, method.Type)
						if b.isOutDoc { // output doc  如果是OutDoc，则...  了解这里parse结构体的意义
							docReq, docResp := b.parserStruct(req, resp, astPkgs, modPkg, modFile)
							for _, gc := range gcs {
								doc.AddOne(objName, gc.RouterPath, gc.Methods, gc.Note, docReq, docResp)
							}
						}

						for _, gc := range gcs {
							checkOnceAdd(objName+"."+method.Name, *gc)
						}
					}
				}
			}
		}
	}

	if b.isOutDoc {
		doc.GenSwagger(modFile + "/docs/swagger/")
		doc.GenMarkdown(modFile + "/docs/markdown/")
	}
	genOutPut(b.outPath, modFile) // generate code
	return true
}

//传入gin.IRoutes 获取basePath
func (b *BaseGin) BasePath(router gin.IRoutes) string {
	switch r := router.(type) {
	case *gin.RouterGroup:
		return r.BasePath()
	case *gin.Engine:
		return r.BasePath()
	}
	return ""
}

// checkHandlerFunc Judge whether to match rules  检查处理rest请求的func，获得它的参数个数和 bool？ todo 校验rest func的请求参数数量，且是否为正确传参
func (b *BaseGin) checkHandlerFunc(typ reflect.Type, isObj bool) (int, bool) { // 判断是否匹配规则,返回参数个数
	offset := 0
	if isObj {
		offset = 1
	}
	num := typ.NumIn() - offset
	if num == 1 || num == 2 { // Parameter checking 参数检查
		ctxType := typ.In(0 + offset)

		// go-gin default method
		if ctxType == reflect.TypeOf(&gin.Context{}) {
			return num, true
		}

		// Customized context . 自定义的context
		if ctxType == b.apiType {
			return num, true
		}

		// maybe interface
		if b.apiType.ConvertibleTo(ctxType) {
			return num, true
		}

	}
	return num, true
}

// 解析内容，目前看来主要是为了填充 路由注释信息，参数 和doc文档等 --可以在此处获得关键注释内容   imports 的键值对就是想要的 import信息 objPkg 应该就是包信息；注意，这里是一个restful方法
func (b *BaseGin) parserComments(f *ast.FuncDecl, objName, objFunc string, imports map[string]string, objPkg string, num int, t reflect.Type) ([]*utils.GenComment, *utils.ParmInfo, *utils.ParmInfo) {
	//for i := range f.Type.Params.List {
	//	fmt.Println(f.Type.Params.List[i].Type)
	//	fmt.Println(f.Type.Params.List[i].Names)
	//}

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
					if b.isBigCamel { // big camel style.大驼峰
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
		gc.RouterPath, gc.Methods = b.getDefaultComments(objName, objFunc, num)
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

//从结构体解析出内容，最终服务于doc文档 todo 以后填充
func (b *BaseGin) parserStruct(req, resp *utils.ParmInfo, astPkg *ast.Package, modPkg, modFile string) (r, p *mydoc.StructInfo) {
	ant := myast.NewStructAnalys(modPkg, modFile)
	if req != nil {
		tmp := astPkg
		if len(req.Pkg) > 0 {
			objFile := myast.EvalSymlinks(modPkg, modFile, req.Import)
			tmp, _ = myast.GetAstPkgs(req.Pkg, objFile) // get ast trees.
		}
		r = ant.ParserStruct(tmp, req.Type)
	}

	if resp != nil {
		tmp := astPkg
		if len(resp.Pkg) > 0 {
			objFile := myast.EvalSymlinks(modPkg, modFile, resp.Import)
			tmp, _ = myast.GetAstPkgs(resp.Pkg, objFile) // get ast trees.
		}
		p = ant.ParserStruct(tmp, resp.Type)
	}

	return
}

//todo 了解它的具体意义 目前来看是添加 路由和controller方法然后输出文档？
func checkOnceAdd(handFunName string, gc utils.GenComment) {
	consolePrint.Do(func() {
		serviceMapMu.Lock()
		defer serviceMapMu.Unlock()
		_genInfo.Tm = time.Now().Unix()
		_genInfo.List = []utils.GenRouterInfo{} // reset
	})

	AddGenOne(handFunName, gc)
}

// AddGenOne add one to base case 添加一个路由规则到规则列表
func AddGenOne(handFunName string, gc utils.GenComment) {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()
	_genInfo.List = append(_genInfo.List, utils.GenRouterInfo{
		HandFunName: handFunName,
		GenComment:  gc,
	})
}

// 生成路由文件
func genOutPut(outDir, modFile string) {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()

	b := genCode(outDir, modFile) // gen .go file

	_genInfo.Tm = time.Now().Unix()
	_data, _ := serializing.Encode(&_genInfo) // gob serialize 序列化
	_path := path.Join(tools.GetCurrentDirectory(), utils.GetRouter)
	if !b {
		tools.BuildDir(_path)
	}
	f, err := os.Create(_path)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(_data)
}

//  生成路由信息文件
func genCode(outDir, modFile string) bool {
	_genInfo.Tm = time.Now().Unix()
	if len(outDir) == 0 {
		outDir = modFile + "/routers/"
	}
	pkgName := getPkgName(outDir)
	// 这个时候的data里面的 PkgImportList 是键值对形式，非常恶心，思考下来 最好的方式就是原封不动，然后原封不动导入回去 由于键值对不好
	//  目前里面存在冗余import内容,把现在的 key value颠倒使用最佳 ginplus/reqtest reqtest 然后自动去重了 已经完成
	//在template中使用，直接拼接字符串更好，然后放list
	data := struct {
		utils.GenInfo
		PkgName string
	}{
		GenInfo: _genInfo,
		PkgName: pkgName,
	}
	//拼接 template需要import的包，键值对直接拼接为完成字符串 类似：annotation "ginPlus/annotation" todo 由于rest所在的文件内存在很多冗余import内容，需要除去，否则go build 会报错
	for s := range data.PkgImportList {
		s3 := data.PkgImportList[s]
		data.PkgImportStrs = append(data.PkgImportStrs, s+" "+"\""+s3+"\"")
	}

	for i := range data.GenInfo.List {
		parms := data.GenInfo.List[i].GenComment.Parms
		for _, parm := range parms {
			parm.ParmKindStr = utils.Kind2String(parm.ParmKind)
			//fmt.Println(parm.ParmType.Name() + "----parm.ParmType.Name()")
			fmt.Println(parm.ParmType.String() + "----parm.ParmType.String()")
			if parm.ParmKind != reflect.Ptr {
				randString := utils.RandString(10)
				//todo 由于多个rest请求的存在，会会导致name重复，建议name为关键字的拼接，或者不重复的随机数
				parm.NewValueStr = randString + " := new(" + parm.ParmType.String() + ")"
				parm.StrInTypeOf = "*" + randString
				//todo bug 只有指针类型才应该采用下方的方式，基本数据类型和结构体和数组都应当采用上方的
			} else {
				parm.NewValueStr = ""
				parm.StrInTypeOf = "new" + "(" + strings.TrimPrefix(parm.ParmType.String(), "*") + ")"
			}

		}
		results := data.GenInfo.List[i].GenComment.Result
		for _, result := range results {
			result.ParmKindStr = utils.Kind2String(result.ParmKind)
			//fmt.Println(parm.ParmType.Name() + "----parm.ParmType.Name()") //name不带前缀的包名，而string是带包名的
			fmt.Println(result.ParmType.String() + "----parm.ParmType.String()")
			if result.ParmKind != reflect.Ptr {
				randString := utils.RandString(10)
				result.NewResultStr = randString + " := new(" + result.ParmType.String() + ")"
				result.StrInTypeOf = "*" + randString
			} else {
				result.NewValueStr = ""
				result.StrInTypeOf = "new" + "(" + strings.TrimPrefix(result.ParmType.String(), "*") + ")"
			}

		}

	}
	//todo 上方传入template的结构体的思路
	// 1。要有imports的内容，这个会是一个map结构，目前已经完成
	// 2。PkgName 这个目前已经有
	// 3。GenRouterInfo 结构体信息
	// 4。目前缺少的信息，new的结构体对象以及它的名称，且这个名称需要和parm里面的信息对齐
	// 5。reflect.Int ParmKind信息，在现有的结构体存的方式无法取出类似的，应该需要另外构建，字符串，然后填充
	// 6。ParmType 现有的方式其实能够满足除了值传递的其他方式，但是为了兼容值传递，最好是把 new(bind.ReqTest) 弄成字符串传过去，然后类似*b 也就直接用
	// 7。result 返回值 目前考虑如上保持一致

	//fmt.Println(data)
	//for i := range data.genInfo.List {
	//	for i2 := range data.genInfo.List[i].GenComment.Parms {
	//		fmt.Println(data.genInfo.List[i].GenComment.Parms[i2])
	//	}
	//}

	tmpl, err := template.New("gen_out").Funcs(template.FuncMap{"GetStringList": GetStringList}).Parse(utils.GenTemp)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	f, err := os.Create(outDir + "temroute.go")
	if err != nil {
		return false
	}
	defer f.Close()
	f.Write(buf.Bytes())

	// format 格式化代码
	exec.Command("gofmt", "-l", "-s", "-w", outDir).Output()
	// goimports 移除非必要的依赖，需要安装goimports到go/bin目录下
	exec.Command("goimports", "-w", outDir).Output()
	return true
}

//获取包名称
func getPkgName(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	dir = strings.TrimRight(dir, "/")

	var pkgName string
	list := strings.Split(dir, "/")
	if len(list) > 0 {
		pkgName = list[len(list)-1]
	}

	if len(pkgName) == 0 || pkgName == "." {
		list = strings.Split(tools.GetCurrentDirectory(), "/")
		if len(list) > 0 {
			pkgName = list[len(list)-1]
		}
	}

	return pkgName
}

//  format string
func GetStringList(list []string) string {
	return `"` + strings.Join(list, `","`) + `"`
}

//格式化参数的方法  目测是服务于注释
func (b *BaseGin) getDefaultComments(objName, objFunc string, num int) (routerPath string, methods []string) {
	methods = []string{"ANY"}
	if num == 2 { // parm 2 , post default
		methods = []string{"post"}
	}

	if b.isBigCamel { // big camel style.大驼峰
		routerPath = objName + "." + objFunc
	} else {
		routerPath = mybigcamel.UnMarshal(objName) + "." + mybigcamel.UnMarshal(objFunc)
	}

	return
}

//从ast树解析出参数信息
func analysisParm(f *ast.FieldList, imports map[string]string, objPkg string, n int) (parm *utils.ParmInfo) {
	if f != nil {
		if f.NumFields() > 1 {
			parm = &utils.ParmInfo{}
			d := f.List[n].Type
			switch exp := d.(type) {
			case *ast.SelectorExpr: // 非本文件包
				parm.Type = exp.Sel.Name
				if x, ok := exp.X.(*ast.Ident); ok {
					parm.Import = imports[x.Name]
					parm.Pkg = myast.GetImportPkg(parm.Import)
				}
			case *ast.StarExpr: // 本文件
				switch expx := exp.X.(type) {
				case *ast.SelectorExpr: // 非本地包
					parm.Type = expx.Sel.Name
					if x, ok := expx.X.(*ast.Ident); ok {
						parm.Pkg = x.Name
						parm.Import = imports[parm.Pkg]
					}
				case *ast.Ident: // 本文件
					parm.Type = expx.Name
					parm.Import = objPkg // 本包
				default:
					mylog.ErrorString(fmt.Sprintf("not find any expx.(%v) [%v]", reflect.TypeOf(expx), objPkg))
				}
			case *ast.Ident: // 本文件
				parm.Type = exp.Name
				parm.Import = objPkg // 本包
			default:
				mylog.ErrorString(fmt.Sprintf("not find any exp.(%v) [%v]", reflect.TypeOf(d), objPkg))
			}
		}
	}

	if parm != nil {
		if len(parm.Pkg) > 0 {
			var pkg string
			n := strings.LastIndex(parm.Import, "/")
			if n > 0 {
				pkg = parm.Import[n+1:]
			}
			if len(pkg) > 0 {
				parm.Pkg = pkg
			}
		}
	}
	return
}

// Register Registered by struct object,[prepath + bojname.]
func (b *BaseGin) Register(router gin.IRoutes, cList ...interface{}) bool {
	if b.isDev {
		b.tryGenRegister(router, cList...)
	}

	return b.register(router, cList...)
}

// register Registered by struct object,[prepath + bojname.]
func (b *BaseGin) register(router gin.IRoutes, cList ...interface{}) bool {
	// groupPath := b.BasePath(router)
	// genRouterInfo 实际上是获取到通过init注册上去的 路由信息
	mp := getInfo()
	for _, c := range cList {
		refTyp := reflect.TypeOf(c)
		refVal := reflect.ValueOf(c)
		t := reflect.Indirect(refVal).Type()
		objName := t.Name()

		// Install the methods
		for m := 0; m < refTyp.NumMethod(); m++ {
			method := refTyp.Method(m)
			num, _b := b.checkHandlerFunc(method.Type /*.Interface()*/, true)
			if _b {
				if v, ok := mp[objName+"."+method.Name]; ok {
					for _, v1 := range v { // 第一格是方法的 refTyp.Method(m) 第二个传入结构体的 reflect.ValueOf(c)
						b.registerHandlerObjTemp(router, v1.GenComment.Methods, v1.GenComment.RouterPath, method.Name, method.Func, refVal, v1)
					}
				} else { // not find using default case
					routerPath, methods := b.getDefaultComments(objName, method.Name, num)
					b.registerHandlerObj(router, methods, routerPath, method.Name, method.Func, refVal)
				}
			}
		}
	}
	return true
}

//获取 genRouterInfo
func getInfo() map[string][]utils.GenRouterInfo {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()

	genInfo := _genInfo
	if utils.GenInfoCnf.Tm > genInfo.Tm { // config to update more than coding 替换旧版本的
		genInfo = utils.GenInfoCnf
	}

	mp := make(map[string][]utils.GenRouterInfo, len(genInfo.List))
	for _, v := range genInfo.List {
		tmp := v
		mp[tmp.HandFunName] = append(mp[tmp.HandFunName], tmp)
	}
	return mp
}

// registerHandlerObj Multiple registration methods.获取并过滤要绑定的参数  主要开发内容
func (b *BaseGin) registerHandlerObj(router gin.IRoutes, httpMethod []string, relativePath, methodName string, tvl, obj reflect.Value) error {
	call := b.handlerFuncObj(tvl, obj, methodName)

	for _, v := range httpMethod {
		// method := strings.ToUpper(v)
		// switch method{
		// case "ANY":
		// 	router.Any(relativePath,list...)
		// default:
		// 	router.Handle(method,relativePath,list...)
		// }
		// or
		switch strings.ToUpper(v) {
		case "POST":
			router.POST(relativePath, call)
		case "GET":
			router.GET(relativePath, call)
		case "DELETE":
			router.DELETE(relativePath, call)
		case "PATCH":
			router.PATCH(relativePath, call)
		case "PUT":
			router.PUT(relativePath, call)
		case "OPTIONS":
			router.OPTIONS(relativePath, call)
		case "HEAD":
			router.HEAD(relativePath, call)
		case "ANY":
			router.Any(relativePath, call)
		default:
			return errors.Errorf("method:[%v] not support", httpMethod)
		}
	}

	return nil
}

// registerHandlerObj Multiple registration methods.获取并过滤要绑定的参数  主要开发内容
func (b *BaseGin) registerHandlerObjTemp(router gin.IRoutes, httpMethod []string, relativePath, methodName string, tvl, obj reflect.Value, v utils.GenRouterInfo) error {
	call := b.handlerFuncObjTemp(tvl, obj, methodName, v)

	for _, v := range httpMethod {
		// method := strings.ToUpper(v)
		// switch method{
		// case "ANY":
		// 	router.Any(relativePath,list...)
		// default:
		// 	router.Handle(method,relativePath,list...)
		// }
		// or
		switch strings.ToUpper(v) {
		case "POST":
			router.POST(relativePath, call)
		case "GET":
			router.GET(relativePath, call)
		case "DELETE":
			router.DELETE(relativePath, call)
		case "PATCH":
			router.PATCH(relativePath, call)
		case "PUT":
			router.PUT(relativePath, call)
		case "OPTIONS":
			router.OPTIONS(relativePath, call)
		case "HEAD":
			router.HEAD(relativePath, call)
		case "ANY":
			router.Any(relativePath, call)
		default:
			return errors.Errorf("method:[%v] not support", httpMethod)
		}
	}

	return nil
}

// HandlerFunc Get and filter the parameters to be bound (object call type)  核心开发板块
func (b *BaseGin) handlerFuncObj(tvl, obj reflect.Value, methodName string) gin.HandlerFunc { // 获取并过滤要绑定的参数(obj 对象类型)
	//tvl是方法的反射对象
	typ := tvl.Type()
	fmt.Println(typ.NumIn())
	for i := 0; i < typ.NumIn(); i++ {
		fmt.Println(typ.In(i))
	}
	//判断该方法参数数量- 如果是两个，则绑定上gin.context 和自定义结构体
	if typ.NumIn() == 2 { // Parameter checking 参数检查
		ctxType := typ.In(1)

		// go-gin default method
		apiFun := func(c *gin.Context) interface{} { return c }
		if ctxType == b.apiType { // Customized context . 自定义的context
			apiFun = b.apiFun
		} else if !(ctxType == reflect.TypeOf(&gin.Context{})) {
			panic("method " + runtime.FuncForPC(tvl.Pointer()).Name() + " not support!")
		}

		return func(c *gin.Context) {
			defer func() {
				if err := recover(); err != nil {
					b.recoverErrorFunc(err)
				}
			}()
			tvl.Call([]reflect.Value{obj, reflect.ValueOf(apiFun(c))})
		}
	}
	////表示有3个参数，第一个是调用对象的结构体，然后可以在这里分段绑定
	//if typ.NumIn()==4{
	//
	//}

	// Custom context type with request parameters .自定义的context类型,带request 请求参数   当匹配不上上面的时候，执行此处call
	call, err := b.getCallObj3(tvl, obj, methodName)
	if err != nil { // Direct reporting error.
		panic(err)
	}

	return call
}

// HandlerFunc Get and filter the parameters to be bound (object call type) todo 核心开发板块
func (b *BaseGin) handlerFuncObjTemp(tvl, obj reflect.Value, methodName string, v utils.GenRouterInfo) gin.HandlerFunc { // 获取并过滤要绑定的参数(obj 对象类型)
	//使用下面这种方式可以第一次加载的时候就参数都对齐，而不是每次请求都加载一遍。

	//parmType := v.GenComment.Parms[3].ParmType  //值
	//parmType4 := v.GenComment.Parms[4].ParmType //指针
	//
	//value4 := reflect.New(parmType4.Elem()) //传指针
	//
	//vValue := reflect.New(parmType) //传值

	//results := tvl.Call([]reflect.Value{obj, reflect.ValueOf("name"), reflect.ValueOf("password"), reflect.ValueOf(10), vValue.Elem(), value4})

	//for i := range results {
	//	fmt.Println(reflect.ValueOf(results[i]))
	//}

	//typ := tvl.Type()
	////输出参数数量
	//fmt.Println(typ.NumIn())
	//for i := 0; i < typ.NumIn(); i++ {
	//	//逐个输出参数类型-- 第一个方法调用者结构体
	//	fmt.Println(typ.In(i))
	//}

	//parms := _genInfo.List[0].GenComment.Parms  //这种方式在dev环境是可以的，但是通过路由文件注册的时候，是没办法获取到对应的reflect.Type的
	//for i := range parms {
	//	fmt.Println(parms[i].ParmType)
	//	var name = parms[i].ParmType
	//	if name.Kind() == reflect.Struct {
	//		field := name.NumField()
	//		for i := 1; i < field; i++ {
	//			fmt.Println(name.Field(i))
	//			fmt.Println(name.Field(i).Type)
	//			fmt.Println(name.Field(i).Name)
	//			fmt.Println(name.Field(i).Anonymous)
	//			fmt.Println(name.Field(i).Offset)
	//			fmt.Println(name.Field(i).PkgPath)
	//			fmt.Println(name.Field(i).Tag)
	//		}
	//	}
	//	//fmt.Println(parms[i].ParmType.In())
	//}

	//var ctxname = typ.In(4)
	//s := reflect.ValueOf(name).Elem()
	//s.FieldByName("Access_token").Set(reflect.ValueOf("aqweqwe"))
	//s.FieldByName("UserName").Set(reflect.ValueOf("zhangsan"))
	//s.FieldByName("Password").Set(reflect.ValueOf("qwerty"))
	//s.FieldByName("Age").Set(reflect.ValueOf(10))
	//
	//marshal, errr := json.Marshal(ctxname)
	//if errr == nil {
	//	fmt.Printf("%s\n", marshal)
	//}

	//for i := 0; i < v.NumField(); i++ {
	//
	//	fieldInfo := v.Type().Field(i) // a reflect.StructField
	//	tag := fieldInfo.Tag           // a reflect.StructTag
	//	name := tag.Get("json")
	//
	//	if name == "" {
	//		name = strings.ToLower(fieldInfo.Name)
	//	}
	//	//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
	//	name = strings.Split(name, ",")[0]
	//
	//	if value, ok := fields[name]; ok {
	//
	//		//给结构体赋值
	//		//保证赋值时数据类型一致
	//		if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
	//			v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
	//		}
	//
	//	}
	//}

	//field := ctxname.NumField() //????我并不需要那么麻烦的操作啊，直接调用 c。getjosn或者类似方法，把对象和ctx传进去就好啊--目前只能拿到，无法赋值貌似
	//for i := 1; i < field; i++ {
	//	fmt.Println(ctxname.Field(i))
	//	fmt.Println(ctxname.Field(i).Type)
	//	fmt.Println(ctxname.Field(i).Name)
	//	fmt.Println(ctxname.Field(i).Anonymous)
	//	fmt.Println(ctxname.Field(i).Offset)
	//	fmt.Println(ctxname.Field(i).PkgPath)
	//	fmt.Println(ctxname.Field(i).Tag)
	//}
	//
	//for i := 1; i < typ.NumIn(); i++ {
	//	//逐个输出参数类型-- 第一个方法调用者结构体--所以从1开始
	//
	//}
	//
	//v2 := reflect.New(_genInfo.List[0].GenComment.Parms[0].ParmType)
	//fmt.Println(v2)
	////value := reflect.New(_genInfoCnf.List[0].GenComment.Parms[0].ParmType)
	//for i := range _genInfo.List[0].GenComment.Parms {
	//	fmt.Println(_genInfo.List[0].GenComment.Parms[i].Name)
	//	fmt.Println(_genInfo.List[0].GenComment.Parms[i].ParmType)
	//	fmt.Println(_genInfo.List[0].GenComment.Parms[i].ParmTypeX)
	//	fmt.Println(_genInfo.List[0].GenComment.Parms[i].IsMust)
	//	fmt.Println(_genInfo.List[0].GenComment.Parms[i].ParmName)
	//}
	//fmt.Println(parmType)
	//p := v2.(parmType)

	//typqwe := tvl.Type()
	//var reqTmp = typ.In(5) //参数是指针类型类型
	//reqTmp.FieldByName("Access_token")
	//value := reflect.New(reqTmp.Elem())
	//todo 如果是传值的话，是先new出来，然后再.Elem获取到值，然后传到call参数，类似下面代码

	//var reqTmp = typ.In(4) //参数是 值类型
	//valueWith := reflect.New(reqTmp)
	//with:=valueWith.Elem()  然后with就可以作为参数了- -

	//reqType.Elem()
	//value.FieldByName("Access_token").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("UserName").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("Password").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("Age").Set(reflect.ValueOf(1))

	// 下方是调用4个参数的时候的代码 第四个参数是指针，暂时注释
	//data, err := json.Marshal(value.Interface())
	//if err == nil {
	//	fmt.Printf("%s\n", data)
	//}
	//values := tvl.Call([]reflect.Value{obj, reflect.ValueOf("name"), reflect.ValueOf("password"), reflect.ValueOf(10), value})

	//fmt.Println(values)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				b.recoverErrorFunc(err)
			}
		}()
		//在参数绑定的时候，首先查询 _genInfoCnf 内的类型 和约束 比如 name string must
		//然后根据类型断言，如果是string,则 执行代码如下 c.Query("name")

		//由于会存在一些结构体传递
		//当断言到结构体的时候，我们首先要获得创建一个这样的结构体，然后填充它的成员变量

		//i := v2.Interface().(parmType)

		//value := reflect.ValueOf(_genInfoCnf.List[0].GenComment.Parms[0].ParmTypeX).Elem().Interface()
		//i := value.(_genInfoCnf.List[0].GenComment.Parms[0].ParmTypetype)

		//todo 简单的处理下参数绑定吧，复杂的需要另外弄 如果参数传了gin.context 原封不动把ctx放进去，但是返回参数不再允许
		//todo 解决有些时候参数是必传，有些时候不需要传的问题，先判断，如果没有取得就给一个默认，否则就给赋值
		for _, method := range v.GenComment.Methods {
			switch method {
			//todo 规范参数 如果是post请求，那么默认为表单方式提交？！除非指定了parm获取，但是参数只有一个的话肯定可能也不现实
			case "POST":
				//所有请求头的内容都优先处理了 todo 取出请求头里面的内容
				for index, parm := range v.GenComment.Parms {
					//var values []reflect.Value
					if parm.IsHeaderOrBody == utils.Header {
						switch parm.ParmKind {
						case reflect.String:
							value := reflect.New(parm.ParmType)
							str := c.Query(parm.ParmName)
							var s = value.Elem().Interface().(string)
							s = str
							parm.Value = reflect.ValueOf(s)
						case reflect.Int:
							parm.Value.SetInt(c.GetInt64(parm.ParmName))
						case reflect.Int64:
							parm.Value.SetInt(int64(c.GetInt(parm.ParmName)))
						case reflect.Float64:
							parm.Value.SetFloat(c.GetFloat64(parm.ParmName))
						case reflect.Float32:
							parm.Value.SetFloat(c.GetFloat64(parm.ParmName))
						case reflect.Ptr:
							value := reflect.New(v.GenComment.Parms[0].ParmType)
							err := c.ShouldBind(value.Interface())
							if err != nil {
								fmt.Println("----错误")
							}
							parm.Value = value.Elem()
						case reflect.Struct:
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							c.ShouldBind(value.Interface())
							parm.Value = value.Elem()
						}
					}
				}

				// 数组是值传递，切记  应当和指针传递区分开来
				//如果只有一个参数，那么直接默认为json body,而且是传指针的话  --数组貌似也应该是值传递 指针数组尚未尝试，忘记了这样传是否可以
				if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Ptr || v.GenComment.Parms[0].ParmKind == reflect.Slice {
					//var arr []bind.ReqTest
					//err3 := c.ShouldBind(&arr)
					//if err3!=nil {
					//	fmt.Println(err3)
					//}
					//fmt.Println(arr)

					value := reflect.New(v.GenComment.Parms[0].ParmType)
					err := c.ShouldBind(value.Interface())
					if err != nil {
						c.JSON(500, "传值错误")
					}
					//只有一个参数的话，把调用者obj和参数都传进去
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					//判断第二个参数是否为nil，为nil的话说明正常，否则服务端报错
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
					//如果post请求参数大于1，那么只能部分是请求头，部分是请求体上/表单提交部分 或者完全是表单提交里面
				} else if len(v.GenComment.Parms) > 1 {
					//首先查看请求体提交的参数的数量，大于1则查看是否为表单提交，不是的话就抛出错误
					var temp int
					for _, parm := range v.GenComment.Parms {
						if parm.IsHeaderOrBody == utils.Body {
							temp++
						}
					}
					if temp >= 1 {
						contentType := c.Request.Header.Get("Content-Type")
						//如果是表单提交（此时也是post请求） todo 更多提交方式还需补充
						if "application/x-www-form-urlencoded" == contentType {
							for index, parm := range v.GenComment.Parms {
								if parm.IsHeaderOrBody == utils.Body {
									//请求体的参数且非表单提交且参数为多个且参数类型为指针或者数组
									//todo 最终form里面的value好像存放的时候都是作为字符串存放，需要手动转一下
									if parm.ParmKind == reflect.Ptr {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										c.Bind(value.Interface())
										parm.Value = value.Elem()
									}
									// form表单本身无法传json，这里是把value json string转为json结构体
									if parm.ParmKind == reflect.Struct {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formString := c.PostForm(parm.ParmName)
										err := json.Unmarshal([]byte(formString), value.Interface())
										if err != nil {
											fmt.Println(err)
										}
										fmt.Println(value.Interface())
										parm.Value = value.Elem()
									}
									if parm.ParmKind == reflect.Array {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formArray := c.PostFormArray(parm.ParmName)
										value.SetPointer(unsafe.Pointer(&formArray))
										parm.Value = value.Elem()
									}
									if parm.ParmKind == reflect.String {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formString := c.PostForm(parm.ParmName)
										fmt.Println(formString)
										var s = value.Elem().Interface().(string)
										s = formString
										parm.Value = reflect.ValueOf(s)
									}
									if parm.ParmKind == reflect.Int {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formInt := c.PostForm(parm.ParmName)
										atoi, err := strconv.Atoi(formInt)
										if err != nil {
											fmt.Println("--转换错误")
										}
										value.SetInt(int64(atoi))
										parm.Value = value.Elem()
									}
									if parm.ParmKind == reflect.Int64 {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formInt := c.PostForm(parm.ParmName)
										atoi, err := strconv.ParseInt(formInt, 10, 64)
										if err != nil {
											fmt.Println("--转换错误")
										}
										value.SetInt(int64(atoi))
										parm.Value = value.Elem()
									}
									if parm.ParmKind == reflect.Float64 {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formInt := c.PostForm(parm.ParmName)
										atoi, err := strconv.ParseFloat(formInt, 64)
										if err != nil {
											fmt.Println("--转换错误")
										}
										value.SetFloat(atoi)
										parm.Value = value.Elem()
									}
									if parm.ParmKind == reflect.Float32 {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formInt := c.PostForm(parm.ParmName)
										atoi, err := strconv.ParseFloat(formInt, 32)
										if err != nil {
											fmt.Println("--转换错误")
										}
										value.SetFloat(atoi)
										parm.Value = value.Elem()
									}
									//不确定是否能用
									if parm.ParmKind == reflect.Map {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										formMap, err := c.GetPostFormMap(parm.ParmName)
										if !err {
											fmt.Println(err, "无法转换")
										}
										value.SetPointer(unsafe.Pointer(&formMap))
										parm.Value = value.Elem()
									}

								}
							}

							//不是表单提交,然后请求体参数也是1
						} else if temp == 1 && c.Request.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
							for index, parm := range v.GenComment.Parms {
								if parm.IsHeaderOrBody == utils.Body {
									//请求体的参数且非表单提交且参数为1且参数类型为指针或者数组
									if parm.ParmKind == reflect.Ptr || parm.ParmKind == reflect.Array {
										value := reflect.New(v.GenComment.Parms[index].ParmType)
										err := c.ShouldBind(value.Interface())
										//todo 这种方式很不恰当，后续纠正
										if err != nil {
											c.JSON(500, "传参错误")
										}
										parm.Value = value.Elem()
									} else if parm.ParmKind == reflect.Struct {
										value := reflect.New(parm.ParmType)
										err := c.ShouldBind(value.Interface())
										if err != nil {
											fmt.Println(err)
										}
										parm.Value = value.Elem()
									}
								}
							}

						} else {
							c.JSON(500, "---参数过多，不支持的类型")
						}
					}

					//遍历parms，如果是请求头的，就根据类型去请求头拿，如果是请求体的，查看请求体的数量，数量大于1则为表单提交，等于1则为body内容

				}
				var values []reflect.Value
				values = append(values, obj)
				for _, parm := range v.GenComment.Parms {
					values = append(values, parm.Value)
				}
				results := tvl.Call(values)

				c.JSON(200, results[0].Interface())

				//get 请求也是可以表单提交的，这里理解有误，可能需要后续更正 todo
			//todo 如果是get请求，那么参数只能从url中获取，ShouldBind非常友好，貌似一样的用，如果是单个结构体对象的话，get也是可以的，数组结构体
			case "GET":
				//for i, parm := range v.GenComment.Parms {
				//
				//}

				//结构体指针的时候这样足够，类似：[req *bind.ReqTest]
				if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Ptr {
					value := reflect.New(v.GenComment.Parms[0].ParmType.Elem())
					err := c.ShouldBind(value.Interface())
					if err != nil {
						fmt.Println(err)
					}
					values := tvl.Call([]reflect.Value{obj, value})
					//判断第二个参数是否为nil，为nil的话说明正常，否则服务端报错
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
					//如果是一个参数且这个参数为结构体对象，类似：[reqList bind.ReqTest]
				} else if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Struct {
					value := reflect.New(v.GenComment.Parms[0].ParmType)
					err := c.ShouldBind(value.Interface())
					if err != nil {
						c.JSON(500, "传值错误")
					}
					//只有一个参数的话，把调用者obj和参数都传进去
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					//判断第二个参数是否为nil，为nil的话说明正常，否则服务端报错
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
				} else if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.String {
					value := reflect.New(v.GenComment.Parms[0].ParmType)
					//当SetString
					value.SetString(c.Query(v.GenComment.Parms[0].ParmName))
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
				} else if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Int {
					value := reflect.New(v.GenComment.Parms[0].ParmType)
					//当SetInt里面类型对不上会宕机
					value.SetInt(int64(c.GetInt(v.GenComment.Parms[0].ParmName)))
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
				} else if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Int64 {
					value := reflect.New(v.GenComment.Parms[0].ParmType)
					//当SetInt里面类型对不上会宕机
					value.SetInt(c.GetInt64(v.GenComment.Parms[0].ParmName))
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
				} else if len(v.GenComment.Parms) == 1 && v.GenComment.Parms[0].ParmKind == reflect.Float64 {
					value := reflect.New(v.GenComment.Parms[0].ParmType)
					//当SetInt里面类型对不上会宕机
					value.SetFloat(c.GetFloat64(v.GenComment.Parms[0].ParmName))
					values := tvl.Call([]reflect.Value{obj, value.Elem()})
					valueOf := values[1].Interface()
					if valueOf != nil {
						c.JSON(500, valueOf)
					} else {
						c.JSON(200, values[0].Interface())
					}
					//当参数大于一的时候，里面如果是基本数据类型，通过反射赋值
				} else if len(v.GenComment.Parms) > 1 {
					//var values []reflect.Value
					//values = append(values, obj)
					for index, parm := range v.GenComment.Parms {
						if parm.ParmKind == reflect.Float64 {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							value.SetFloat(c.GetFloat64(v.GenComment.Parms[index].ParmName))
							parm.Value = value.Elem()
						} else if parm.ParmKind == reflect.Int64 {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							value.SetInt(c.GetInt64(v.GenComment.Parms[index].ParmName))
							parm.Value = value.Elem()
						} else if parm.ParmKind == reflect.Int {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							getInt := c.Query(v.GenComment.Parms[index].ParmName)
							var s = value.Elem().Interface().(int)
							atoi, err := strconv.Atoi(getInt)
							if err != nil {
								fmt.Println(err)
							}
							s = atoi
							parm.Value = reflect.ValueOf(s)
						} else if parm.ParmKind == reflect.String {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							str := c.Query(v.GenComment.Parms[index].ParmName)
							var s = value.Elem().Interface().(string)
							s = str
							parm.Value = reflect.ValueOf(s)
						}
					}
					var values []reflect.Value
					values = append(values, obj)
					for _, parm := range v.GenComment.Parms {
						values = append(values, parm.Value)
					}
					results := tvl.Call(values)

					//c.JSON(200, results[0].Interface())
					if len(results) == 2 {
						valueOut := results[1].Interface()
						if valueOut != nil {
							c.JSON(500, valueOut)
						} else {
							c.JSON(200, results[0].Interface())
						}
					} else {
						valueOut := results[0].Interface()
						c.JSON(200, valueOut)
					}

				}
				//var values []reflect.Value
				//values = append(values, obj)
				//for _, parm := range v.GenComment.Parms {
				//	values = append(values, parm.Value)
				//}
				//results := tvl.Call(values)
				////todo 更友好的处理rest的返回结果
				//c.JSON(200, results[0].Interface())

			case "DELETE":

			case "PATCH":

			case "PUT":

			case "OPTIONS":

			case "HEAD":

			case "ANY":

			default:
				panic("匹配不到路由方法")
			}
		}

		//name := c.Query("name")
		//password := c.Query("password")
		//age := c.GetInt("age")
		//err := c.BindJSON(value.Interface())
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	fmt.Println(value.Interface())
		//}
		//
		//typ := tvl.Type()
		//var reqTmp = typ.In(4)
		//data, err := json.Marshal(reqTmp)
		//if err == nil {
		//	fmt.Printf("%s\n", data)
		//}
		//values := tvl.Call([]reflect.Value{obj, reflect.ValueOf(name), reflect.ValueOf(password), reflect.ValueOf(age), reflect.ValueOf(reqTmp)})
		//for _, value := range values {
		//	fmt.Println(reflect.ValueOf(value))
		//	c.JSON(200, reflect.ValueOf(value))
		//}
	}

	return nil
}

// Custom context type with request parameters
func (b *BaseGin) getCallObj3(tvl, obj reflect.Value, methodName string) (func(*gin.Context), error) {
	typ := tvl.Type()
	//if typ.NumIn() != 3 { // Parameter checking 参数检查
	//	return nil, errors.New("method " + runtime.FuncForPC(tvl.Pointer()).Name() + " not support!")
	//}
	//
	//if typ.NumOut() != 0 {
	//	if typ.NumOut() == 2 { // Parameter checking 参数检查
	//		if returnType := typ.Out(1); returnType != typeOfError {
	//			return nil, errors.Errorf("method : %v , returns[1] %v not error",
	//				runtime.FuncForPC(tvl.Pointer()).Name(), returnType.String())
	//		}
	//	} else {
	//		return nil, errors.Errorf("method : %v , Only 2 return values (obj, error) are supported", runtime.FuncForPC(tvl.Pointer()).Name())
	//	}
	//}

	ctxType, reqType := typ.In(1), typ.In(2)

	reqIsGinCtx := false
	if ctxType == reflect.TypeOf(&gin.Context{}) {
		reqIsGinCtx = true
	}

	// ctxType != reflect.TypeOf(gin.Context{}) &&
	// ctxType != reflect.Indirect(reflect.ValueOf(b.iAPIType)).Type()
	if !reqIsGinCtx && ctxType != b.apiType && !b.apiType.ConvertibleTo(ctxType) {
		return nil, errors.New("method " + runtime.FuncForPC(tvl.Pointer()).Name() + " first parm not support!")
	}

	reqIsValue := true
	if reqType.Kind() == reflect.Ptr {
		reqIsValue = false
	}
	apiFun := func(c *gin.Context) interface{} { return c }
	if !reqIsGinCtx {
		apiFun = b.apiFun
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				b.recoverErrorFunc(err)
			}
		}()

		req := reflect.New(reqType)
		if !reqIsValue {
			req = reflect.New(reqType.Elem())
		}
		if err := b.unmarshal(c, req.Interface()); err != nil { // Return error message.返回错误信息
			b.handErrorString(c, req, err)
			return
		}

		if reqIsValue {
			req = req.Elem()
		}

		bainfo, is := b.beforCall(c, tvl, obj, req.Interface(), methodName)
		if !is {
			c.JSON(http.StatusBadRequest, bainfo.Resp)
			return
		}

		// var returnValues []reflect.Value
		returnValues := tvl.Call([]reflect.Value{obj, reflect.ValueOf(apiFun(c)), req})

		if returnValues != nil {
			bainfo.Resp = returnValues[0].Interface()
			rerr := returnValues[1].Interface()
			if rerr != nil {
				bainfo.Error = rerr.(error)
			}

			is = b.afterCall(bainfo, obj)
			if is {
				c.JSON(http.StatusOK, bainfo.Resp)
			} else {
				c.JSON(http.StatusBadRequest, bainfo.Resp)
			}
		}
	}, nil
}

// Custom context type with request parameters
func (b *BaseGin) getCallObj3Temp(tvl, obj reflect.Value, methodName string) (func(*gin.Context), error) {
	typ := tvl.Type()
	//if typ.NumIn() != 3 { // Parameter checking 参数检查
	//	return nil, errors.New("method " + runtime.FuncForPC(tvl.Pointer()).Name() + " not support!")
	//}
	//
	//if typ.NumOut() != 0 {
	//	if typ.NumOut() == 2 { // Parameter checking 参数检查
	//		if returnType := typ.Out(1); returnType != typeOfError {
	//			return nil, errors.Errorf("method : %v , returns[1] %v not error",
	//				runtime.FuncForPC(tvl.Pointer()).Name(), returnType.String())
	//		}
	//	} else {
	//		return nil, errors.Errorf("method : %v , Only 2 return values (obj, error) are supported", runtime.FuncForPC(tvl.Pointer()).Name())
	//	}
	//}

	ctxType, reqType := typ.In(1), typ.In(2)

	reqIsGinCtx := false
	if ctxType == reflect.TypeOf(&gin.Context{}) {
		reqIsGinCtx = true
	}

	// ctxType != reflect.TypeOf(gin.Context{}) &&
	// ctxType != reflect.Indirect(reflect.ValueOf(b.iAPIType)).Type()
	if !reqIsGinCtx && ctxType != b.apiType && !b.apiType.ConvertibleTo(ctxType) {
		return nil, errors.New("method " + runtime.FuncForPC(tvl.Pointer()).Name() + " first parm not support!")
	}

	reqIsValue := true
	if reqType.Kind() == reflect.Ptr {
		reqIsValue = false
	}
	apiFun := func(c *gin.Context) interface{} { return c }
	if !reqIsGinCtx {
		apiFun = b.apiFun
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				b.recoverErrorFunc(err)
			}
		}()

		req := reflect.New(reqType)
		if !reqIsValue {
			req = reflect.New(reqType.Elem())
		}
		if err := b.unmarshal(c, req.Interface()); err != nil { // Return error message.返回错误信息
			b.handErrorString(c, req, err)
			return
		}

		if reqIsValue {
			req = req.Elem()
		}

		bainfo, is := b.beforCall(c, tvl, obj, req.Interface(), methodName)
		if !is {
			c.JSON(http.StatusBadRequest, bainfo.Resp)
			return
		}

		// var returnValues []reflect.Value
		returnValues := tvl.Call([]reflect.Value{obj, reflect.ValueOf(apiFun(c)), req})

		if returnValues != nil {
			bainfo.Resp = returnValues[0].Interface()
			rerr := returnValues[1].Interface()
			if rerr != nil {
				bainfo.Error = rerr.(error)
			}

			is = b.afterCall(bainfo, obj)
			if is {
				c.JSON(http.StatusOK, bainfo.Resp)
			} else {
				c.JSON(http.StatusBadRequest, bainfo.Resp)
			}
		}
	}, nil
}

//参数绑定逻辑
func (b *BaseGin) unmarshal(c *gin.Context, v interface{}) error {
	err := c.ShouldBind(v)
	if err != nil || strings.EqualFold(c.Request.Method, "get") { // get 模式 补刀json
		err = utils.MapJson(v, c.Request.Form)
	}
	return err
}

//处理error 报错的方法，500错误，然后err信息
func (b *BaseGin) handErrorString(c *gin.Context, req reflect.Value, err error) {
	var fields []string
	if _, ok := err.(validator.ValidationErrors); ok {
		for _, err := range err.(validator.ValidationErrors) {
			tmp := fmt.Sprintf("%v:%v", myreflect.FindTag(req.Interface(), err.Field(), "json"), err.Tag())
			if len(err.Param()) > 0 {
				tmp += fmt.Sprintf("[%v](but[%v])", err.Param(), err.Value())
			}
			fields = append(fields, tmp)
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			// fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
		}
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		err := err.(*json.UnmarshalTypeError)
		tmp := fmt.Sprintf("%v:%v(but[%v])", err.Field, err.Type.String(), err.Value)
		fields = append(fields, tmp)

	} else {
		fields = append(fields, err.Error())
	}

	msg := message.GetErrorMsg(message.ParameterInvalid)
	msg.Error = fmt.Sprintf("req param : %v", strings.Join(fields, ";"))
	c.JSON(http.StatusBadRequest, msg)
}

//调用具体controller方法前执行的放啊
func (b *BaseGin) beforCall(c *gin.Context, tvl, obj reflect.Value, req interface{}, methodName string) (*GinBeforeAfterInfo, bool) {
	info := &GinBeforeAfterInfo{
		C:        c,
		FuncName: fmt.Sprintf("%v.%v", reflect.Indirect(obj).Type().Name(), methodName), // 函数名
		Req:      req,                                                                   // 调用前的请求参数
		Context:  context.Background(),                                                  // 占位参数，可用于存储其他参数，前后连接可用
	}

	is := true
	if bfobj, ok := obj.Interface().(GinBeforeAfter); ok { // 本类型
		is = bfobj.GinBefore(info)
	}
	if is && b.beforeAfter != nil {
		is = b.beforeAfter.GinBefore(info)
	}
	return info, is
}

//掉用controller方法后执行的逻辑
func (b *BaseGin) afterCall(info *GinBeforeAfterInfo, obj reflect.Value) bool {
	is := true
	if bfobj, ok := obj.Interface().(GinBeforeAfter); ok { // 本类型
		is = bfobj.GinAfter(info)
	}
	if is && b.beforeAfter != nil {
		is = b.beforeAfter.GinAfter(info)
	}
	return is
}
