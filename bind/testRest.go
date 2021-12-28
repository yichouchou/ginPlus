package bind

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

// @POST /UserRest222
// @resp-custom-user
type UserRest222 struct {
	Tel  int
	Time string
}

// [name string, age int]
// @POST /RegistUser2
func (receiver *UserRest222) RegistUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// [name string, age int]
// @GET /LogOutUser2
func (receiver *UserRest222) LogOutUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return false
}
