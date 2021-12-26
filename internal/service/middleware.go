package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
	"mmdm/internal/model"
	"net/http"
)

var Middleware = serviceMiddleware{}

type serviceMiddleware struct{}

// Ctx injects custom business context variable into context of current request.
func (s *serviceMiddleware) Ctx(r *ghttp.Request) {
	gsession.NewStorageFile()
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if admin := Session.GetAdmin(r.Context()); admin != nil {
		customCtx.Admin = &model.ContextAdmin{
			Id:       admin.Id,
			Passport: admin.Password,
			UserName: admin.UserName,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}
func (s *serviceMiddleware) CheckAdminLogin(r *ghttp.Request) {
	admin := Session.GetAdmin(r.Context())
	if admin == nil {
		r.Response.WriteStatus(http.StatusForbidden)
	}
	r.Middleware.Next()
}
