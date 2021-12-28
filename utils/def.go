package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/xxjwxc/public/errors"

	"github.com/gin-gonic/gin"
	//"github.com/xxjwxc/ginrpc/api"
)

// NewAPIFunc Custom context support
type NewAPIFunc func(*gin.Context) interface{}

// RecoverErrorFunc recover 错误设置
type RecoverErrorFunc func(interface{})

// ParmInfo 参数类型描述
type ParmInfo struct {
	Pkg    string // 包名
	Type   string // 类型
	Import string // import 包
}

//注解事务保留结构体
//todo 在启动的时候指定需要数据库事务管理的对象，然后该对象下具有事务注解的方法都能够被事务管理；对事务方法的调用不再能够调，而是采用一个包装类，然后调用(类似gomonkey的mock方法操作)。类似以下代码：
/*

	//以下是一段类似使用的方式,先开启事务，然后调用，如果TrySet内返回了error 或者中途出现了异常，都会出现回滚。
	Transactional.try(Transactional.UserDao.TrySet,"zhangsan","18");




@Transactional(roolback="xxx err")
func(userDao UserDao)  (form formSource) TrySet() error{
	//下方是数据库操作
	........

	return err.new("xxxxx");

}


*/

type TransactionalComment struct {
}

// store the comment for the controller method. 生成注解路由--由ast解析出来的内容，包括RouterPath路由，note注释文档，以及rest controller方法
type GenComment struct {
	RouterPath string
	Note       string            // api注释
	Headers    map[string]string // 请求头键值对
	Methods    []string          //请求方式
	Consumes   map[string]string //请求头支持的请求内容类型 --todo 待开发
	Produces   map[string]string //响应时，响应的头上的内容类型 --todo 待开发
	Parms      []*Parm
	Result     []*Parm //组装返回参数的结构体，强烈建议，struct/基本数据类型 +err的返回方式 err是为了辨认是否为500服务器错误
}

type HeaderOrBody uint //判断是请求头内的参数还是请求体内的参数
const (
	Header  HeaderOrBody = 0
	Body    HeaderOrBody = 1
	Default HeaderOrBody = 2
)

type Parm struct {
	FiledNote string //参数说明 todo 保留字段
	ParmName  string
	Name      string
	Value     reflect.Value //新增字段，方便后续call的时候塞真正的参数
	ParmType  reflect.Type  //在注释阶段，已经塞进去了内容了
	ParmKind  reflect.Kind  //在   这个字段保存参数的种类，比如reflect.Int reflect.String  reflect.Struct 参数是什么类型（ maybe应当禁止值和接口传递，目前看起来暂时没有必要，接口未必）
	//ParmTypetype reflect.Type  //在
	//可能还需要保存对应的名字，比如string int bind.ReqTest{}
	IsMust         bool
	NewValueStr    string // 保存 创建结构体的 string 内容 例如：b := new(bind.ReqTest)
	StrInTypeOf    string // 保存 new(bind.ReqTest) 或者 *b 或者 new(error)的内容
	ParmKindStr    string // 保存kind分类的字段 reflect.String 类似这样
	NewResultStr   string // 保存 创建结构体的 string 内容 例如：b := new(bind.ReqTest)
	IsHeaderOrBody HeaderOrBody
	//ContentType  string // 传输的格式 比如：表单提交

}

//存储gen_router的路径 todo 完全不知道这个什么用途，里面内容看不到，预期是服务于生成doc
const (
	GetRouter = "/conf/gen_router.data"
)

//路由规则 正则表达式
var routeRegex = regexp.MustCompile(`@Router\s+(\S+)(?:\s+\[(\S+)\])?`)

// router style list.路由规则列表，这个是存放rest结构体上边的注解的参数信息的实体
type GenRouterInfo struct {
	GenComment  *GenComment
	HandFunName string
	RouterPath  string
	Note        string            // api注释
	Headers     map[string]string // 请求头键值对 --todo 必须是一个map
	Methods     []string          //请求方式
	Consumes    map[string]string //请求头支持的请求内容类型 --todo 待开发，是一个map
	Produces    map[string]string //响应时，响应的头上的内容类型 --todo 待开发，是一个map
}

//rest obj 信息结构体
type GenRestObjInfo struct {
	HandFunName string
	RouterPath  string
	Note        string
	Methods     []string
	Headers     map[string]string
	Consumes    map[string]string
	Produces    map[string]string
}

//路由规则信息
type GenInfo struct {
	List          []GenRouterInfo
	Tm            int64 //genout time
	PkgImportList map[string]string
	PkgImportStrs []string
}

var GenInfoCnf GenInfo

