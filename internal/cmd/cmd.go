package cmd

import (
	"context"
	"mmdm/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"mmdm/internal/handler"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware.Ctx,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					handler.Hello,
					handler.Admin,
				)
			})
			s.Run()
			return nil
		},
	}
)
