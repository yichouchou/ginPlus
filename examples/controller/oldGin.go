package controller

import "fmt"

type DemoRest struct {
	Tel  int
	Time string
}

type UserRest struct {
	Tel  int
	Time string
}

// [name string, age int]
// @POST /RegistUser
func (receiver *UserRest) RegistUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// @GET /LogOutUser
func (receiver *UserRest) LogOutUser(name string, age int, rest DemoRest) (success bool) {
	fmt.Println(name, age, "-----user")
	fmt.Println(rest.Tel, "-----Tel")
	fmt.Println(rest.Time, "-----Time")
	return false
}
