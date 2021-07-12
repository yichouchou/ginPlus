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
