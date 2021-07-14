package main

import (
	fmt "fmt"

	bind "ginPlus/bind"
	"ginPlus/examples"
	"github.com/gin-gonic/gin"
)

//type ReqTest struct {
//	Access_token string `json:"access_token"`
//	UserName     string `json:"user_name" binding:"required"` // With verification mode
//	Password     string `json:"password"`
//	Age          string `json:"age"`
//}

// Hello ...
type Hello struct {
}

// Hello Annotated route (bese on beego way)
// @Router /block [post]
//func (s *Hello) Hello(c *gin.Context, req *ReqTest) {
//	fmt.Println(req)
//	c.JSON(http.StatusOK, "ok")
//}

//// Hello2 Route without annotation (the parameter is 2 default post)
//func (s *Hello) Hello2(c *gin.Context, req ReqTest) {
//	fmt.Println(req)
//	c.JSON(http.StatusOK, "ok")
//}
//

// [name string, password string, age int]
// @GET /block1
func (s *Hello) Hi1(parm1 string, parm2 string, parm3 int, hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi1 string, errHi1 error) {
	fmt.Println(parm1 + parm2)
	fmt.Println(parm3)
	fmt.Println(hi)
	fmt.Println(hiValue)
	return "ni hao", nil
}

// @GET /block2
func (s *Hello) Hi2(ctx *gin.Context, hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi2 bind.ReqTest, errHi2 error) {
	fmt.Println(hi)
	fmt.Println(hiValue)
	return bind.ReqTest{}, nil
}

// [name, password string, age, year int]
// @GET /block3
func (s *Hello) Hi3(name, password string, age, year int) (commentHi3 int, errHi3 error) {

	return 10, nil
}

// Hello Annotated route (bese on beego way)
// {List []bind.ReqTest}
// @POST /block4
func (s *Hello) Hi4(reqList []bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("这是post请求接受结构体数组")

	for i := range reqList {
		fmt.Println(reqList[i])
	}
	return 4, nil
}

// Hello Annotated route (bese on beego way)
// [reqList *bind.ReqTest]
// @GET /block5
func (s *Hello) Hi5(reqList *bind.ReqTest) (index int, errHi5 error) {
	fmt.Println(reqList)
	fmt.Println("这是get请求接受结构体指针")
	return 5, nil
}

//
//func init() {
//	annotation.SetVersion(1625327764)
//	annotation.AddGenOne("Hello.Hello", "/block", []string{"post"})
//	_= map[string]map[string]string{}
//}
//
////annotation.AddGenOne("Hello.Hello", "/block", []string{"post"})
//
//// []string{"post"} 这样的描述太过于简洁，我希望能够把入参的名称和类型都注册上去，把返回值类型也注册上去-这个不是绝对的
////													入参1类型：string  入参名字:name
////													入参2类型：string  入参名字:password
////													入参3类型：string  入参名字:age
////												          * * *
////													以什么协议传输：josn -默认
////大致的类型： _= map[string]map[string]string{}   post:
//
////然后在请求进入的时候，根据这个map 获取到相关的参数，

// Example ...
type Example struct {
}

// Hello Annotated route (bese on beego way)
// [str1, str2, str3 examples.DemoRest]
// @GET /Say1
func (example *Example) Say1(str1, str2, str3 examples.DemoRest) (str4 string) {
	fmt.Println(str1, str2, str3)
	return "这个是example say1方法rest"
}

// Hello Annotated route (bese on beego way)
// [str1, str2, str3 string]
// {rest examples.DemoRest}
// @GET /Say2
func (example *Example) Say2(str1, str2, str3 string, rest examples.DemoRest) (str4 string) {
	fmt.Println(&str1, &str2, &str3)
	fmt.Println(rest)
	return "这个是example say2方法rest"
}