func MapJson(ptr interface{}, form map[string][]string) error {
	return mapFormByTag(ptr, form, "json")
}

func mapFormByTag(ptr interface{}, form map[string][]string, tag string) error {
	// Check if ptr is a map
	ptrVal := reflect.ValueOf(ptr)
	var pointed interface{}
	if ptrVal.Kind() == reflect.Ptr {
		ptrVal = ptrVal.Elem()
		pointed = ptrVal.Interface()
	}
	if ptrVal.Kind() == reflect.Map &&
		ptrVal.Type().Key().Kind() == reflect.String {
		if pointed != nil {
			ptr = pointed
		}
		return setFormMap(ptr, form)
	}

	return mappingByPtr(ptr, formSource(form), tag)
}

func setFormMap(ptr interface{}, form map[string][]string) error {
	el := reflect.TypeOf(ptr).Elem()

	if el.Kind() == reflect.Slice {
		ptrMap, ok := ptr.(map[string][]string)
		if !ok {
			return errors.New("cannot convert to map slices of strings")
		}
		for k, v := range form {
			ptrMap[k] = v
		}

		return nil
	}

	ptrMap, ok := ptr.(map[string]string)
	if !ok {
		return errors.New("cannot convert to map of strings")
	}
	for k, v := range form {
		ptrMap[k] = v[len(v)-1] // pick last
	}

	return nil
}

type setOptions struct {
	isDefaultExists bool
	defaultValue    string
}

type setter interface {
	TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (isSetted bool, err error)
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form formSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (isSetted bool, err error) {
	return setByForm(value, field, form, tagValue, opt)
}

type formSource map[string][]string

//todo 存在大量大量的参数校验，获取，绑定的方法，以后都抽取到外部，且使用接口，因为方便其他人修改和各自调整

func setByForm(value reflect.Value, field reflect.StructField, form map[string][]string, tagValue string, opt setOptions) (isSetted bool, err error) {
	vs, ok := form[tagValue]
	if !ok && !opt.isDefaultExists {
		return false, nil
	}

	switch value.Kind() {
	case reflect.Slice:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		return true, setSlice(vs, value, field)
	case reflect.Array:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		if len(vs) != value.Len() {
			return false, fmt.Errorf("%q is not valid value for %s", vs, value.Type().String())
		}
		return true, setArray(vs, value, field)
	default:
		var val string
		if !ok {
			val = opt.defaultValue
		}

		if len(vs) > 0 {
			val = vs[0]
		}
		return true, setWithProperType(val, value, field)
	}
}

var emptyField = reflect.StructField{}

func mappingByPtr(ptr interface{}, setter setter, tag string) error {
	_, err := mapping(reflect.ValueOf(ptr), emptyField, setter, tag)
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

var errUnknownType = errors.New("unknown type")

func setTimeDuration(val string, value reflect.Value, field reflect.StructField) error {
	d, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(d))
	return nil
}

func setWithProperType(val string, value reflect.Value, field reflect.StructField) error {
	switch value.Kind() {
	case reflect.Int:
		return setIntField(val, 0, value)
	case reflect.Int8:
		return setIntField(val, 8, value)
	case reflect.Int16:
		return setIntField(val, 16, value)
	case reflect.Int32:
		return setIntField(val, 32, value)
	case reflect.Int64:
		switch value.Interface().(type) {
		case time.Duration:
			return setTimeDuration(val, value, field)
		}
		return setIntField(val, 64, value)
	case reflect.Uint:
		return setUintField(val, 0, value)
	case reflect.Uint8:
		return setUintField(val, 8, value)
	case reflect.Uint16:
		return setUintField(val, 16, value)
	case reflect.Uint32:
		return setUintField(val, 32, value)
	case reflect.Uint64:
		return setUintField(val, 64, value)
	case reflect.Bool:
		return setBoolField(val, value)
	case reflect.Float32:
		return setFloatField(val, 32, value)
	case reflect.Float64:
		return setFloatField(val, 64, value)
	case reflect.String:
		value.SetString(val)
	case reflect.Struct:
		switch value.Interface().(type) {
		case time.Time:
			return setTimeField(val, field, value)
		}
		return json.Unmarshal(StringToBytes(val), value.Addr().Interface())
	case reflect.Map:
		return json.Unmarshal(StringToBytes(val), value.Addr().Interface())
	default:
		return errUnknownType
	}
	return nil
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}

	switch tf := strings.ToLower(timeFormat); tf {
	case "unix", "unixnano":
		tv, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}

		d := time.Duration(1)
		if tf == "unixnano" {
			d = time.Second
		}

		t := time.Unix(tv/int64(d), tv%int64(d))
		value.Set(reflect.ValueOf(t))
		return nil

	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func mapping(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	if field.Tag.Get(tag) == "-" { // just ignoring this field
		return false, nil
	}

	var vKind = value.Kind()

	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSetted, err := mapping(vPtr.Elem(), field, setter, tag)
		if err != nil {
			return false, err
		}
		if isNew && isSetted {
			value.Set(vPtr)
		}
		return isSetted, nil
	}

	if vKind != reflect.Struct || !field.Anonymous {
		ok, err := tryToSetValue(value, field, setter, tag)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	if vKind == reflect.Struct {
		tValue := value.Type()

		var isSetted bool
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			ok, err := mapping(value.Field(i), tValue.Field(i), setter, tag)
			if err != nil {
				return false, err
			}
			isSetted = isSetted || ok
		}
		return isSetted, nil
	}
	return false, nil
}

