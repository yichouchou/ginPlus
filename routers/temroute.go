package routers

import (
	"reflect"

	"ginPlus/annotation"
	"ginPlus/utils"
	gin "github.com/gin-gonic/gin"

	_ "ginPlus/bind"

	bind "ginPlus/bind"
)

func init() {
	annotation.SetVersion(1625736705)
	//
	//abc3 := new(bind.ReqTest)
	//
	//abc1 := new(bind.ReqTest)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "hello.hi1",
		Note:       "",
		Methods:    []string{"ANY"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(*bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi1",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "errHi1",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Array,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi2", utils.GenComment{
		RouterPath: "hello.hi2",
		Note:       "",
		Methods:    []string{"ANY"},
		Parms: []*utils.Parm{

			{
				ParmName: "ctx",
				ParmType: reflect.TypeOf(new(*gin.Context)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(*bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi2",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "errHi2",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Array,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi3", utils.GenComment{
		RouterPath: "hello.hi3",
		Note:       "",
		Methods:    []string{"ANY"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "year",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi3",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi3",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Array,
			},
		},
	})
}
