package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"` //结构中添加header标签
	Domain string `header:"Domain"`
}

type tesHeader struct {
	Rate   int    `header:"Rate"` //结构中添加header标签
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		h := testHeader{}
		//​    //ShouldBindHeader是c.ShouldBindWith(obj, binding.Header)的快捷方法
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		fmt.Printf("%#v", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})
	r.Run()
}

//模拟请求// curl -H "rate:300" -H "domain:music" http://localhost:8080/// 参考输出:// {"Domain":"music","Rate":300}