func head(str, sep string) (head string, tail string) {
	idx := strings.Index(str, sep)
	if idx < 0 {
		return str, ""
	}
	return str[:idx], str[idx+len(sep):]
}

func tryToSetValue(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	var tagValue string
	var setOpt setOptions

	tagValue = field.Tag.Get(tag)
	tagValue, opts := head(tagValue, ",")

	if tagValue == "" { // default value is FieldName
		tagValue = field.Name
	}
	if tagValue == "" { // when field is "emptyField" variable
		return false, nil
	}

	var opt string
	for len(opts) > 0 {
		opt, opts = head(opts, ",")

		if k, v := head(opt, "="); k == "default" {
			setOpt.isDefaultExists = true
			setOpt.defaultValue = v
		}
	}

	return setter.TrySet(value, field, tagValue, setOpt)
}

func setSlice(vals []string, value reflect.Value, field reflect.StructField) error {
	slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))
	err := setArray(vals, slice, field)
	if err != nil {
		return err
	}
	value.Set(slice)
	return nil
}

func setArray(vals []string, value reflect.Value, field reflect.StructField) error {
	for i, s := range vals {
		err := setWithProperType(s, value.Index(i), field)
		if err != nil {
			return err
		}
	}
	return nil
}

// // router style list.路由规则列表
// type genRouterList struct {
// 	list []genRouterInfo
// }

var (
	// Precompute the reflect type for error. Can't use error directly
	// because Typeof takes an empty interface value. This is annoying.
	typeOfError = reflect.TypeOf((*error)(nil)).Elem()
	//	ginrpc.AddGenOne("Hello.Hello", "/block", []string{"post"})
	//GenTemp = `
	//package {{.PkgName}}
	//
	//import (
	//	"ginPlus/annotation"
	//)
	//
	//func init() {
	//	annotation.SetVersion({{.Tm}})
	//	{{range .List}}annotation.AddGenOne("{{.HandFunName}}", "{{.GenComment.RouterPath}}", []string{ {{GetStringList .GenComment.Methods}} })
	//	{{end}} }
	//`

	//todo 把请求头 响应头 consumersCheck producesSet 添加到GenTemp 中

	GenTemp = `
	package {{.PkgName}}
	
	import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"
{{range $i, $v := .PkgImportStrs}}
	{{ $v}}
{{end}}
	)
	
	func init() {
		annotation.SetVersion({{.Tm}})
		
		{{range .List}}
			{{range .GenComment.Parms}}
				{{.NewValueStr}}
			{{end -}}
		{{end}}

		{{range .List}}
			{{range .GenComment.Result}}
				{{.NewResultStr}}
			{{end -}}
		{{end}}

		{{range .List}}annotation.AddGenOne("{{.HandFunName}}", utils.GenRouterInfo{
		HandFunName: "{{.HandFunName}}",
		RouterPath: "{{.RouterPath}}",
		Note:        "",
		Methods:    []string{ {{GetStringList .Methods}} },
        Headers:    map[string]string{
					{{range $key, $value := .Headers }}
            			"{{ $key}}" : "{{ $value}}",
  					{{end}}
					},
		Consumes:    map[string]string{
					{{range $key, $value := .Consumes }}
            			"{{$key}}":"{{$value}}",
  					{{end}}
					},
 		Produces:    map[string]string{
					{{range $key, $value := .Produces }}
            			"{{$key}}": "{{$value}}",
  					{{end}}
					},
		GenComment: &utils.GenComment{

					RouterPath: "{{.GenComment.RouterPath}}",

					Note:       "{{.GenComment.Note}}",

					Methods:    []string{ {{GetStringList .GenComment.Methods}} },

					Headers:     map[string]string{
								{{range $key, $value := .GenComment.Headers }}
            						"{{$key}}": "{{$value}}",
											{{end}}
								},

					Consumes:   map[string]string{
								{{range $key, $value := .GenComment.Consumes }}
            						"{{$key}}": "{{$value}}",
											{{end}}
								},

					Produces:   map[string]string{
								{{range $key, $value := .GenComment.Produces }}
            						"{{$key}}": "{{$value}}",
											{{end}}
								},

					Parms:      []*utils.Parm{
		
								{{range .GenComment.Parms}}
								{
									ParmName: "{{.ParmName}}",
									ParmType: reflect.TypeOf({{.StrInTypeOf}}),
									IsMust:   {{.IsMust}},
									ParmKind: {{.ParmKindStr}},
								},	

								{{end -}}
											},

					Result:     []*utils.Parm{	
		
								{{range .GenComment.Result}}

								{
									ParmName: "{{.ParmName}}",
									ParmType: reflect.TypeOf({{.StrInTypeOf}}),
									IsMust:   {{.IsMust}},
									ParmKind: {{.ParmKindStr}},
								},	

								{{end -}}


										},
							},
		
	})
{{end}} }
	`

	// 运行时绑定存在非常大的阻碍，结构体的绑定困难很大 尤其是使用一些自由的结构体的时候，无从获取需要注入的package --目前想到的方式，
	// 扩大获取的ast内容，然后根据interface中是否存在. 比如bind.ReqTest ，那么就去impots的内容中寻找存在匹配的import内容，如果有，则存放它的全限定名称，比如github.com/gin-gonic/gin.context
	// 如果没有import内容，则去查找到它的包名，然后存放
	// 存放方式 import的 “” 或者package 的 “”  然后parm的名称 这个目前已经有了，然后它的反射type

	// todo  动态生产 handlerFuncObjTemp 方法代码块，或者抽取到外面，动态生成这一部分代码。
	//  考虑到权限的参数绑定方式 兼容传值与传结构体， 还是需要pkg imports包

	//在dev环境，解析出注释的内容，包括
	//	name string, password string, age int,hi bind.ReqTest
)

