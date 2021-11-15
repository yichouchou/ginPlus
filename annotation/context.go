package annotation

import (
	"github.com/gin-gonic/gin"
)

// Context Wrapping gin context to custom context
type Context struct { // 包装gin的上下文到自定义context
	*gin.Context
}

// GetVersion Get the version by req url
func (c *Context) GetVersion() string { // 获取版本号
	return c.Param("version")
}

//WriteJSON 写入json对象
func (c *Context) WriteJSON(obj interface{}) {
	c.JSON(200, obj)
}

// NewCtx Create a new custom context
func NewCtx(c *gin.Context) *Context { // 新建一个自定义context
	return &Context{c}
}

// NewAPIFunc default of custom handlefunc
func NewApiFunc(c *gin.Context) interface{} {
	return NewCtx(c)
}
