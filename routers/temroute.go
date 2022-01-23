package routers

import (
	"reflect"

	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"

	bind "github.com/yichouchou/ginPlus/bind"

	utils "github.com/yichouchou/ginPlus/utils"
)

func init() {
	annotation.SetVersion(1642940588)

	parm8498081 := new(string)

	parm9727887 := new(string)

	parm7131847 := new(int)

	parm9984059 := new(bind.ReqTest)

	parm954425 := new(bind.ReqTest)

	parm6203300 := new(string)

	parm6410694 := new(string)

	parm7278511 := new(int)

	parm128162 := new(int)

	parm6933274 := new([]bind.ReqTest)

	parm6340495 := new([]*bind.ReqTest)

	parm2186258 := new(bind.ReqTest)

	parm6138287 := new(string)

	parm3632888 := new(int)

	parm6193015 := new(string)

	parm4895541 := new(int)

	parm1902081 := new(string)

	parm4941318 := new(error)

	parm6122540 := new(bind.ReqTest)

	parm8240456 := new(error)

	parm7455089 := new(int)

	parm3024728 := new(error)

	parm7811211 := new(int)

	parm9431445 := new(error)

	parm8323237 := new(int)

	parm9339106 := new(error)

	parm4965466 := new(int)

	parm5511528 := new(error)

	parm9458047 := new(int)

	parm7979947 := new(error)

	parm8292790 := new(bool)

	parm780408 := new(bool)

	annotation.AddGenOne("main.Hello.Hi1", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi1",
		RouterPath:  "/hello",
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

			RouterPath: "/block1",

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
					ParmType: reflect.TypeOf(*parm8498081),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm9727887),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm7131847),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm9984059),
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
					ParmType: reflect.TypeOf(*parm1902081),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "errHi1",
					ParmType: reflect.TypeOf(*parm4941318),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi2", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi2",
		RouterPath:  "/hello",
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

			RouterPath: "/block2",

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
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm954425),
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
					ParmType: reflect.TypeOf(*parm6122540),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "errHi2",
					ParmType: reflect.TypeOf(*parm8240456),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi3", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi3",
		RouterPath:  "/hello",
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

			RouterPath: "/block3",

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
					ParmType: reflect.TypeOf(*parm6203300),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm6410694),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm7278511),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "year",
					ParmType: reflect.TypeOf(*parm128162),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi3",
					ParmType: reflect.TypeOf(*parm7455089),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi3",
					ParmType: reflect.TypeOf(*parm3024728),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi4", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi4",
		RouterPath:  "/hello",
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

			RouterPath: "/block4",

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
					ParmName: "List",
					ParmType: reflect.TypeOf(*parm6933274),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm7811211),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm9431445),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi5", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi5",
		RouterPath:  "/hello",
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

			RouterPath: "/block5",

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
					ParmName: "reqList",
					ParmType: reflect.TypeOf(new(bind.ReqTest)),
					IsMust:   false,
					ParmKind: reflect.Ptr,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm8323237),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm9339106),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi6", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi6",
		RouterPath:  "/hello",
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

			RouterPath: "/block6",

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
					ParmName: "List",
					ParmType: reflect.TypeOf(*parm6340495),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm4965466),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm5511528),
					IsMust:   false,
					ParmKind: reflect.Interface,
				},
			},
		},
	})
	annotation.AddGenOne("main.Hello.Hi7", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi7",
		RouterPath:  "/hello",
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

			RouterPath: "/block7",

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
					ParmName: "reqList",
					ParmType: reflect.TypeOf(*parm2186258),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm9458047),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm7979947),
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
					ParmType: reflect.TypeOf(*parm6138287),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm3632888),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm8292790),
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
					ParmType: reflect.TypeOf(*parm6193015),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm4895541),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm780408),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
