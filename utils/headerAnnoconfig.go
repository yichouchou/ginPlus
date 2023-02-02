package utils

//todo 提供向该集合添加新的自定义头注解的方法

// todo 统一rest请求 响应 以及参数绑定 校验 解析 等等的优先顺序，配置来源包括：排序前则优先级高
// 1. CommonHeaderAnnoConfig 标准 公共 全局配置
// 2. rest obj对象tag 上的rest局部配置 就是该obj对象的所有rest方法都有此类配置
// 3. rest方法上局部方法的请求 响应配置内容，只针对该方法生效。（注意，这里方法上的注解可能是预先定义好的，
//也可以是自定义好的；比如 @ resp-custom-user 可能在多个rest方法上使用，也可能不用，预先定义好其中的内容项 ）

// todo 由于提供了多种rest 请求 响应的配置方式，自由组合情况复杂
// CommonHeaderAnnoConfig 全局生效，可以自由配置其中内容，比如固定需要某个请求头等等
// todo： rest obj对象，可以在对象头上添加注解  考虑进去 ；
// todo 关于参数校验，如果没有额外注解，则默认为： 非必须传，有则绑定，无则赋空对象，不对空对象负责；
// todo 关于请求响应 header的配置：尽量简洁，然后个性化的配置参考以及对标 spring mvc

// 标准公共基础header头内容，就是所有的rest都有的内容
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
