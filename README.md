# ginPlus

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

基于 [Gin](https://github.com/gin-gonic/gin) 框架扩展，支持更简洁的路由注解和自动参数绑定，对齐 Spring MVC 风格。

## 特性

### 🏷️ 简洁路由注解

支持类似 Beego 的注解方式定义路由：

```go
// @POST /user/create
func (s *User) CreateUser(c *gin.Context, req *CreateUserRequest) {
    // 处理创建用户逻辑
    c.JSON(200, gin.H{"status": "ok"})
}
```

### 🔗 自动参数绑定

自动从请求头、URL 参数、JSON Body、Form 表单绑定参数到结构体：

```go
// 自动绑定 GET 参数
// @GET /user/info
func (s *User) GetUserInfo(name string, age int) (string, error) {
    return fmt.Sprintf("Name: %s, Age: %d", name, age), nil
}

// 自动绑定 POST JSON Body
// @POST /user/create
func (s *User) CreateUser(req *CreateUserRequest) (string, error) {
    // req 会自动从 JSON body 绑定
    return "created", nil
}
```

### 📝 参数来源注解

支持显式指定参数来源：

```go
// [name, age] 表示从请求头获取
// {name, age} 表示从请求体获取
// @POST /user/create
func (s *User) Create([name, age string], {info RequestBody}) {
    // ...
}
```

## 快速开始

### 安装

```bash
go get github.com/yichouchou/ginPlus
```

### 创建项目

```bash
mkdir myapp && cd myapp
go mod init myapp
go get github.com/yichouchou/ginPlus
```

### 示例代码

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yichouchou/ginPlus/annotation"
    "github.com/yichouchou/ginPlus/examples/controller"
)

func main() {
    engine := gin.Default()
    base := annotation.New()
    base.Dev(true)
    base.Register(engine, new(controller.UserRest))
    engine.Run(":8088")
}
```

### 定义 Controller

```go
package controller

import (
    "github.com/gin-gonic/gin"
)

// UserRest 用户接口
type UserRest struct{}

// @POST /user/create
// content-type=application/json
func (u *UserRest) Create(c *gin.Context, req *CreateRequest) {
    c.JSON(200, gin.H{
        "code": 0,
        "msg":  "success",
        "data": req,
    })
}

// CreateRequest 创建用户请求
type CreateRequest struct {
    Name  string `json:"name" binding:"required"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}
```

## 项目结构

```
ginPlus/
├── annotation/          # 注解处理核心
│   ├── ProcessAnnot.go # 注解解析器
│   ├── middleware.go   # 中间件
│   └── context.go      # 上下文扩展
├── bind/              # 参数自动绑定
│   └── autoBind.go     # 自动参数绑定实现
├── utils/             # 工具函数
│   ├── headerCheck.go  # 请求头校验
│   ├── parameterCheck.go # 参数校验
│   ├── producesSet.go  # 响应类型设置
│   └── consumersCheck.go # Accept 头校验
└── examples/          # 示例代码
    ├── controller/     # 控制器示例
    └── routers/       # 生成的路由
```

## 参数绑定说明

| 注解格式 | 来源 | 示例 |
|---------|------|------|
| [param1, param2] | 请求头/Query | [name, age] |
| {param1, param2} | 请求体 JSON | {name, age} |
| 无注解 | 自动推断 | 根据方法签名 |

## 性能测试

使用 jmeter 简单测试结果：

| 场景 | 请求数 | 平均响应时间 | TPS |
|------|--------|-------------|-----|
| Gin | 100万 | 21ms | ~22917 |
| ginPlus | 100万 | 21ms | ~22901 |

性能损耗约 1%，对整体性能影响极小。

## 依赖

- Go 1.16+
- Gin v1.7.2+
- go-playground/validator v10.4.1+

## 注意事项

⚠️ 项目目前处于维护状态，部分代码（如 main.go）可能被注释。如需使用请参考 examples。

## License

MIT License
