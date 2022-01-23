package annotation

import (
	"context"
	"github.com/gin-gonic/gin"
)

// GinBeforeAfterInfo 对象调用前后执行中间件参数
type GinBeforeAfterInfo struct {
	C        *gin.Context
	FuncName string      // 函数名
	Req      interface{} // 调用前的请求参数
	Resp     interface{} // 调用后的返回参数
	Error    error       // 错误信息
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context // 占位参数，可用于存储其他参数，前后连接可用

}

// GinBeforeAfter 对象调用前后执行中间件(支持总的跟对象单独添加)
type GinBeforeAfter interface {
	GinBefore(req *GinBeforeAfterInfo) bool
	GinAfter(req *GinBeforeAfterInfo) bool
}

// DefaultGinBeforeAfter 创建一个默认 BeforeAfter Middleware
type DefaultGinBeforeAfter struct {
}

type timeTrace struct{}
