package utils

//提供想该集合添加新的自定义头注解的方法
var CommonHeaderAnnoConfig []ReqHeaderInfoWithTag

//这里只提供内置的方法，启动之后，会自动注册这些内置头注解，然后加载用户自定义头注解，然后注解都是预先定义的，也就是dev仍然有这些内容，然后解析rest注解的时候，到这里进行查找，匹配的话就解析到route

//头注解选取策略：0 rest method上边注解，1 rest obj上边的field 和 tag， 2 rest obj上边的注解， 4、如果都没有，查找是否有设置全局的header，如果没有就不加限制
func init() {
	// todo 提供一个init方法，在启动之后，加载所有内置的、自定义的 rest 头注解

	CommonHeaderAnnoConfig = []ReqHeaderInfoWithTag{
		{
			Tag: "@resp-custom-user",

			Type: "common-header",

			Headers: map[string]string{
				"Content-Type": "application/json",
			},

			Consumes: map[string]string{
				"RefererReferer": "www.baidu.com",
			},

			Produces: map[string]string{
				"Accept-Language": "cn",
			},
		},
	}

}
