package examples

import "fmt"

type DemoRest struct {
	Tel  int
	Time string
}

// TestRest TestRest
type TestRest struct {
	Tel  int
	Time string
}

// test UserRest
type UserRest struct {
	Tel  int
	Time string
}

// [name string, age int]
// @POST /RegistUser3
func (receiver *UserRest) RegistUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// [name string, age int]
// @GET /LogOutUser3
func (receiver *UserRest) LogOutUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return false
}
