
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
		annotation.SetVersion(1773231227)
		
		
			
				parm6860590 := new(string)
			
				parm4374764 := new(string)
			
				parm8156836 := new(int)
			
				parm5919727 := new(bind.ReqTest)
			
				
			
			
				parm6697646 := new(bind.ReqTest)
			
				
			
			
				parm803589 := new(string)
			
				parm4577502 := new(string)
			
				parm2346044 := new(int)
			
				parm3975074 := new(int)
			
			
				parm1002271 := new([]bind.ReqTest)
			
			
				
			
			
				parm2247818 := new([]*bind.ReqTest)
			
			
				parm6380668 := new(bind.ReqTest)
			
			
				parm2598702 := new(string)
			
				parm7781867 := new(int)
			
			
				parm5022392 := new(string)
			
				parm5101788 := new(int)
			

		
			
				parm3298168 := new(string)
			
				parm4499461 := new(error)
			
			
				parm3022635 := new(bind.ReqTest)
			
				parm967254 := new(error)
			
			
				parm7255155 := new(int)
			
				parm6588146 := new(error)
			
			
				parm131986 := new(int)
			
				parm3063569 := new(error)
			
			
				parm6405499 := new(int)
			
				parm8936093 := new(error)
			
			
				parm127227 := new(int)
			
				parm9226618 := new(error)
			
			
				parm6221826 := new(int)
			
				parm3438766 := new(error)
			
			
				parm5428601 := new(bool)
			
			
				parm4023752 := new(bool)
			

		annotation.AddGenOne("main.Hello.Hi1", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi1",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block1",

					Note:       "",

					Methods:    []string{ "GET" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "name",
									ParmType: reflect.TypeOf(*parm6860590),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "password",
									ParmType: reflect.TypeOf(*parm4374764),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm8156836),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								
								{
									ParmName: "hiValue",
									ParmType: reflect.TypeOf(*parm5919727),
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

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "commentHi1",
									ParmType: reflect.TypeOf(*parm3298168),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								

								{
									ParmName: "errHi1",
									ParmType: reflect.TypeOf(*parm4499461),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi2", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi2",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block2",

					Note:       "",

					Methods:    []string{ "POST" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "hiValue",
									ParmType: reflect.TypeOf(*parm6697646),
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

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "commentHi2",
									ParmType: reflect.TypeOf(*parm3022635),
									IsMust:   false,
									ParmKind: reflect.Struct,
								},	

								

								{
									ParmName: "errHi2",
									ParmType: reflect.TypeOf(*parm967254),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi3", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi3",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block3",

					Note:       "",

					Methods:    []string{ "GET" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "name",
									ParmType: reflect.TypeOf(*parm803589),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "password",
									ParmType: reflect.TypeOf(*parm4577502),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm2346044),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								
								{
									ParmName: "year",
									ParmType: reflect.TypeOf(*parm3975074),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "commentHi3",
									ParmType: reflect.TypeOf(*parm7255155),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi3",
									ParmType: reflect.TypeOf(*parm6588146),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi4", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi4",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block4",

					Note:       "",

					Methods:    []string{ "POST" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "List",
									ParmType: reflect.TypeOf(*parm1002271),
									IsMust:   false,
									ParmKind: reflect.Slice,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm131986),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi4",
									ParmType: reflect.TypeOf(*parm3063569),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi5", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi5",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block5",

					Note:       "",

					Methods:    []string{ "GET" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "reqList",
									ParmType: reflect.TypeOf(new(bind.ReqTest)),
									IsMust:   false,
									ParmKind: reflect.Ptr,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm6405499),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi5",
									ParmType: reflect.TypeOf(*parm8936093),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi6", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi6",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block6",

					Note:       "",

					Methods:    []string{ "POST" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "List",
									ParmType: reflect.TypeOf(*parm2247818),
									IsMust:   false,
									ParmKind: reflect.Slice,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm127227),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi4",
									ParmType: reflect.TypeOf(*parm9226618),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("main.Hello.Hi7", utils.GenRouterInfo{
		HandFunName: "main.Hello.Hi7",
		RouterPath: "",
		Note:        "",
		Methods:    []string{ "" },
        Headers:    map[string]string{
					
					},
		Consumes:    map[string]string{
					
					},
 		Produces:    map[string]string{
					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/block7",

					Note:       "",

					Methods:    []string{ "GET" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
								},

					Produces:   map[string]string{
								
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "reqList",
									ParmType: reflect.TypeOf(*parm6380668),
									IsMust:   false,
									ParmKind: reflect.Struct,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm6221826),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi5",
									ParmType: reflect.TypeOf(*parm3438766),
									IsMust:   false,
									ParmKind: reflect.Interface,
								},	

								},
							},
		
	})
annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.LogOutUser",
		RouterPath: "/UserRest222",
		Note:        "",
		Methods:    []string{ "POST" },
        Headers:    map[string]string{
					
            			"Content-Type" : "application/json",
  					
					},
		Consumes:    map[string]string{
					
            			"RefererReferer":"www.baidu.com",
  					
					},
 		Produces:    map[string]string{
					
            			"Accept-Language": "cn",
  					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/LogOutUser2",

					Note:       "",

					Methods:    []string{ "GET" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
            						"RefererReferer": "www.baidu.com",
											
								},

					Produces:   map[string]string{
								
            						"Accept-Language": "cn",
											
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "name",
									ParmType: reflect.TypeOf(*parm2598702),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm7781867),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "success",
									ParmType: reflect.TypeOf(*parm5428601),
									IsMust:   false,
									ParmKind: reflect.Bool,
								},	

								},
							},
		
	})
annotation.AddGenOne("github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser", utils.GenRouterInfo{
		HandFunName: "github.com/yichouchou/ginPlus/bind.UserRest222.RegistUser",
		RouterPath: "/UserRest222",
		Note:        "",
		Methods:    []string{ "POST" },
        Headers:    map[string]string{
					
            			"Content-Type" : "application/json",
  					
					},
		Consumes:    map[string]string{
					
            			"RefererReferer":"www.baidu.com",
  					
					},
 		Produces:    map[string]string{
					
            			"Accept-Language": "cn",
  					
					},
		GenComment: &utils.GenComment{

					RouterPath: "/RegistUser2",

					Note:       "",

					Methods:    []string{ "POST" },

					Headers:     map[string]string{
								
								},

					Consumes:   map[string]string{
								
            						"RefererReferer": "www.baidu.com",
											
								},

					Produces:   map[string]string{
								
            						"Accept-Language": "cn",
											
								},

					Parms:      []*utils.Parm{
		
								
								{
									ParmName: "name",
									ParmType: reflect.TypeOf(*parm5022392),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm5101788),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "success",
									ParmType: reflect.TypeOf(*parm4023752),
									IsMust:   false,
									ParmKind: reflect.Bool,
								},	

								},
							},
		
	})
 }
	