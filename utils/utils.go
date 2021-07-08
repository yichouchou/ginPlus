package utils

import (
	"reflect"
	"strings"
)

//todo 一个完整的注释至少包括这样的内容
// [!name string, !password string, !age int]
// @GET /block

//方法1
//校验注解内容是否包含： @GET @POST @HEAD @DELETE --- 等开头，然后表示可接收的类型
func ContainsHttpMethod(annoDoc string) (httpMethond string, isHas bool) {
	if strings.Contains(annoDoc, "GET") {
		return "GET", true
	}
	if strings.Contains(annoDoc, "POST") {
		return "POST", true
	}
	if strings.Contains(annoDoc, "DELETE") {
		return "DELETE", true
	}
	if strings.Contains(annoDoc, "PATCH") {
		return "PATCH", true
	}
	if strings.Contains(annoDoc, "PUT") {
		return "PUT", true
	}
	if strings.Contains(annoDoc, "OPTIONS") {
		return "OPTIONS", true
	}
	if strings.Contains(annoDoc, "HEAD") {
		return "HEAD", true
	}
	return "ANY", false
}

//方法2
//校验是否存在类似/hello 内容的路由
func ContainsHttpRouter(annoDoc string) (router string, contains bool) {
	if strings.Contains(annoDoc, "/") {
		//如果包含'/' 取出'/'后面的内容
		indexStart := strings.Index(annoDoc, "/")
		return annoDoc[indexStart:], true
	}
	return "", false
}

//方法3
//校验是否存在[]的内容，里面是装的参数  由于可以通过反射获取到参数的类型，
//但是无法获得传递的参数名称 比如 name string 获的不了name，所以需要使用类似[!name string, !password string, !age int]这样的方式，然后可以对于参数是否必传做校验
func ContainsBraces(annoDoc string) (parms string, contains bool) {
	if strings.Contains(annoDoc, "[") {
		indexStart := strings.Index(annoDoc, "[")
		if strings.Contains(annoDoc, "]") {
			indexEnd := strings.Index(annoDoc, "]")
			return annoDoc[indexStart:indexEnd], true
		}
	}
	return "", false
}

//解析[!name string, !password string, !age int] 组装到annotation.Parm 结构体数组
func SplitParms(parmsDoc string) (parms []Parm) {
	var parmList []Parm
	split := strings.Split(parmsDoc, ",")
	//遍历数组，组装参数然后返回
	for _, s := range split {
		var must bool
		if s[0:1] == "!" {
			must = true
		} else {
			must = false
		}
		//这里根据/n 空格获取其实很有风险-可以判断当前系统，然后再解析... 或者根据 _ - 等标识
		end := strings.Index(s, "/n")
		parmName := s[1:end]
		parmTypeStr := s[end:] //todo 这里其中是不需要获取到type的，因为在后面传递参数的时候，有办法反射获取，但是这里保留的话在考虑这样提前保留，然后后面无需运行时获取会快一些

		parmList = append(parmList, Parm{
			ParmName: parmName,
			IsMust:   must,
			ParmType: reflect.TypeOf(parmTypeStr),
		})
	}
	return parmList

}

func Kind2String(kind reflect.Kind) (kinStr string) {
	switch kind {
	case reflect.String:
		return "reflect.String"
	case reflect.Int:
		return "reflect.Int"
	case reflect.Ptr:
		return "reflect.Ptr"
	case reflect.Struct:
		return "reflect.Struct"
	case reflect.Array:
		return "reflect.Slice"
	default:
		return "reflect.Array"
	}
}
