package main

import (
	fmt "fmt"

	bind "github.com/yichouchou/ginPlus/bind"
	"github.com/yichouchou/ginPlus/utils"
)

// @POST /hello
// @resp-custom-user

type BaseRestSetting struct {
	CommonHeader string
}

type Hello struct {
	BaseRestSetting
	ReqContentType    utils.ReqHeaderInfo `head:"application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8"`
	ReqUserAgent      utils.ReqHeaderInfo `head:"PostmanRuntime/7.26.8"`
	ReqAcceptEncoding utils.ReqHeaderInfo `head:"gzip, deflate, br"`

	RespContentType    utils.RespHeaderInfo `head:"application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8"`
	RespUserAgent      utils.RespHeaderInfo `head:"PostmanRuntime/7.26.8"`
	RespAcceptEncoding utils.RespHeaderInfo `head:"gzip, deflate, br"`
}

// [parm1 string, parm2 string, parm3 int]
// @GET /block1
// @resp-custom-user
// 请求参数: parm1, parm2 (string), parm3 (int), hiValue, hi (struct)
// 响应内容: "ni hao"
func (s *Hello) Hi1(parm1 string, parm2 string, parm3 int, hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi1 string, errHi1 error) {
	fmt.Println("=== Block1 GET ===")
	fmt.Println("基本参数:", parm1, parm2, parm3)
	fmt.Println("结构体值:", hiValue)
	fmt.Println("结构体指针:", hi)
	return "ni hao", nil
}

// @POST /block2
// @resp-custom-user
// 请求参数: hiValue, hi (struct)
// 响应内容: bind.ReqTest{}
func (s *Hello) Hi2(hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi2 bind.ReqTest, errHi2 error) {
	fmt.Println("=== Block2 POST ===")
	fmt.Println("结构体值:", hiValue)
	fmt.Println("结构体指针:", hi)
	return bind.ReqTest{}, nil
}

// @GET /block3
// @resp-custom-user
// 请求参数: name, password (string), age, year (int)
// 响应内容: 10
func (s *Hello) Hi3(name, password string, age, year int) (commentHi3 int, errHi3 error) {
	fmt.Println("=== Block3 GET ===")
	fmt.Println("多基本类型参数:", name, password, age, year)
	return 10, nil
}

// {List []bind.ReqTest}
// @POST /block4
// @resp-custom-user
// 请求参数: reqList ([]bind.ReqTest) 结构体数组
// 响应内容: 4
func (s *Hello) Hi4(reqList []bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("=== Block4 POST ===")
	fmt.Println("结构体数组:", reqList)
	return 4, nil
}

// [reqList *bind.ReqTest]
// @GET /block5
// 请求参数: reqList (*bind.ReqTest) 结构体指针
// 响应内容: 5
func (s *Hello) Hi5(reqList *bind.ReqTest) (index int, errHi5 error) {
	fmt.Println("=== Block5 GET ===")
	fmt.Println("结构体指针:", reqList)
	return 5, nil
}

// {List []*bind.ReqTest}
// @POST /block6
// 请求参数: reqList ([]*bind.ReqTest) 结构体指针数组
// 响应内容: 6
func (s *Hello) Hi6(reqList []*bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("=== Block6 POST ===")
	fmt.Println("结构体指针数组:", reqList)
	return 6, nil
}

// [reqList bind.ReqTest]
// @GET /block7
// 请求参数: reqList (bind.ReqTest) 结构体值
// 响应内容: 5
func (s *Hello) Hi7(reqList bind.ReqTest) (index int, errHi5 error) {
	fmt.Println("=== Block7 GET ===")
	fmt.Println("结构体值:", reqList)
	return 5, nil
}

// @GET /Example
// @resp-custom-user
type Example struct {
}
