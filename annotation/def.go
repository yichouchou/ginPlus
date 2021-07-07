package annotation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xxjwxc/public/errors"

	"github.com/gin-gonic/gin"
	//"github.com/xxjwxc/ginrpc/api"
)

/*
 Description: Support three types of interfaces
 func(*gin.Context) go-gin raw interface
 func(*Context)  Custom context type
 func(*Context,req)  Custom context type with request request request parameters
*/

//func _fun1(*gin.Context)              {}
//func _fun2(*api.Context)              {}
//func _fun3(*api.Context, interface{}) {}

// NewAPIFunc Custom context support
type NewAPIFunc func(*gin.Context) interface{}

// RecoverErrorFunc recover 错误设置
type RecoverErrorFunc func(interface{})

// parmInfo 参数类型描述
type parmInfo struct {
	Pkg    string // 包名
	Type   string // 类型
	Import string // import 包
}

// store the comment for the controller method. 生成注解路由--由ast解析出来的内容，包括RouterPath路由，note注释文档，以及rest controller方法
type GenComment struct {
	RouterPath string
	Note       string   // 注释
	Methods    []string //用不上？因为parm里面有了
	Parms      []*Parm
}

type Parm struct {
	ParmName string
	Name     string
	ParmType reflect.Type //在注释阶段，已经塞进去了内容了
	ParmKind reflect.Kind //在   这个字段保存参数的种类，比如reflect.Int reflect.String  reflect.Struct 参数是什么类型（todo maybe应当禁止指针和接口传递）
	//ParmTypetype reflect.Type  //在
	//可能还需要保存对应的名字，比如string int bind.ReqTest{}
	IsMust bool
	//PkgImport string //新增字段，参数对象有可能需要引入包--但是放到这里是不合适的，不方便取，我最终只需要便利取出来不需要和参数绑定
}

//存储gen_router的路径 todo 完全不知道这个什么用途，里面内容看不到，预期是服务于生成doc
const (
	getRouter = "/conf/gen_router.data"
)

//路由规则 正则表达式
var routeRegex = regexp.MustCompile(`@Router\s+(\S+)(?:\s+\[(\S+)\])?`)

// router style list.路由规则列表
type genRouterInfo struct {
	GenComment  GenComment
	HandFunName string
}

//路由规则信息
type genInfo struct {
	List          []genRouterInfo
	Tm            int64 //genout time
	PkgImportList map[string]string
}

var _genInfoCnf genInfo

func mapJson(ptr interface{}, form map[string][]string) error {
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

var _mmu sync.Mutex

// SetVersion user timestamp to replace version
func SetVersion(tm int64) {
	_mmu.Lock()
	defer _mmu.Unlock()
	_genInfo.Tm = tm
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
	//genTemp = `
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
	genTemp = `
	package {{.PkgName}}
	
	import (
		"ginPlus/annotation"
		"reflect"
	)
	
	func init() {
		annotation.SetVersion({{.Tm}})
		{{range .List}}annotation.AddGenOne("{{.HandFunName}}", annotation.GenComment{
		RouterPath: "{{.GenComment.RouterPath}}",
		Note:       "",
		Methods:    []string{ {{GetStringList .GenComment.Methods}} },
		Parms: []*annotation.Parm{	
		
	{{range .GenComment.Parms}}

		{
			ParmName: "{{.ParmName}}",
			ParmType: reflect.TypeOf({{.ParmType}}),
			IsMust:   {{.IsMust}},
		},	

	{{end -}}


		},
	})
		{{end}} }
	`
	//todo 运行时绑定存在非常大的阻碍，结构体的绑定困难很大 尤其是使用一些自由的结构体的时候，无从获取需要注入的package --目前想到的方式，
	//todo 扩大获取的ast内容，然后根据interface中是否存在. 比如bind.ReqTest ，那么就去impots的内容中寻找存在匹配的import内容，如果有，则存放它的全限定名称，比如github.com/gin-gonic/gin.context
	//todo 如果没有import内容，则去查找到它的包名，然后存放
	//todo 存放方式 import的 “” 或者package 的 “”  然后parm的名称 这个目前已经有了，然后它的反射type

	// todo 由于go反射无法动态赋值  将参数转为 结构体对象存在困难，目前想到的方式是通过 gen 生成动态代码来解决。
	// todo  动态生产 handlerFuncObjTemp 方法代码块，或者抽取到外面，动态生成这一部分代码。

	//在dev环境，解析出注释的内容，包括
	//	name string, password string, age int,hi bind.ReqTest
)

//
//
//func init() {
//	annotation.SetVersion(1625627016)
//	annotation.AddGenOne("Hello.Hi", annotation.GenComment{
//		RouterPath: "hello.hi",
//		Note:       "",
//		Methods:    []string{"ANY"},
//		Parms: []*annotation.Parm{
//
//
//			{
//				ParmName: "",
//				Name:reflect.TypeOf(new(string)).Name(),
//				ParmTypeX: reflect.String,
//				ParmType:reflect.TypeOf(new(string)),
//				IsMust:   false,
//			},
//
//			{
//				ParmName: "",
//				ParmTypeX: reflect.String,
//				ParmType: reflect.TypeOf(new(string)),
//				IsMust:   false,
//			},
//
//			{
//				ParmName: "",
//				ParmTypeX: reflect.Int,
//				ParmType: reflect.TypeOf(new(string)),
//				IsMust:   false,
//			},
//
//			{
//				ParmName: "",
//				ParmTypeX: reflect.TypeOf(new(bind.ReqTest)).Kind(),  //这里 是否可以考虑直接 reflect.Struct
//				ParmType: reflect.TypeOf(new(bind.ReqTest)),
//				IsMust:   false,
//			},
//		},
//	})
//}
