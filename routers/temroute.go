
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
		annotation.SetVersion(1773220539)
		
		
			
				parm3968124 := new(string)
			
				parm1693233 := new(string)
			
				parm4441728 := new(int)
			
				parm4373843 := new(bind.ReqTest)
			
				
			
			
				parm2702294 := new(bind.ReqTest)
			
				
			
			
				parm6886727 := new(string)
			
				parm2311946 := new(string)
			
				parm6154448 := new(int)
			
				parm7242845 := new(int)
			
			
				parm9311821 := new([]bind.ReqTest)
			
			
				
			
			
				parm6763598 := new([]*bind.ReqTest)
			
			
				parm8721907 := new(bind.ReqTest)
			
			
				parm7841609 := new(string)
			
				parm8106771 := new(int)
			
			
				parm107356 := new(string)
			
				parm2498931 := new(int)
			

		
			
				parm8342123 := new(string)
			
				parm2763675 := new(error)
			
			
				parm3264729 := new(bind.ReqTest)
			
				parm9074263 := new(error)
			
			
				parm7519786 := new(int)
			
				parm1531490 := new(error)
			
			
				parm2341015 := new(int)
			
				parm5978059 := new(error)
			
			
				parm7486031 := new(int)
			
				parm8050417 := new(error)
			
			
				parm117852 := new(int)
			
				parm1106088 := new(error)
			
			
				parm1269234 := new(int)
			
				parm7761403 := new(error)
			
			
				parm537778 := new(bool)
			
			
				parm8482301 := new(bool)
			

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
									ParmType: reflect.TypeOf(*parm3968124),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "password",
									ParmType: reflect.TypeOf(*parm1693233),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm4441728),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								
								{
									ParmName: "hiValue",
									ParmType: reflect.TypeOf(*parm4373843),
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
									ParmType: reflect.TypeOf(*parm8342123),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								

								{
									ParmName: "errHi1",
									ParmType: reflect.TypeOf(*parm2763675),
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
									ParmType: reflect.TypeOf(*parm2702294),
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
									ParmType: reflect.TypeOf(*parm3264729),
									IsMust:   false,
									ParmKind: reflect.Struct,
								},	

								

								{
									ParmName: "errHi2",
									ParmType: reflect.TypeOf(*parm9074263),
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
									ParmType: reflect.TypeOf(*parm6886727),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "password",
									ParmType: reflect.TypeOf(*parm2311946),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm6154448),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								
								{
									ParmName: "year",
									ParmType: reflect.TypeOf(*parm7242845),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "commentHi3",
									ParmType: reflect.TypeOf(*parm7519786),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi3",
									ParmType: reflect.TypeOf(*parm1531490),
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
									ParmType: reflect.TypeOf(*parm9311821),
									IsMust:   false,
									ParmKind: reflect.Slice,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm2341015),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi4",
									ParmType: reflect.TypeOf(*parm5978059),
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
									ParmType: reflect.TypeOf(*parm7486031),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi5",
									ParmType: reflect.TypeOf(*parm8050417),
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
									ParmType: reflect.TypeOf(*parm6763598),
									IsMust:   false,
									ParmKind: reflect.Slice,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm117852),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi4",
									ParmType: reflect.TypeOf(*parm1106088),
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
									ParmType: reflect.TypeOf(*parm8721907),
									IsMust:   false,
									ParmKind: reflect.Struct,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "index",
									ParmType: reflect.TypeOf(*parm1269234),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								

								{
									ParmName: "errHi5",
									ParmType: reflect.TypeOf(*parm7761403),
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
									ParmType: reflect.TypeOf(*parm7841609),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm8106771),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "success",
									ParmType: reflect.TypeOf(*parm537778),
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
									ParmType: reflect.TypeOf(*parm107356),
									IsMust:   false,
									ParmKind: reflect.String,
								},	

								
								{
									ParmName: "age",
									ParmType: reflect.TypeOf(*parm2498931),
									IsMust:   false,
									ParmKind: reflect.Int,
								},	

								},

					Result:     []*utils.Parm{	
		
								

								{
									ParmName: "success",
									ParmType: reflect.TypeOf(*parm8482301),
									IsMust:   false,
									ParmKind: reflect.Bool,
								},	

								},
							},
		
	})
 }
	