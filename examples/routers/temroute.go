package routers

import (
	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"
	"reflect"
)

func init() {
	annotation.SetVersion(1772677437)

	FMWHJDCGEE3372058799042066494 := new(string)

	FMWHJDCGEE6616879910558991571 := new(int)

	FMWHJDCGEE3558074866267376968 := new(string)

	FMWHJDCGEE2148399144987937656 := new(int)

	FMWHJDCGEE1218531123504490226 := new(bool)

	FMWHJDCGEE2180221784322843576 := new(bool)

	annotation.AddGenOne("UserRest.LogOutUser", utils.GenComment{
		RouterPath: "/LogOutUser",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*FMWHJDCGEE3372058799042066494),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*FMWHJDCGEE6616879910558991571),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*FMWHJDCGEE1218531123504490226),
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
				ParmType: reflect.TypeOf(*FMWHJDCGEE3558074866267376968),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*FMWHJDCGEE2148399144987937656),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*FMWHJDCGEE2180221784322843576),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
