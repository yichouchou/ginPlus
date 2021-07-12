package routers

import (
	"reflect"

	"ginPlus/annotation"
	"ginPlus/utils"

	bind "ginPlus/bind"

	examples "ginPlus/examples"

	gin "github.com/gin-gonic/gin"
)

func init() {
	annotation.SetVersion(1626058797)

	abchiValue3 := new(bind.ReqTest)

	abchiValue1 := new(bind.ReqTest)

	abcreqList0 := new([]bind.ReqTest)

	abcstr10 := new(examples.DemoRest)

	abcstr21 := new(examples.DemoRest)

	abcstr32 := new(examples.DemoRest)

	cbacommentHi20 := new(bind.ReqTest)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "hello.hi1",
		Note:       "",
		Methods:    []string{"GET"},
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
				ParmType: reflect.TypeOf(*abchiValue3),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
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
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi2", utils.GenComment{
		RouterPath: "hello.hi2",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "ctx",
				ParmType: reflect.TypeOf(new(gin.Context)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*abchiValue1),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi2",
				ParmType: reflect.TypeOf(*cbacommentHi20),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "errHi2",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi3", utils.GenComment{
		RouterPath: "hello.hi3",
		Note:       "",
		Methods:    []string{"GET"},
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
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi4", utils.GenComment{
		RouterPath: "hello.hi4",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "reqList",
				ParmType: reflect.TypeOf(*abcreqList0),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi5", utils.GenComment{
		RouterPath: "hello.hi5",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "req",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(new(error)),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Example.Say1", utils.GenComment{
		RouterPath: "example.say1",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "str1",
				ParmType: reflect.TypeOf(*abcstr10),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*abcstr21),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*abcstr32),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},
		},
	})
	annotation.AddGenOne("Example.Say2", utils.GenComment{
		RouterPath: "example.say2",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "str1",
				ParmType: reflect.TypeOf(new(examples.DemoRest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(new(examples.DemoRest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(new(examples.DemoRest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},
		},
	})
}
