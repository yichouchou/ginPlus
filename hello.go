package main

import (
	fmt "fmt"

	bind "github.com/yichouchou/ginPlus/bind"
	"github.com/yichouchou/ginPlus/examples"
	"github.com/yichouchou/ginPlus/utils"
)

// @POST /hello
// @resp-custom-user
type Hello struct {
	ReqContentType    utils.ReqHeaderInfo `head:"application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8"`
	ReqUserAgent      utils.ReqHeaderInfo `head:"PostmanRuntime/7.26.8"`
	ReqAcceptEncoding utils.ReqHeaderInfo `head:"gzip, deflate, br"`

	RespContentType    utils.RespHeaderInfo `head:"application/json;charset=UTF-8,text/html;charset=UTF-8,multipart/form-data;charset=UTF-8"`
	RespUserAgent      utils.RespHeaderInfo `head:"PostmanRuntime/7.26.8"`
	RespAcceptEncoding utils.RespHeaderInfo `head:"gzip, deflate, br"`
}

// [name string, password string, age int]
// @GET /block1
// @resp-custom-user
func (s *Hello) Hi1(parm1 string, parm2 string, parm3 int, hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi1 string, errHi1 error) {
	fmt.Println(parm1 + parm2)
	fmt.Println(parm3)
	fmt.Println(hi)
	fmt.Println(hiValue)
	return "ni hao", nil
}

// @POST /block2
// @resp-custom-user
func (s *Hello) Hi2(hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi2 bind.ReqTest, errHi2 error) {
	fmt.Println(hi)
	fmt.Println(hiValue)
	return bind.ReqTest{}, nil
}

// @GET /block3
// @resp-custom-user
func (s *Hello) Hi3(name, password string, age, year int) (commentHi3 int, errHi3 error) {
	fmt.Println("---get请求，无参数注解，多基本数据类型已经调通")
	return 10, nil
}

// {List []bind.ReqTest}
// @POST /block4
// @resp-custom-user
func (s *Hello) Hi4(reqList []bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("这是post请求接受结构体数组")
	fmt.Println("---post请求，有参数注解，结构体值数组已经调通")

	for i := range reqList {
		fmt.Println(reqList[i])
	}
	return 4, nil
}

// [reqList *bind.ReqTest]
// @GET /block5
func (s *Hello) Hi5(reqList *bind.ReqTest) (index int, errHi5 error) {
	fmt.Println(reqList)
	fmt.Println("这是get请求接受结构体指针")
	fmt.Println("---这是get请求接受结构体指针已经调通")
	return 5, nil
}

// {List []*bind.ReqTest}
// @POST /block6
func (s *Hello) Hi6(reqList []*bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("这是post请求接受指针结构体数组")
	fmt.Println("---post请求，有参数注解，结构体指针数组已经调通")
	for i := range reqList {
		fmt.Println(reqList[i])
	}
	return 6, nil
}

// [reqList bind.ReqTest]
// @GET /block7
func (s *Hello) Hi7(reqList bind.ReqTest) (index int, errHi5 error) {
	fmt.Println(reqList)
	fmt.Println("这是get请求接受结构体指针")
	fmt.Println("---这是get请求接受结构体指针已经调通")
	return 5, nil
}

// @GET /Example
// @resp-custom-user
type Example struct {
}

// [str1, str2, str3 examples.DemoRest]
// @GET /Say1
func (example *Example) Say1(str1, str2, str3 examples.DemoRest) (str4 string) {
	fmt.Println("-- Say1 rest接口--")
	return "这个是example say1方法rest"
}

// {str1, str2, str3 string,rest examples.DemoRest}
// @POST /Say2
func (example *Example) Say2(str1, str2, str3 string, rest examples.DemoRest) (str4 string) {
	fmt.Println("--post请求，既包含请求体的内容，也包含请求头的内容，参数绑定成功---")
	fmt.Println(str1, str2, str3)
	fmt.Println(rest)
	return "这个是example say2方法rest"
}

type DemoRest struct {
	Tel  int
	Time string
}

// TestRest TestRest
type TestRest struct {
	Tel  int
	Time string
}
