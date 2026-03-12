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
	annotation.SetVersion(1773308689)

	parm3821835 := new(string)

	parm4316971 := new(string)

	parm9398562 := new(int)

	parm6403331 := new(bind.ReqTest)

	parm2882520 := new(bind.ReqTest)

	parm6761615 := new(string)

	parm667983 := new(string)

	parm8779942 := new(int)

	parm5633910 := new(int)

	parm4082962 := new([]bind.ReqTest)

	parm2000471 := new([]*bind.ReqTest)

	parm583728 := new(bind.ReqTest)

	parm196951 := new(string)

	parm3683134 := new(int)

	parm676844 := new(string)

	parm9577346 := new(int)

	parm5418717 := new(string)

	parm7025467 := new(error)

	parm7512277 := new(bind.ReqTest)

	parm115842 := new(error)

	parm9917814 := new(int)

	parm9096512 := new(error)

	parm1447909 := new(int)

	parm4635332 := new(error)

	parm5895782 := new(int)

	parm3384230 := new(error)

	parm632881 := new(int)

	parm7583395 := new(error)

	parm25649 := new(int)

	parm9686977 := new(error)

	parm9436330 := new(bool)

	parm1796724 := new(bool)

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
					ParmType: reflect.TypeOf(*parm3821835),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm2",
					ParmType: reflect.TypeOf(*parm4316971),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm3",
					ParmType: reflect.TypeOf(*parm9398562),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm6403331),
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
					ParmType: reflect.TypeOf(*parm5418717),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "errHi1",
					ParmType: reflect.TypeOf(*parm7025467),
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
					ParmType: reflect.TypeOf(*parm2882520),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "响应内容:",
					ParmType: reflect.TypeOf(new(bind.ReqTest)),
					IsMust:   false,
					ParmKind: reflect.Ptr,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi2",
					ParmType: reflect.TypeOf(*parm7512277),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "errHi2",
					ParmType: reflect.TypeOf(*parm115842),
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
					ParmType: reflect.TypeOf(*parm6761615),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm667983),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8779942),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "year",
					ParmType: reflect.TypeOf(*parm5633910),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi3",
					ParmType: reflect.TypeOf(*parm9917814),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi3",
					ParmType: reflect.TypeOf(*parm9096512),
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
					ParmType: reflect.TypeOf(*parm4082962),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm1447909),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm4635332),
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
					ParmType: reflect.TypeOf(*parm5895782),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm3384230),
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
					ParmType: reflect.TypeOf(*parm2000471),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm632881),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm7583395),
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
					ParmType: reflect.TypeOf(*parm583728),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm25649),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm9686977),
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

			"aaaaaa": "bbbbb",
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
					ParmType: reflect.TypeOf(*parm196951),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm3683134),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm9436330),
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

			"aaaaaa": "bbbbb",
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
					ParmType: reflect.TypeOf(*parm676844),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm9577346),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm1796724),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
