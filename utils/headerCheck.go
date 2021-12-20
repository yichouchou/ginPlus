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

var StandardResponseHeader = ResponseHeader{
	ContentType:    ApplicationJson,
	AcceptEncoding: GZIP,
	RefererReferer: "",
	Accept:         "",
	AcceptLanguage: UTF8,
	Attribute:      map[string]string{},
}

//自定义请求头，启动的时候可以往这里添加，然后注册rest的时候提前绑定
var CustomRequestHeader = RequestHeader{}

var CustomResponseHeader = ResponseHeader{}

//todo 再提供多种不同差异的标准请求头，且配套提供注解内容