/*   模板生成的内容大致会参照这样子，方便拓展，可能还会引入类似于swagger的接口文档

func init() {

// 这里下方imports的内容，目前已经能够拿到
	"reflect"
	"ginPlus/annotation"
	"ginPlus/bind"
	"ginPlus/utils"

// 下方创建结构体对象 需要的要素： 1。名称 2。new()
// 3.括号内的 bind.type （目前应该都可以拿到的，但是名称b类似的需要和下方传值时候使用的对应起来，可能需要使用map）
// 存储格式如下： name  refletx.type  --如果遇到值传递，把所有值传递的内容都需要在下方new出来
	b := new(bind.ReqTest)


//  下方版本设置已经有解法
	annotation.SetVersion(1625627016)

//   下方路由方法的注册，{{range .List}} 参考之前的，遍历，然后取
	annotation.AddGenOne("Hello.Hi", utils.GenComment{
		RouterPath: "hello.hi",
		Note:       "",
//  请求方式多种的话可以考虑先拼接再传string过来
		Methods:    []string{"ANY"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmKind: reflect.String,
//  下方固定格式写法reflect.TypeOf(new(string))，除非遇到传值的情况

				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "password",
				ParmKind: reflect.String,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "age",
				ParmKind: reflect.Int,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "hiValue",
				ParmKind: reflect.Struct,
				//reflect.TypeOf(new(bind.ReqTest)).Kind(), 这里 是否可以考虑直接 reflect.Struct

				ParmType: reflect.TypeOf(*b), //这里是传递值参数
				//由于在启动后不论dev 还是生产，运行后都可以加载对应参数，所以这里不用ParmType字段貌似也可以!! 在生产环境，无法做到注入 都会多一个 *  todo 确定了可以不用，因为无法很好的存放
				IsMust: false,
			},
			{
				ParmName: "hi",
				ParmKind: reflect.Ptr,
				//reflect.TypeOf(new(bind.ReqTest)).Kind(), 这里 是否可以考虑直接 reflect.Struct
				ParmType: reflect.TypeOf(new(bind.ReqTest)), //这里传递指针参数
				//由于在启动后不论dev 还是生产，运行后都可以加载对应参数，所以这里不用ParmType字段貌似也可以!! 在生产环境，无法做到注入 都会多一个 *  todo 确定了可以不用，因为无法很好的存放
				IsMust: false,
			},
		},
		Result: []*utils.Parm{
			{
				ParmName: "name",
				ParmKind: reflect.String,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "password",
				ParmKind: reflect.Interface,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},
		},
	})
}


*/
