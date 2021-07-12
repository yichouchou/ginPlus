package examples

import "fmt"

type DemoRest struct {
	Tel  int
	Time string
}

type UserRest struct {
}

// Hello Annotated route (bese on beego way)
// [str1, str2, str3 examples.DemoRest] [commentHi1 string,errHi1 error]
// @POST /RegistUser
func (receiver *UserRest) RegistUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// Hello Annotated route (bese on beego way)
// [str1, str2, str3 examples.DemoRest] [commentHi1 string,errHi1 error]
// @GET /LogOutUser
func (receiver *UserRest) LogOutUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return false
}
