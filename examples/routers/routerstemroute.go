package routers

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/examples/controller"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"
)

func init() {
	annotation.SetVersion(1772610818)

	XHWSYFLPUW7535484998337571078 := new(string)

	XHWSYFLPUW1837728370458510504 := new(int)

	XHWSYFLPUW2145367275043906616 := new(controller.DemoRest)

	XHWSYFLPUW1382764017443833647 := new(string)

	XHWSYFLPUW6303292334592687199 := new(int)

	XHWSYFLPUW4969082095760678017 := new(bool)

	XHWSYFLPUW5093615082521380362 := new(bool)

	annotation.AddGenOne("UserRest.LogOutUser", utils.GenComment{
		RouterPath: "/LogOutUser",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*XHWSYFLPUW7535484998337571078),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*XHWSYFLPUW1837728370458510504),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "rest",
				ParmType: reflect.TypeOf(*XHWSYFLPUW2145367275043906616),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*XHWSYFLPUW4969082095760678017),
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
				ParmType: reflect.TypeOf(*XHWSYFLPUW1382764017443833647),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*XHWSYFLPUW6303292334592687199),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*XHWSYFLPUW5093615082521380362),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
