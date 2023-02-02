package routers

import (
	"reflect"

	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/examples/controller"
	"github.com/yichouchou/ginPlus/utils"
)

func init() {
	annotation.SetVersion(1675324967)

	NQKRVEOAEI5577006791947779410 := new(string)

	NQKRVEOAEI8674665223082153551 := new(int)

	PEIDSEUDLU6129484611666145821 := new(controller.DemoRest)

	PEIDSEUDLU3916589616287113937 := new(string)

	PEIDSEUDLU6334824724549167320 := new(int)

	PEIDSEUDLU4037200794235010051 := new(bool)

	PEIDSEUDLU605394647632969758 := new(bool)

	annotation.AddGenOne("UserRest.LogOutUser", utils.GenComment{
		RouterPath: "/LogOutUser",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*NQKRVEOAEI5577006791947779410),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*NQKRVEOAEI8674665223082153551),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "rest",
				ParmType: reflect.TypeOf(*PEIDSEUDLU6129484611666145821),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*PEIDSEUDLU4037200794235010051),
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
				ParmType: reflect.TypeOf(*PEIDSEUDLU3916589616287113937),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*PEIDSEUDLU6334824724549167320),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*PEIDSEUDLU605394647632969758),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
