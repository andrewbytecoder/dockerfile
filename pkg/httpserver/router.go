package httpserver

import (
	"context"

	"github.com/andrewbytecoder/dockerfile/pkg/ctx"
	"github.com/andrewbytecoder/dockerfile/pkg/httpserver/api/debug"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Router struct {
	ctx          context.Context
	log          *zap.Logger
	engine       *gin.Engine
	controllers  *Controllers
	serverConfig *ServerConfig
}

func New(ctx *ctx.Ctx) *Router {
	engine := gin.Default()
	return &Router{
		ctx:          ctx.Context(),
		log:          ctx.Logger(),
		engine:       engine,
		serverConfig: &ServerConfig{},
		controllers:  NewControllers(ctx, engine),
	}
}

type ServerConfig struct {
	Addr    string
	Port    int
	TlsPort int
}

func (r *Router) GetServerConfig() *ServerConfig {
	return r.serverConfig
}

func (r *Router) ParseFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&r.serverConfig.Addr, "address", "a", "", "server listen address")
	cmd.Flags().IntVarP(&r.serverConfig.Port, "port", "p", 8080, "server listen port")
	cmd.Flags().IntVarP(&r.serverConfig.TlsPort, "tls-port", "t", 8443, "server listen tls port")
}

type Controllers struct {
	debug *debug.Debug
}

func NewControllers(ctx *ctx.Ctx, engine *gin.Engine) *Controllers {

	debugger := debug.NewDebugger(ctx, engine.Group("/debug"))

	return &Controllers{
		debug: debugger,
	}
}

// Run 启动 Gin 引擎并监听指定地址
//
// @Summary 启动 Gin 引擎并监听指定地址
// @Description 启动 Gin 引擎并监听指定的地址，如果启动失败则记录错误日志
// @Tags GinRouter
// @Param addr path string true "监听地址，可选参数"
// @Success 200 {object} nil "成功启动 Gin 引擎并监听指定地址"
// @Failure 500 {object} error "启动 Gin 引擎失败"
// @Router /runGinRouter [post]
func (r *Router) Run(addr ...string) {
	err := r.engine.Run(addr...)
	if err != nil {
		r.log.Error("ginRouter.engin.Run", zap.Error(err))
		return
	}
}

func (r *Router) RunTLS(addr string, certFile, keyFile string) {
	// "cert.pem", "key.pem"
	err := r.engine.RunTLS(addr, certFile, keyFile)
	if err != nil {
		r.log.Error("ginRouter.engin.RunTLS", zap.Error(err))
		return
	}
}

func (r *Router) Routers() {
	r.controllers.debug.Routers()
}
