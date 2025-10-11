package httpserver

import (
	"fmt"

	"github.com/andrewbytecoder/dockerfile/pkg/ctx"
	"github.com/andrewbytecoder/dockerfile/pkg/httpserver"
	"github.com/spf13/cobra"
)

func GetHttpServerCmd(ctx *ctx.Ctx) []*cobra.Command {
	var cmds []*cobra.Command
	cmds = append(cmds, newHttpServer(ctx))

	return cmds
}

// newHttpServer returns a cobra command for fetching versions
func newHttpServer(ctx *ctx.Ctx) *cobra.Command {
	httpserver := httpserver.New(ctx)

	cmd := &cobra.Command{
		Use:     "httpserver",
		Short:   "nexa httpserver",
		Long:    `nexa httpserver`,
		Example: `nexa httpserver`,
		// stop printing usage when the command errors
		SilenceUsage: true,
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		// 添加路由信息
		httpserver.Routers()

		addr := fmt.Sprintf("%s:%d", httpserver.GetServerConfig().Addr, httpserver.GetServerConfig().Port)

		// 启动服务
		httpserver.Run(addr)
	}

	httpserver.ParseFlags(cmd)
	return cmd
}
