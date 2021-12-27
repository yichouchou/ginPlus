package routers

import (
	"reflect"

	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"

	examples "github.com/yichouchou/ginPlus/examples"

	bind "github.com/yichouchou/ginPlus/bind"
)

func init() {
	annotation.SetVersion(1639924621)

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

	parm6138287 := new(examples.DemoRest)

	parm3632888 := new(examples.DemoRest)

	parm8292790 := new(examples.DemoRest)

	parm4895541 := new(string)

	parm780408 := new(string)

	parm7387 := new(string)

	parm4066831 := new(examples.DemoRest)

	parm625356 := new(string)

	parm7341737 := new(int)

	parm6111485 := new(string)

	parm7515026 := new(int)

	parm2003090 := new(string)

	parm8565194 := new(int)

	parm8712433 := new(string)

	parm9424147 := new(int)

	parm5764324 := new(string)

	parm3516159 := new(int)

	parm951957 := new(string)

	parm8043721 := new(int)

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

	parm6193015 := new(string)

	parm4315429 := new(string)

	parm6960631 := new(bool)

	parm2086413 := new(bool)

	parm8090563 := new(bool)

	parm4474078 := new(bool)

	parm6971353 := new(bool)

	parm627189 := new(bool)

	annotation.AddGenOne("main.Hello.Hi1", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi1",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block1",

			Note: "/block1",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block2",

			Note: "/block2",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block3",

			Note: "/block3",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block4",

			Note: "/block4",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block5",

			Note: "/block5",

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block6",

			Note: "/block6",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/block7",

			Note: "/block7",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

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
	annotation.AddGenOne("main.Example.Say1", utils.GenRouterInfo{
		HandFunName: "main.Example.Say1",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/Say1",

			Note: "/Say1",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "str1",
					ParmType: reflect.TypeOf(*parm6138287),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "str2",
					ParmType: reflect.TypeOf(*parm3632888),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "str3",
					ParmType: reflect.TypeOf(*parm8292790),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "str4",
					ParmType: reflect.TypeOf(*parm6193015),
					IsMust:   false,
					ParmKind: reflect.String,
				},
			},
		},
	})
	annotation.AddGenOne("main.Example.Say2", utils.GenRouterInfo{
		HandFunName: "main.Example.Say2",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/Say2",

			Note: "/Say2",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "str1",
					ParmType: reflect.TypeOf(*parm4895541),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "str2",
					ParmType: reflect.TypeOf(*parm780408),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "str3",
					ParmType: reflect.TypeOf(*parm7387),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "rest",
					ParmType: reflect.TypeOf(*parm4066831),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "str4",
					ParmType: reflect.TypeOf(*parm4315429),
					IsMust:   false,
					ParmKind: reflect.String,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/examples.UserRest.LogOutUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/examples.UserRest.LogOutUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/LogOutUser3",

			Note: "/LogOutUser3",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm625356),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm7341737),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm6960631),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/examples.UserRest.RegistUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/examples.UserRest.RegistUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/RegistUser3",

			Note: "/RegistUser3",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm6111485),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm7515026),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm2086413),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/examples/simpleExample/controllor.UserRest222.LogOutUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/examples/simpleExample/controllor.UserRest222.LogOutUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/LogOutUser4",

			Note: "/LogOutUser4",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm2003090),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8565194),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm8090563),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/examples/simpleExample/controllor.UserRest222.RegistUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/examples/simpleExample/controllor.UserRest222.RegistUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/RegistUser4",

			Note: "/RegistUser4",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm8712433),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm9424147),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm4474078),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/LogOutUser2",

			Note: "/LogOutUser2",

			Methods: []string{"GET"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm5764324),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm3516159),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm6971353),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
	annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser",
		RouterPath:  "/hi",
		Note:        "",
		Methods:     []string{"get", "post"},
		Headers: map[string]string{

			"Content-Type": "application/json",
		},
		Consumes: map[string]string{

			"Content-Type": "application/Consumes",
		},
		Produces: map[string]string{

			"Content-Type": "application/Produces",
		},
		GenComment: &utils.GenComment{

			RouterPath: "/RegistUser2",

			Note: "/RegistUser2",

			Methods: []string{"POST"},

			Headers: map[string]string{},

			Consumes: map[string]string{},

			Produces: map[string]string{},

			Parms: []*utils.Parm{

				{
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm951957),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8043721),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm627189),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
