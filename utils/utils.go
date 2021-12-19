package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//todo 目前这里的util太杂乱，需要拆分解耦

//todo 一个完整的注释至少包括这样的内容，如果没有，则自动绑定书写的类型 issues#7
// [!name string, !password string, !age int]
// @GET /block

//方法1
//校验注解内容是否包含： @GET @POST @HEAD @DELETE --- 等开头，然后表示可接收的类型 todo 将来可能支持多请求方式，返回的是httpmethond数组
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

//校验是否存在出参入残相关的注释，有的话返回true和parm入参数组和返回参数的数组
//一个完整的路由注释：[str1, str2, str3 *examples.DemoRest] [commentHi1 string,errHi1 error]
//如果复杂的情况，既包含请求头又包含请求体里面的内容的话，可以考虑这种写法:写两层内容，第一行是请求头，下方是请求体。如果请求体参数不止一个，则为表单方式提交
//？但是问题来了，如果只写一行，该如何认为呢？有可能是get请求，也可能会是表单提交-感觉会走不通
//另外一种方式，[]表示请求头里面的内容  { }表示请求体里面的内容 如果请求体参数很多，则为表单方式提交！ todo 采用的方式
//{name , password string, age int, hiValue bind.ReqTest}
//[hi *bind.ReqTest]
func ContainsParmsOrResults(annoDoc string, gc *GenComment) (resultGc *GenComment) {
	//先把parms里面能获得的入参获取到
	//如果存在请求头注释，解析出里面的参数，把parm里面的字段填充，比如 ParmName，IsMust，还有IsHeader，type不需要也无法获得
	//todo 请求头内的参数是从前往后
	if strings.HasPrefix(annoDoc, "[") && strings.HasSuffix(annoDoc, "]") {
		//对于注释里面的内容处理非常麻烦，思路如下：1.先除去[]，然后根据，切割，没一组对应一个parm，然后获得键值对，然后只需要取得前面的name

		//考虑到有可能有些人不按照顺序写，就是说请求参数内，请求头和请求体混淆写在一起，然后在注释进行区分，这个时候需要按照顺序排列，
		//但是根据注释是无法获得排列顺序的，所以最好的方式是返回一个parm数组，然后在后续再根据parm内的参数（）进行排序组装？
		//但是当参数名称不一致的时候，注视的parm名称与实际的parm参数名称不同，导致无法判断谁是谁，就很困难；
		//todo 解决思路，强制要求实际的parm参数按照顺序填写，前面的放请求头，后面的放请求体，且在注解内保持一致--可以完全不用关心，
		//todo 如果存在【】 则是请求头的内容，然后根据类型的name去找，然后标注为请求头的内容
		// todo 如果存在 str1 ,str2,str3 string的情况，则第一个参数绑定为str1
		//todo 由于指针类型 无法获取到parm对应的名称，导致难以绑定，所以只能强制性要求按照顺序写参数，请求头放到前面，请求体的内容放到后面
		// todo 如果不特别标注，默认的话自己来判断应该是从哪里取
		trimPrefix := strings.TrimPrefix(annoDoc, "[")
		anno := strings.TrimSuffix(trimPrefix, "]")
		split := strings.Split(anno, ",")
		for index, annoParm := range split {
			gc.Parms[index] = &Parm{
				ParmName:       annto2Parm(annoParm),
				IsHeaderOrBody: Header,
			}

		}

		//如果存在请求体相关注释，解析出里面的参数，把parm里面的字段填充，type不需要也无法获得
		//todo 请求体的参数是从后往前
	} else if strings.Contains(annoDoc, "{") && strings.Contains(annoDoc, "}") {
		trimPrefix := strings.TrimPrefix(annoDoc, "{")
		anno := strings.TrimSuffix(trimPrefix, "}")
		split := strings.Split(anno, ",")
		Reverse(&split)
		for index, annoParm := range split {
			gc.Parms[len(gc.Parms)-index-1] = &Parm{
				ParmName:       annto2Parm(annoParm),
				IsHeaderOrBody: Body,
			}

		}
	}

	return gc
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
		return "reflect.Array"
	case reflect.Interface:
		return "reflect.Interface"
	case reflect.Bool:
		return "reflect.Bool"
	default:
		return "reflect.Slice"
	}
}

//todo 写一个接口，里面定义了处理各种形式参数和请求方式的 方法-需要抽象，方法内包含两个东西：
// todo 1. tvl reflect.Value 指向具体的方法rest方法-  2. c *gin.Context 从gin.context解析参数，必不可少
// todo 3. obj reflect.Value 这个是方法调用者 call时候的第一个参数 4. v utils.GenRouterInfo 与该rest方法绑定的参数列表（入参与出参都有）

//这个时候的anno为切割之后的，比如：
// name, password string, age, year int,切割后为
// name / password string /age /year int
func annto2Parm(anno string) string {
	//删除首位连续的空白
	annoWithoutSpace := strings.TrimSpace(anno)
	extra_space := delete_extra_space(annoWithoutSpace)
	split := strings.Split(extra_space, " ")
	return split[0]
}

/*
函数名：delete_extra_space(s string) string
功  能:删除字符串中多余的空格(含tab)，有多个空格时，仅保留一个空格，同时将字符串中的tab换为空格
参  数:s string:原始字符串
返回值:string:删除多余空格后的字符串
创建时间:2018年12月3日
修订信息:
*/
func delete_extra_space(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)       //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

//数组倒序函数
func Reverse(arr *[]string) {
	var temp string
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

//这个是保留map重复的内容，取交集
func MapMerge(first, second map[string]string) map[string]string {
	n := make(map[string]string)
	for i, v := range first {
		for j, w := range second {
			if i == j {
				n[i] = w

			} else {
				if _, ok := n[i]; !ok {
					n[i] = v
				}
				if _, ok := n[j]; !ok {
					n[j] = w
				}
			}
		}
	}
	return n

}

//这个是保留map重复的内容，取合集
func MapMergeMost(first, second map[string]string) map[string]string {
	n := make(map[string]string)
	for i, v := range first {
		n[i] = v
	}
	for j, w := range second {
		n[j] = w
	}
	return n

}

func ReplenishParmsOrResults(gc *GenComment) {
	//如果是get请求，则所有参数都是请求头内
	if len(gc.Methods) == 1 && gc.Methods[0] == "GET" {
		for _, parm := range gc.Parms {
			parm.IsHeaderOrBody = Header
		}
	}

	var others HeaderOrBody
	for _, parm := range gc.Parms {
		if parm.IsHeaderOrBody == Header {
			others = Body
			break
		} else if parm.IsHeaderOrBody == Body {
			others = Header
			break
		}
	}

	for _, parm := range gc.Parms {
		if parm.IsHeaderOrBody != Header && parm.IsHeaderOrBody != Body {
			if others == Header || others == Body {
				parm.IsHeaderOrBody = others
			}
		}
	}

}

//随机生成controller对象的名称；todo 采用更加合理的方式
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes) + strconv.Itoa(rand.Int())
}
