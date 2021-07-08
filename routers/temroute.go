package routers

import (
	"reflect"

	"ginPlus/annotation"
	"ginPlus/bind"
	"ginPlus/utils"
)

func init() {
	b := new(bind.ReqTest)
	utils.SetVersion(1625627016)
	annotation.AddGenOne("Hello.Hi", utils.GenComment{
		RouterPath: "hello.hi",
		Note:       "",
		Methods:    []string{"ANY"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmKind: reflect.String,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "password",
				ParmKind: reflect.String,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "age",
				ParmKind: reflect.Int,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "hiValue",
				ParmKind: reflect.Struct,
				//reflect.TypeOf(new(bind.ReqTest)).Kind(), 这里 是否可以考虑直接 reflect.Struct
				ParmType: reflect.TypeOf(*b), //这里是传递值参数
				//由于在启动后不论dev 还是生产，运行后都可以加载对应参数，所以这里不用ParmType字段貌似也可以!! 在生产环境，无法做到注入 都会多一个 *  todo 确定了可以不用，因为无法很好的存放
				IsMust: false,
			},
			{
				ParmName: "hi",
				ParmKind: reflect.Ptr,
				//reflect.TypeOf(new(bind.ReqTest)).Kind(), 这里 是否可以考虑直接 reflect.Struct
				ParmType: reflect.TypeOf(new(bind.ReqTest)), //这里传递指针参数
				//由于在启动后不论dev 还是生产，运行后都可以加载对应参数，所以这里不用ParmType字段貌似也可以!! 在生产环境，无法做到注入 都会多一个 *  todo 确定了可以不用，因为无法很好的存放
				IsMust: false,
			},
		},
		Result: []*utils.Parm{
			{
				ParmName: "name",
				ParmKind: reflect.String,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},

			{
				ParmName: "password",
				ParmKind: reflect.Interface,
				ParmType: reflect.TypeOf(new(string)),
				IsMust:   false,
			},
		},
	})
}
