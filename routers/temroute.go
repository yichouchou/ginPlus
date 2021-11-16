package routers

import (
	"reflect"

	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"

	bind "github.com/yichouchou/ginPlus/bind"

	examples "github.com/yichouchou/ginPlus/examples"
)

func init() {
	annotation.SetVersion(1637041512)

	GIQZEUFCIE := new(string)

	OYNCLYUMYU := new(string)

	VCWKRUHZDF := new(int)

	ZTWJVDKFZH := new(bind.ReqTest)

	OCVGVJEGEP := new(bind.ReqTest)

	ZQXXXUQPQA := new(string)

	XCCHRDDDRH := new(string)

	ATMEOYVPDY := new(int)

	RIVXUIHUXV := new(int)

	XUUKLTJKWO := new([]bind.ReqTest)

	NMDGJTDMQU := new([]*bind.ReqTest)

	PNRTSJKREE := new(bind.ReqTest)

	LXAXBSAZOG := new(examples.DemoRest)

	KXHDACOLKO := new(examples.DemoRest)

	GKHHYWXFTB := new(examples.DemoRest)

	PSHQNGITUQ := new(string)

	HPBBVNTHBA := new(string)

	WWMKNCTHPL := new(string)

	BCDAAWEYHC := new(examples.DemoRest)

	IVDZRHPPSG := new(string)

	CKAZVOJBOF := new(int)

	VNKEPGOZVU := new(string)

	GQKHWFNALN := new(int)

	ZPCPTUJVEX := new(string)

	VBJCQTBLPL := new(error)

	XABVLBKBEU := new(bind.ReqTest)

	GHYLFXXZGM := new(error)

	CMWJDXPXFW := new(int)

	OPAULNVJUN := new(error)

	CDGTMQNSMD := new(int)

	TFIOGOAHCF := new(error)

	FVHSSNLPAJ := new(int)

	NIVJJVTNZX := new(error)

	ICHLWOXAEM := new(int)

	YYZMLOFEPP := new(error)

	JRKWLYSTHG := new(int)

	SNZDZZTRZO := new(error)

	FMIPQGCTGO := new(string)

	MBJUHWEMLD := new(string)

	WPRKKQQSLD := new(bool)

	HMNNZENSOB := new(bool)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "/block1",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*GIQZEUFCIE),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*OYNCLYUMYU),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*VCWKRUHZDF),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*ZTWJVDKFZH),
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
				ParmType: reflect.TypeOf(*ZPCPTUJVEX),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "errHi1",
				ParmType: reflect.TypeOf(*VBJCQTBLPL),
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
				ParmType: reflect.TypeOf(*OCVGVJEGEP),
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
				ParmType: reflect.TypeOf(*XABVLBKBEU),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "errHi2",
				ParmType: reflect.TypeOf(*GHYLFXXZGM),
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
				ParmType: reflect.TypeOf(*ZQXXXUQPQA),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*XCCHRDDDRH),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*ATMEOYVPDY),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "year",
				ParmType: reflect.TypeOf(*RIVXUIHUXV),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi3",
				ParmType: reflect.TypeOf(*CMWJDXPXFW),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi3",
				ParmType: reflect.TypeOf(*OPAULNVJUN),
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
				ParmType: reflect.TypeOf(*XUUKLTJKWO),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*CDGTMQNSMD),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*TFIOGOAHCF),
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
				ParmType: reflect.TypeOf(*FVHSSNLPAJ),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*NIVJJVTNZX),
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
				ParmType: reflect.TypeOf(*NMDGJTDMQU),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*ICHLWOXAEM),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*YYZMLOFEPP),
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
				ParmType: reflect.TypeOf(*PNRTSJKREE),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*JRKWLYSTHG),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*SNZDZZTRZO),
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
				ParmType: reflect.TypeOf(*LXAXBSAZOG),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*KXHDACOLKO),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*GKHHYWXFTB),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(*FMIPQGCTGO),
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
				ParmType: reflect.TypeOf(*PSHQNGITUQ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*HPBBVNTHBA),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*WWMKNCTHPL),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "rest",
				ParmType: reflect.TypeOf(*BCDAAWEYHC),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(*MBJUHWEMLD),
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
				ParmType: reflect.TypeOf(*IVDZRHPPSG),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*CKAZVOJBOF),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*WPRKKQQSLD),
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
				ParmType: reflect.TypeOf(*VNKEPGOZVU),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*GQKHWFNALN),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*HMNNZENSOB),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
