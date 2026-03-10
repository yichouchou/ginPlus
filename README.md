# ginPlus

基于 Gin 框架的增强型 RESTful API 框架，支持简洁路由注解和自动参数绑定。

## 特性

- 🚀 **简洁路由注解** - 通过方法注释定义路由，如 `@GET /hello`、`@POST /user`
- 🔗 **自动参数绑定** - 对齐 Spring MVC 的参数绑定方式，支持从请求头、URL参数、请求体自动绑定
- 📦 **请求头/响应头控制** - 支持通过结构体 tag 定义请求头和响应头
- 📚 **Swagger 集成** - 自动生成 API 文档
- ⚡ **高性能** - 根据测试，自动参数绑定性能损耗仅约 1%

## 安装

```bash
go get github.com/yichouchou/ginPlus
```

## 快速开始

### 1. 定义 Controller

使用注解定义路由，支持 GET、POST、PUT、DELETE、PATCH 等 HTTP 方法：

```go
type Hello struct {
    ReqContentType utils.ReqHeaderInfo `head:"application/json"`
    RespContentType utils.RespHeaderInfo `head:"application/json"`
}

// @GET /hello
// @resp-custom-user
func (s *Hello) Hi(name string, age int) (string, error) {
    return fmt.Sprintf("Hello, %s! You are %d years old.", name, age), nil
}
```

### 2. 注册路由

```go
func main() {
    engine := gin.Default()
    base := annotation.New()
    base.Dev(true)
    base.Register(engine, new(Hello))
    
    engine.Run(":8088")
}
```

### 3. 访问 API

```
GET /hello?name=张三&age=25
```

## 注解语法

### 路由注解

| 注解 | 说明 | 示例 |
|------|------|------|
| `@GET` | GET 请求 | `@GET /user` |
| `@POST` | POST 请求 | `@POST /user` |
| `@PUT` | PUT 请求 | `@PUT /user` |
| `@DELETE` | DELETE 请求 | `@DELETE /user` |
| `@PATCH` | PATCH 请求 | `@PATCH /user` |

### 参数注解

参数可以通过注释指定来源：

```go
// [name string, password string] 表示从请求头获取
// {body *RequestBody} 表示从请求体获取
// @POST /login
func (s *User) Login([name string, password string], body *LoginReq) (string, error)
```

### 响应头定义

通过结构体字段定义请求头和响应头：

```go
type BaseRestSetting struct {
    CommonHeader string
}

type Hello struct {
    BaseRestSetting
    ReqContentType  utils.ReqHeaderInfo `head:"application/json"`
    RespContentType utils.RespHeaderInfo `head:"application/json"`
}
```

## 参数绑定规则

### GET 请求

- 单一结构体参数：从 URL query 绑定
- 多个基本类型参数：依次从 URL query 获取

### POST 请求

- 单一参数（指针/结构体）：从请求体 JSON 绑定
- 多个参数：部分从请求头获取，部分从表单/请求体获取

### 请求头绑定

使用 `[]` 包裹的参数从请求头获取：
```go
// [token string, appId string]
```

### 请求体绑定

使用 `{}` 包裹的参数从请求体获取：
```go
// {user *UserInfo}
```

## 项目结构

```
ginPlus/
├── annotation/          # 注解处理核心
│   ├── ProcessAnnot.go # 路由解析与注册
│   ├── context.go      # 上下文扩展
│   └── middleware.go   # 中间件
├── bind/               # 参数自动绑定
│   └── autoBind.go
├── routers/            # 生成的路由文件
├── swagger/            # Swagger 文档
├── utils/              # 工具函数
├── hello.go           # 示例代码
└── main.go            # 入口文件
```

## 性能测试

使用 jmeter 测试结果（500 线程，2000 次，总计 100 万请求）：

| 请求数 | 框架 | 平均响应时间 | TPS |
|--------|------|-------------|-----|
| 100万 | gin | 21ms | 22901 |
| 100万 | ginPlus | 21ms | 22917 |

自动参数绑定的性能损耗约为 **1%**，对实际使用影响微乎其微。

## 依赖

- Go 1.16+
- gin-gonic/gin v1.7+
- go-playground/validator v10
- swaggo 相关包

## License

MIT
