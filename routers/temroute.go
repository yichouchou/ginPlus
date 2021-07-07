package routers

import (
	"reflect"

	"ginPlus/annotation"
)

func init() {
	annotation.SetVersion(1625627016)
	annotation.AddGenOne("Hello.Hi", annotation.GenComment{
		RouterPath: "hello.hi",
		Note:       "",
		Methods:    []string{"ANY"},
		Parms: []*annotation.Parm{

			{
				ParmName: "name",
				ParmKind: reflect.String,
				//ParmType: reflect.TypeOf(new(string)),
				IsMust: false,
			},

			{
				ParmName: "password",
				ParmKind: reflect.String,
				//ParmType: reflect.TypeOf(new(string)),
				IsMust: false,
			},

			{
				ParmName: "age",
				ParmKind: reflect.Int,
				//ParmType: reflect.TypeOf(new(string)),
				IsMust: false,
			},

			{
				ParmName: "hi",
				ParmKind: reflect.Struct,
				//reflect.TypeOf(new(bind.ReqTest)).Kind(), 这里 是否可以考虑直接 reflect.Struct
				//ParmType: reflect.TypeOf(new(bind.ReqTest)),
				//由于在启动后不论dev 还是生产，运行后都可以加载对应参数，所以这里不用ParmType字段貌似也可以!! 在生产环境，无法做到注入 都会多一个 *  todo 确定了可以不用，因为无法很好的存放
				IsMust: false,
			},
		},
	})
}
