package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"mmdm/internal/handler"
	"mmdm/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 设置session 目录
			//s.SetConfigWithMap(g.Map{
			//	"SessionMaxAge": time.Minute,
			//	"SessionStorage":gsession.NewStorageFile("D:/opt"),
			//})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware.Ctx,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					handler.Hello,
					handler.Admin,
					handler.Datasource,
				)
			})
			s.Run()
			return nil
		},
	}
)
