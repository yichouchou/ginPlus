package utils

//todo 专门校验请求头  需要在 GenComment 结构体新增字段，【】数组类型，存放从注解解析出来的请求头内容

type RequestHeader struct {
	ContentType    string `header:"Content-Type"`
	AcceptEncoding string `header:"Accept-Encoding"`
	RefererReferer string `header:"RefererReferer"`
	Accept         string `header:"Accept"`
	AcceptLanguage string `header:"Accept-Language"`
	Attribute      map[string]string
}

type ResponseHeader struct {
	ContentType    string `header:"Content-Type"`
	AcceptEncoding string `header:"Accept-Encoding"`
	RefererReferer string `header:"RefererReferer"`
	Accept         string `header:"Accept"`
	AcceptLanguage string `header:"Accept-Language"`
	Attribute      map[string]string
}

// todo 自定义的特殊的rest注解，首先需要注册。根据tag为key，Attribute内为详细请求头信息，然后在获取restObj restMethod头上的注解之后，去根据注册的集合内查找是否有符合匹配，找到了的话就绑定请求头信息
// todo 需要调整目前的header/Consumes/Produces 绑定策略，
// todo 优化目前设计，headers：请求和响应头；Consumes调整为requestHeaders;Produces调整为responseHeaders,结构都如下设计：tag+Attribute
type CustomRequestHeaderUser struct {
	Tag       string `@req-custom-user`
	Attribute map[string]string
}

var FirstHeaderAnno = CustomRequestHeaderUser{
	Tag: "@req-custom-user",
	Attribute: map[string]string{
		"Content-Type": "application/Consumes",
	},
}

//定义很多的请求头常量
const ApplicationJson = "application/json"
const GZIP = "gzip"
const UTF8 = "utf-8"

//请求头可以由用户自定义修改，在utils内定义多个标准的请求头类型，用户可以直接使用标准的请求头类型配合标准的请求头注解。也可以使用非标准的请求头类型和非标准的请求头注解个性化拓展
var StandardRequestHeader = RequestHeader{
	ContentType:    ApplicationJson,
	AcceptEncoding: GZIP,
	RefererReferer: "",
	Accept:         "",
	AcceptLanguage: UTF8,
	Attribute:      map[string]string{},
}

// 注解方式 @req-custom
var CustomRequestHeaderMap = map[string]string{
	"Content-Type": "application/json",
}

// 注解方式 @resp-stander
var StandardResponseHeader = ResponseHeader{
	ContentType:    ApplicationJson,
	AcceptEncoding: GZIP,
	RefererReferer: "",
	Accept:         "",
	AcceptLanguage: UTF8,
	Attribute:      map[string]string{},
}

// 注解方式：@resp-custom
var CustomResponseHeaderMap = map[string]string{
	"Content-Type": "application/json",
}

//自定义请求头，启动的时候可以往这里添加，然后注册rest的时候提前绑定
var CustomRequestHeader = RequestHeader{}

var CustomResponseHeader = ResponseHeader{}

//todo 再提供多种不同差异的标准请求头，且配套提供注解内容，太复杂了

//提供另外一种方式的请求头绑定，通过tag用结构体属性来进行绑定，类似如下：

type ReqHeaderInfo struct {
}

type RespHeaderInfo struct {
}

// todo 下边是只支持rest obj的方式
//把 名称作为请求头的key,把tag信息作为value，其中的 , 作为分隔符；可以在dev的时候直接解析到router文件中，方便生成rest接口文档，然后生产环境直接加载
type RestObj struct {
	ReqContentType    ReqHeaderInfo `application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8`
	ReqUserAgent      ReqHeaderInfo `PostmanRuntime/7.26.8`
	ReqAcceptEncoding ReqHeaderInfo `gzip, deflate, br`

	RespContentType    RespHeaderInfo `application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8`
	RespUserAgent      RespHeaderInfo `PostmanRuntime/7.26.8`
	RespAcceptEncoding RespHeaderInfo `gzip, deflate, br`
}

// Hello ...
