package controllor

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

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
type UserRest222 struct {
	Tel  int
	Time string
}

// [name string, age int]
// @POST /RegistUser4
func (receiver *UserRest222) RegistUser(ctx gin.Context, name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// [name string, age int]
// @GET /LogOutUser4
func (receiver *UserRest222) LogOutUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return false
}
