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
	"strings"
	"sync"
	"text/template"
	"time"

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
)

// BaseGin  运行时存储结构体
type BaseGin struct {
	isBigCamel       bool // big camel style.大驼峰命名规则
	isDev            bool // if is development
	apiFun           NewAPIFunc
	apiType          reflect.Type
	outPath          string         // output path.输出目录
	beforeAfter      GinBeforeAfter // todo ..
	isOutDoc         bool
	recoverErrorFunc RecoverErrorFunc
}

// Option overrides behavior of Connect.  特有的不定方法参数使用..
type Option interface {
	apply(*BaseGin)
}

type optionFunc func(*BaseGin)

func (f optionFunc) apply(o *BaseGin) {
	f(o)
}

// Model use custom context //使用经典的context 应该是指gin.context 在rest路由中传入这个 todo 其实我用不到应该
func (b *BaseGin) Model(middleware NewAPIFunc) *BaseGin {
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

var serviceMapMu sync.Mutex // protects the serviceMap //保护serviceMap？为了线程安全？ todo

var consolePrint sync.Once //目前来看作用是一次性输出到控制台 把rest和func名称

var _genInfo genInfo //存储路由规则信息的结构体

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
			_genInfo.PkgImportList = imports
			funMp := myast.GetObjFunMp(astPkgs, objName)
			// ast.Print(token.NewFileSet(), astPkgs)
			// fmt.Println(b)

			refTyp := reflect.TypeOf(c)
			// Install the methods
			for m := 0; m < refTyp.NumMethod(); m++ {
				method := refTyp.Method(m)
				num, _b := b.checkHandlerFunc(method.Type /*.Interface()*/, true)
				if _b {
					if sdl, ok := funMp[method.Name]; ok {
						gcs, req, resp := b.parserComments(sdl, objName, method.Name, imports, objPkg, num, method.Type)
						if b.isOutDoc { // output doc  如果是OutDoc，则... todo 了解这里parse结构体的意义
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
	//genOutPut(b.outPath, modFile) // generate code   todo为了测试方便，暂时不生成 init文件
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

// 解析内容，目前看来主要是为了填充 路由注释信息，参数 和doc文档等 --可以在此处获得关键注释内容  todo imports 的键值对就是想要的 import信息 objPkg 应该就是包信息
func (b *BaseGin) parserComments(f *ast.FuncDecl, objName, objFunc string, imports map[string]string, objPkg string, num int, t reflect.Type) ([]*GenComment, *parmInfo, *parmInfo) {
	var note string
	var gcs []*GenComment
	req := analysisParm(f.Type.Params, imports, objPkg, 1)
	resp := analysisParm(f.Type.Results, imports, objPkg, 0)
	ignore := false

	if f.Doc != nil {
		for _, c := range f.Doc.List {
			gc := &GenComment{}
			t := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			//
			//if  {
			//
			//}

			if strings.HasPrefix(t, "@Router") {
				// t := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
				matches := routeRegex.FindStringSubmatch(t)
				if len(matches) == 3 {
					gc.RouterPath = matches[1]
					methods := matches[2]
					if methods != "-" {
						if methods == "" {
							gc.Methods = []string{"get"}
						} else {
							gc.Methods = strings.Split(methods, ",")
						}
						gcs = append(gcs, gc)
					} else {
						ignore = true
					}

				}
				// else {
				// return nil, errors.New("Router information is missing")
				// }
			} else if strings.HasPrefix(t, objFunc) { // find note
				t = strings.TrimSpace(strings.TrimPrefix(t, objFunc))
				note += t
			}
		}

	}

	//defalt
	if len(gcs) == 0 && !ignore {
		gc := &GenComment{}
		gc.RouterPath, gc.Methods = b.getDefaultComments(objName, objFunc, num)
		gcs = append(gcs, gc)
	}

	// add note 添加注释
	for i := 0; i < len(gcs); i++ {
		gcs[i].Note = note
	}
	for _, gc := range gcs {
		for i := 1; i < t.NumIn(); i++ {
			fmt.Println(t.In(i))
			//todo 在这里，整个parm其实在前面绑定参数type之前就应该有了，这里图方便，重新创建的，实际上应该遍历直接赋值就好了
			gc.Parms = append(gc.Parms, &Parm{ParmType: t.In(i)})
		}

	}
	return gcs, req, resp
}

//从结构体解析出内容，最终服务于doc文档 todo 以后填充
func (b *BaseGin) parserStruct(req, resp *parmInfo, astPkg *ast.Package, modPkg, modFile string) (r, p *mydoc.StructInfo) {
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

//todo 了解它的具体意义 目前来看是添加 路由和controller方法然后输出到控制台
func checkOnceAdd(handFunName string, gc GenComment) {
	consolePrint.Do(func() {
		serviceMapMu.Lock()
		defer serviceMapMu.Unlock()
		_genInfo.Tm = time.Now().Unix()
		_genInfo.List = []genRouterInfo{} // reset
	})

	AddGenOne(handFunName, gc)
}

// AddGenOne add one to base case 添加一个路由规则到规则列表 todo
func AddGenOne(handFunName string, gc GenComment) {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()
	_genInfo.List = append(_genInfo.List, genRouterInfo{
		HandFunName: handFunName,
		GenComment:  gc,
	})
}

//todo 生成控制台路由信息？
func genOutPut(outDir, modFile string) {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()

	b := genCode(outDir, modFile) // gen .go file

	_genInfo.Tm = time.Now().Unix()
	_data, _ := serializing.Encode(&_genInfo) // gob serialize 序列化
	_path := path.Join(tools.GetCurrentDirectory(), getRouter)
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

//控制台输出逻辑
func genCode(outDir, modFile string) bool {
	_genInfo.Tm = time.Now().Unix()
	if len(outDir) == 0 {
		outDir = modFile + "/routers/"
	}
	pkgName := getPkgName(outDir)
	//todo 这个时候的data里面的 PkgImportList 是键值对形式，非常恶心，思考下来 最好的方式就是原封不动，然后原封不动导入回去 由于键值对不好
	//在template中使用，直接拼接字符串更好，然后放list
	data := struct {
		genInfo
		PkgName string
	}{
		genInfo: _genInfo,
		PkgName: pkgName,
	}
	//for i := range data.genInfo.List {
	//	for i2 := range data.genInfo.List[i].GenComment.Parms {
	//		fmt.Println(data.genInfo.List[i].GenComment.Parms[i2])
	//	}
	//}

	tmpl, err := template.New("gen_out").Funcs(template.FuncMap{"GetStringList": GetStringList}).Parse(genTemp)
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

	// format
	exec.Command("gofmt", "-l", "-w", outDir).Output()
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

// GetStringList format string
func GetStringList(list []string) string {
	return `"` + strings.Join(list, `","`) + `"`
}

//格式化参数的方法 todo 目测是服务于注释
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
func analysisParm(f *ast.FieldList, imports map[string]string, objPkg string, n int) (parm *parmInfo) {
	if f != nil {
		if f.NumFields() > 1 {
			parm = &parmInfo{}
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
					for _, v1 := range v { //todo 第一格是方法的 refTyp.Method(m) 第二个传入结构体的 reflect.ValueOf(c)
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
func getInfo() map[string][]genRouterInfo {
	serviceMapMu.Lock()
	defer serviceMapMu.Unlock()

	genInfo := _genInfo
	if _genInfoCnf.Tm > genInfo.Tm { // config to update more than coding
		genInfo = _genInfoCnf
	}

	mp := make(map[string][]genRouterInfo, len(genInfo.List))
	for _, v := range genInfo.List {
		tmp := v
		mp[tmp.HandFunName] = append(mp[tmp.HandFunName], tmp)
	}
	return mp
}

// registerHandlerObj Multiple registration methods.获取并过滤要绑定的参数 todo 主要开发内容
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

// registerHandlerObj Multiple registration methods.获取并过滤要绑定的参数 todo 主要开发内容
func (b *BaseGin) registerHandlerObjTemp(router gin.IRoutes, httpMethod []string, relativePath, methodName string, tvl, obj reflect.Value, v genRouterInfo) error {
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

// HandlerFunc Get and filter the parameters to be bound (object call type) todo 核心开发板块
func (b *BaseGin) handlerFuncObj(tvl, obj reflect.Value, methodName string) gin.HandlerFunc { // 获取并过滤要绑定的参数(obj 对象类型)
	//tvl是方法的反射对象
	typ := tvl.Type()
	fmt.Println(typ.NumIn())
	for i := 0; i < typ.NumIn(); i++ {
		fmt.Println(typ.In(i))
	}
	//判断该方法参数数量-todo 如果是两个，则绑定上gin.context 和自定义结构体
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

	// Custom context type with request parameters .自定义的context类型,带request 请求参数  todo 当匹配不上上面的时候，执行此处call
	call, err := b.getCallObj3(tvl, obj, methodName)
	if err != nil { // Direct reporting error.
		panic(err)
	}

	return call
}

// HandlerFunc Get and filter the parameters to be bound (object call type) todo 核心开发板块
func (b *BaseGin) handlerFuncObjTemp(tvl, obj reflect.Value, methodName string, v genRouterInfo) gin.HandlerFunc { // 获取并过滤要绑定的参数(obj 对象类型)

	typ := tvl.Type()
	//输出参数数量
	fmt.Println(typ.NumIn())
	for i := 0; i < typ.NumIn(); i++ {
		//逐个输出参数类型-- 第一个方法调用者结构体
		fmt.Println(typ.In(i))
	}
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

	var ctxname = typ.In(4)
	//s := reflect.ValueOf(name).Elem()
	//s.FieldByName("Access_token").Set(reflect.ValueOf("aqweqwe"))
	//s.FieldByName("UserName").Set(reflect.ValueOf("zhangsan"))
	//s.FieldByName("Password").Set(reflect.ValueOf("qwerty"))
	//s.FieldByName("Age").Set(reflect.ValueOf(10))

	marshal, errr := json.Marshal(ctxname)
	if errr == nil {
		fmt.Printf("%s\n", marshal)
	}

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
	//	//逐个输出参数类型-- 第一个方法调用者结构体--所以从1开始 todo
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
	var reqTmp = typ.In(4) //参数是ptr类型 值类型
	//reqTmp.FieldByName("Access_token")
	value := reflect.New(reqTmp.Elem())
	//reqType.Elem()
	//value.FieldByName("Access_token").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("UserName").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("Password").Set(reflect.ValueOf("aaaa"))
	//value.FieldByName("Age").Set(reflect.ValueOf(1))

	data, err := json.Marshal(value.Interface())
	if err == nil {
		fmt.Printf("%s\n", data)
	}
	values := tvl.Call([]reflect.Value{obj, reflect.ValueOf("name"), reflect.ValueOf("password"), reflect.ValueOf(10), value})

	fmt.Println(values)

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

		name := c.Query("name")
		password := c.Query("password")
		age := c.GetInt("age")
		err := c.BindJSON(value.Interface())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value.Interface())
		}

		typ := tvl.Type()
		var reqTmp = typ.In(4)
		data, err := json.Marshal(reqTmp)
		if err == nil {
			fmt.Printf("%s\n", data)
		}
		values := tvl.Call([]reflect.Value{obj, reflect.ValueOf(name), reflect.ValueOf(password), reflect.ValueOf(age), reflect.ValueOf(reqTmp)})
		for _, value := range values {
			fmt.Println(reflect.ValueOf(value))
			c.JSON(200, reflect.ValueOf(value))
		}
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
		err = mapJson(v, c.Request.Form)
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
