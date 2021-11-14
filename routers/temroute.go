package routers

import (
	"ginPlus/annotation"
	"ginPlus/utils"
	"reflect"

	bind "ginPlus/bind"

	examples "ginPlus/examples"
)

func init() {
	annotation.SetVersion(1636805177)

	abchiValuestruct3 := new(bind.ReqTest)

	abchiValuestruct0 := new(bind.ReqTest)

	abcListslice0 := new([]bind.ReqTest)

	abcListslice02 := new([]*bind.ReqTest)

	abcreqListstruct0 := new(bind.ReqTest)

	abcstr1struct0 := new(examples.DemoRest)

	abcstr2struct1 := new(examples.DemoRest)

	abcstr3struct2 := new(examples.DemoRest)

	abcreststruct3 := new(examples.DemoRest)

	cbacommentHi2struct0 := new(bind.ReqTest)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "/block1",
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
				ParmType: reflect.TypeOf(*abchiValuestruct3),
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
		RouterPath: "/block2",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*abchiValuestruct0),
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
				ParmType: reflect.TypeOf(*cbacommentHi2struct0),
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
		RouterPath: "/block3",
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
		RouterPath: "/block4",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "List",
				ParmType: reflect.TypeOf(*abcListslice0),
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
		RouterPath: "/block5",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "reqList",
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
	annotation.AddGenOne("Hello.Hi6", utils.GenComment{
		RouterPath: "/block6",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "List",
				ParmType: reflect.TypeOf(*abcListslice02),
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
	annotation.AddGenOne("Hello.Hi7", utils.GenComment{
		RouterPath: "/block7",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "reqList",
				ParmType: reflect.TypeOf(*abcreqListstruct0),
				IsMust:   false,
				ParmKind: reflect.Struct,
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
		RouterPath: "/Say1",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "str1",
				ParmType: reflect.TypeOf(*abcstr1struct0),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*abcstr2struct1),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*abcstr3struct2),
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
		RouterPath: "/Say2",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "str1",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "rest",
				ParmType: reflect.TypeOf(*abcreststruct3),
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
	annotation.AddGenOne("UserRest.LogOutUser", utils.GenComment{
		RouterPath: "/LogOutUser",
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
				ParmName: "age",
				ParmType: reflect.TypeOf(new(int)),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(new(bool)),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
	annotation.AddGenOne("UserRest.RegistUser", utils.GenComment{
		RouterPath: "/RegistUser",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
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
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(new(bool)),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
