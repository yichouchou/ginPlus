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
	annotation.SetVersion(1773301654)

	parm4911612 := new(string)

	parm500177 := new(string)

	parm6603643 := new(int)

	parm3788392 := new(bind.ReqTest)

	parm4842980 := new(bind.ReqTest)

	parm2110133 := new(string)

	parm6742842 := new(string)

	parm5791677 := new(int)

	parm3411053 := new(int)

	parm6236563 := new([]bind.ReqTest)

	parm7876385 := new([]*bind.ReqTest)

	parm8695614 := new(bind.ReqTest)

	parm4584738 := new(string)

	parm8552895 := new(int)

	parm1351294 := new(string)

	parm8729395 := new(int)

	parm7181196 := new(string)

	parm3689764 := new(error)

	parm5785209 := new(bind.ReqTest)

	parm1268652 := new(error)

	parm6706049 := new(int)

	parm1206680 := new(error)

	parm315151 := new(int)

	parm5620549 := new(error)

	parm5399271 := new(int)

	parm9452572 := new(error)

	parm1209667 := new(int)

	parm4567171 := new(error)

	parm8687042 := new(int)

	parm375660 := new(error)

	parm1493700 := new(bool)

	parm1156520 := new(bool)

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
					ParmType: reflect.TypeOf(*parm4911612),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm2",
					ParmType: reflect.TypeOf(*parm500177),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "parm3",
					ParmType: reflect.TypeOf(*parm6603643),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "hiValue",
					ParmType: reflect.TypeOf(*parm3788392),
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
					ParmType: reflect.TypeOf(*parm7181196),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "errHi1",
					ParmType: reflect.TypeOf(*parm3689764),
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
					ParmType: reflect.TypeOf(*parm4842980),
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
					ParmType: reflect.TypeOf(*parm5785209),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},

				{
					ParmName: "errHi2",
					ParmType: reflect.TypeOf(*parm1268652),
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
					ParmType: reflect.TypeOf(*parm2110133),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "password",
					ParmType: reflect.TypeOf(*parm6742842),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm5791677),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "year",
					ParmType: reflect.TypeOf(*parm3411053),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "commentHi3",
					ParmType: reflect.TypeOf(*parm6706049),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi3",
					ParmType: reflect.TypeOf(*parm1206680),
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
					ParmType: reflect.TypeOf(*parm6236563),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm315151),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm5620549),
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
					ParmType: reflect.TypeOf(*parm5399271),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm9452572),
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
					ParmType: reflect.TypeOf(*parm7876385),
					IsMust:   false,
					ParmKind: reflect.Slice,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm1209667),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi4",
					ParmType: reflect.TypeOf(*parm4567171),
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
					ParmType: reflect.TypeOf(*parm8695614),
					IsMust:   false,
					ParmKind: reflect.Struct,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "index",
					ParmType: reflect.TypeOf(*parm8687042),
					IsMust:   false,
					ParmKind: reflect.Int,
				},

				{
					ParmName: "errHi5",
					ParmType: reflect.TypeOf(*parm375660),
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
					ParmType: reflect.TypeOf(*parm4584738),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8552895),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm1493700),
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
					ParmType: reflect.TypeOf(*parm1351294),
					IsMust:   false,
					ParmKind: reflect.String,
				},

				{
					ParmName: "age",
					ParmType: reflect.TypeOf(*parm8729395),
					IsMust:   false,
					ParmKind: reflect.Int,
				},
			},

			Result: []*utils.Parm{

				{
					ParmName: "success",
					ParmType: reflect.TypeOf(*parm1156520),
					IsMust:   false,
					ParmKind: reflect.Bool,
				},
			},
		},
	})
}
