package routers

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"

	utils "github.com/yichouchou/ginPlus/utils"

	fmt "fmt"

	bind "github.com/yichouchou/ginPlus/bind"
)

func init() {
	annotation.SetVersion(1773208422)

	parm8637247 := new(string)

	parm3330902 := new(string)

	parm8566705 := new(int)

	parm1919552 := new(bind.ReqTest)

	parm2592152 := new(bind.ReqTest)

	parm2496965 := new(string)

	parm3943993 := new(string)

	parm1477661 := new(int)

	parm6307446 := new(int)

	parm5675470 := new([]bind.ReqTest)

	parm8797882 := new([]*bind.ReqTest)

	parm3660262 := new(bind.ReqTest)

	parm4760904 := new(string)

	parm5584411 := new(int)

	parm2898513 := new(string)

	parm4180025 := new(int)

	parm690207 := new(string)

	parm2643503 := new(error)

	parm411449 := new(bind.ReqTest)

	parm6240293 := new(error)

	parm3370934 := new(int)

	parm1904981 := new(error)

	parm2434926 := new(int)

	parm9081944 := new(error)

	parm3621767 := new(int)

	parm713713 := new(error)

	parm3672799 := new(int)

	parm5583844 := new(error)

	parm3316892 := new(int)

	parm4392087 := new(error)

	parm9576979 := new(bool)

	parm2296521 := new(bool)

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
					ParmName: "name",
					ParmType: reflect.TypeOf(*parm8637247),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm3330902),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8566705),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm1919552),
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
					ParmType: reflect.TypeOf(*parm690207),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "errHi1",
					ParmType: reflect.TypeOf(*parm2643503),
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
					ParmType: reflect.TypeOf(*parm2592152),
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
					ParmType: reflect.TypeOf(*parm411449),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "errHi2",
					ParmType: reflect.TypeOf(*parm6240293),
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
					ParmType: reflect.TypeOf(*parm2496965),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm3943993),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm1477661),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "year",
					ParmType: reflect.TypeOf(*parm6307446),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi3",
					ParmType: reflect.TypeOf(*parm3370934),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi3",
					ParmType: reflect.TypeOf(*parm1904981),
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
					ParmType: reflect.TypeOf(*parm5675470),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm2434926),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm9081944),
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
					ParmType: reflect.TypeOf(*parm3621767),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm713713),
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
					ParmType: reflect.TypeOf(*parm8797882),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm3672799),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm5583844),
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
					ParmType: reflect.TypeOf(*parm3660262),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm3316892),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm4392087),
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
					ParmType: reflect.TypeOf(*parm4760904),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm5584411),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm9576979),
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
					ParmType: reflect.TypeOf(*parm2898513),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm4180025),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm2296521),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
