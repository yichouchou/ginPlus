package routers

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"

	bind "github.com/yichouchou/ginPlus/bind"

	examples "github.com/yichouchou/ginPlus/examples"
)

func init() {
	annotation.SetVersion(1637029395)

	WPDYNTTCES := new(string)

	NUHWDXWJGD := new(string)

	RRZFBGMCEC := new(int)

	GCATEBNRRR := new(bind.ReqTest)

	DNUWCDJTKN := new(bind.ReqTest)

	AUUNUWSNET := new(string)

	EURUTGREUE := new(string)

	PAGYYPVJPR := new(int)

	VRRPSZFCXG := new(int)

	NACKRJLJJE := new([]bind.ReqTest)

	JXSBDAIXII := new([]*bind.ReqTest)

	ZPOEZZPZVM := new(bind.ReqTest)

	EPCMIUPYPR := new(examples.DemoRest)

	GPTZFSWXPD := new(examples.DemoRest)

	JUNNOBPBQU := new(examples.DemoRest)

	ETEQPSTVKZ := new(string)

	NETPLEVWCI := new(string)

	UOVVKWOPYN := new(string)

	FGJBPYMXGZ := new(examples.DemoRest)

	DNGIERTNOB := new(string)

	ULJYQZMPML := new(int)

	CZMUKUNBSB := new(string)

	UVRGBNUKYX := new(int)

	CHGQPGEBOV := new(string)

	MJFPZMSGPJ := new(error)

	GKQOSLVWSL := new(bind.ReqTest)

	LQXDROGJGS := new(error)

	HVYUJSWESK := new(int)

	UHEDEMHVAF := new(error)

	JMVQKERTJC := new(int)

	RQHBYDNZTJ := new(error)

	NBFGBGGZKL := new(int)

	BNFOACDMEL := new(error)

	LJBYZPNMJY := new(int)

	XIZNCALPFP := new(error)

	KGSJPAKKSS := new(int)

	THHQKIBTEU := new(error)

	NZNGKBDOBQ := new(string)

	NYQFIMQGRR := new(string)

	QTSZNLZATK := new(bool)

	OUUWZEQTJW := new(bool)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "/block1",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*WPDYNTTCES),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*NUHWDXWJGD),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*RRZFBGMCEC),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*GCATEBNRRR),
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
				ParmType: reflect.TypeOf(*CHGQPGEBOV),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "errHi1",
				ParmType: reflect.TypeOf(*MJFPZMSGPJ),
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
				ParmType: reflect.TypeOf(*DNUWCDJTKN),
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
				ParmType: reflect.TypeOf(*GKQOSLVWSL),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "errHi2",
				ParmType: reflect.TypeOf(*LQXDROGJGS),
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
				ParmType: reflect.TypeOf(*AUUNUWSNET),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*EURUTGREUE),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*PAGYYPVJPR),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "year",
				ParmType: reflect.TypeOf(*VRRPSZFCXG),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi3",
				ParmType: reflect.TypeOf(*HVYUJSWESK),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi3",
				ParmType: reflect.TypeOf(*UHEDEMHVAF),
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
				ParmType: reflect.TypeOf(*NACKRJLJJE),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*JMVQKERTJC),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*RQHBYDNZTJ),
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
				ParmType: reflect.TypeOf(*NBFGBGGZKL),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*BNFOACDMEL),
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
				ParmType: reflect.TypeOf(*JXSBDAIXII),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*LJBYZPNMJY),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*XIZNCALPFP),
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
				ParmType: reflect.TypeOf(*ZPOEZZPZVM),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*KGSJPAKKSS),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*THHQKIBTEU),
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
				ParmType: reflect.TypeOf(*EPCMIUPYPR),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*GPTZFSWXPD),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*JUNNOBPBQU),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(*NZNGKBDOBQ),
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
				ParmType: reflect.TypeOf(*ETEQPSTVKZ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str2",
				ParmType: reflect.TypeOf(*NETPLEVWCI),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "str3",
				ParmType: reflect.TypeOf(*UOVVKWOPYN),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "rest",
				ParmType: reflect.TypeOf(*FGJBPYMXGZ),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "str4",
				ParmType: reflect.TypeOf(*NYQFIMQGRR),
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
				ParmType: reflect.TypeOf(*DNGIERTNOB),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*ULJYQZMPML),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*QTSZNLZATK),
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
				ParmType: reflect.TypeOf(*CZMUKUNBSB),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*UVRGBNUKYX),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*OUUWZEQTJW),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
