package routers

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"

	fmt "fmt"

	bind "github.com/yichouchou/ginPlus/bind"

	utils "github.com/yichouchou/ginPlus/utils"
)

func init() {
	annotation.SetVersion(1773280529)

	parm6829321 := new(string)

	parm4055450 := new(string)

	parm4276713 := new(int)

	parm2134556 := new(bind.ReqTest)

	parm4656420 := new(bind.ReqTest)

	parm2505285 := new(string)

	parm7250372 := new(string)

	parm4131296 := new(int)

	parm7860789 := new(int)

	parm5729715 := new([]bind.ReqTest)

	parm516087 := new([]*bind.ReqTest)

	parm1954649 := new(bind.ReqTest)

	parm607582 := new(string)

	parm9448238 := new(int)

	parm3150762 := new(string)

	parm7122174 := new(int)

	parm5634657 := new(string)

	parm6569822 := new(error)

	parm4625075 := new(bind.ReqTest)

	parm9907647 := new(error)

	parm6637606 := new(int)

	parm2232202 := new(error)

	parm655846 := new(int)

	parm5154511 := new(error)

	parm3059718 := new(int)

	parm9466384 := new(error)

	parm3292563 := new(int)

	parm1360547 := new(error)

	parm1149410 := new(int)

	parm4763257 := new(error)

	parm3310069 := new(bool)

	parm6485532 := new(bool)

	annotation.AddGenOne("main.Hello.Hi1", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi1",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block1",

			Note: "",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "parm1",
					ParmType: reflect.TypeOf(*parm6829321),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm2",
					ParmType: reflect.TypeOf(*parm4055450),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm3",
					ParmType: reflect.TypeOf(*parm4276713),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm2134556),
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
					ParmType: reflect.TypeOf(*parm5634657),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "errHi1",
					ParmType: reflect.TypeOf(*parm6569822),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi2", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi2",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block2",

			Note: "",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm4656420),
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
					ParmType: reflect.TypeOf(*parm4625075),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "errHi2",
					ParmType: reflect.TypeOf(*parm9907647),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi3", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi3",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block3",

			Note: "",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm2505285),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm7250372),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm4131296),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "year",
					ParmType: reflect.TypeOf(*parm7860789),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi3",
					ParmType: reflect.TypeOf(*parm6637606),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi3",
					ParmType: reflect.TypeOf(*parm2232202),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi4", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi4",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block4",

			Note: "",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "List",
					ParmType: reflect.TypeOf(*parm5729715),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm655846),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm5154511),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi5", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi5",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block5",

			Note: "",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
					ParmType: reflect.TypeOf(*parm3059718),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm9466384),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi6", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi6",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block6",

			Note: "",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "List",
					ParmType: reflect.TypeOf(*parm516087),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm3292563),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm1360547),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi7", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi7",
		RouterPath:  "",
		Note:        "",
		Methods:     []string{""},
		Headers:     map[string]string{},
		Consumes:    map[string]string{},
		Produces:    map[string]string{},
		GenComment: &utils.GenComment{

			RouterPath: "/block7",

			Note: "",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "reqList",
					ParmType: reflect.TypeOf(*parm1954649),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm1149410),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm4763257),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser",
		RouterPath:  "/UserRest222",
		Note:        "",
		Methods:     []string{"POST"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"RefererReferer": "www.baidu.com",
		},
		Produces: map[string]string{

			"Accept-Language": "cn",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/LogOutUser2",

			Note: "",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{

				"RefererReferer": "www.baidu.com",
			},

			Produces: map[string]string{

				"Accept-Language": "cn",
			},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm607582),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm9448238),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm3310069),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser",
		RouterPath:  "/UserRest222",
		Note:        "",
		Methods:     []string{"POST"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"RefererReferer": "www.baidu.com",
		},
		Produces: map[string]string{

			"Accept-Language": "cn",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/RegistUser2",

			Note: "",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{

				"RefererReferer": "www.baidu.com",
			},

			Produces: map[string]string{

				"Accept-Language": "cn",
			},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm3150762),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm7122174),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm6485532),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
