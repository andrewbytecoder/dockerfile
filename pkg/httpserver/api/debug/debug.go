package debug

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/andrewbytecoder/dockerfile/pkg/ctx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Debug struct {
	ctx    *ctx.Ctx
	log    *zap.Logger
	router *gin.RouterGroup
}

func NewDebugger(ctx *ctx.Ctx, router *gin.RouterGroup) *Debug {
	return &Debug{
		ctx:    ctx,
		log:    ctx.Logger(),
		router: router,
	}
}

func (d *Debug) Routers() {
	//d.router.GET("/pprof/*any", d.serverDebug)
	d.router.GET("/pprof/*any", d.ginServerDebug)
}

func (d *Debug) serverDebug(c *gin.Context) {

	if !strings.HasPrefix(c.Request.URL.Path, "/debug/pprof/") {
		fmt.Println("debugger: 404")
		http.NotFound(c.Writer, c.Request)
		return
	}

	subpath := strings.TrimPrefix(c.Request.URL.Path, "/debug/pprof/")

	switch subpath {
	case "cmdline":
		pprof.Cmdline(c.Writer, c.Request)
	case "profile":
		pprof.Profile(c.Writer, c.Request)
	case "symbol":
		pprof.Symbol(c.Writer, c.Request)
	case "trace":
		pprof.Trace(c.Writer, c.Request)
	default:
		c.Request.URL.Path = "/debug/pprof/" + subpath
		pprof.Index(c.Writer, c.Request)
	}
}

func (d *Debug) ginServerDebug(c *gin.Context) {
	// 如果路径不以斜杠结尾，则添加斜杠，避免进行重定向

	// 将请求转发给 net/http/pprof 的处理器
	http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
}
