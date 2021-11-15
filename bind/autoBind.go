package bind

//todo 参数自动绑定逻辑

//需要统一规范
//
// 1。get请求都是通过请求头传参数 类似【name string, password string, age int 】一个个获取， 或者 【hiValue bind.ReqTest】从url解析出结构体放入

// 2。post请求可以把参数放到请求头 也可以放请求体， 如果不特别指定（规则暂时未定），那么从请求体中获取

// 3。表单提交参数从请求体form中根据键值对获取value，然后装填到参数内，比如下方为java post请求传参
/*
（ModelMap modelMap, HttpServletRequest request,
HttpServletResponse response, String resourceId,
String resourceDomain, String resourceVersion,
String attributeName, String value, String fromApp,
String token, String timeZone）
*/

//todo gin.shoudbind非常友好，暂时可以直接拿来用,以后需要个性化修改和调整。

//todo post请求，然后先去找请求头的参数，依次绑定，然后找请求体的参数。
//todo 如果部分注释，则先按照已经有的去寻找，剩下的理解为是其他部分（比如写了请求头的，那么剩下的就是请求体）
//
//todo 如果未有注解，则按照默认的方式（查看参数个数，post请求多参数的话就是表单方式，1个参数就是默认请求体内容）
//todo get请求，全部都在请求体内，依次绑定即可。（注意：如果存在结构体，那么通过shoudbind去绑定，如果都是基本数据类型，则通过query去绑定）
