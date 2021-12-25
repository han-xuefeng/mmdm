package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"mmdm/internal/consts"
	"mmdm/internal/model"
)

var Context = serviceContext{}

type serviceContext struct{}

// Init initializes and injects custom business context object into request context.
func (s *serviceContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get retrieves and returns the user object from context.
// It returns nil if nothing found in given context.
func (s *serviceContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser injects business user object into context.
func (s *serviceContext) SetAdmin(ctx context.Context, ctxAdmin *model.ContextAdmin) {
	s.Get(ctx).Admin = ctxAdmin
}
