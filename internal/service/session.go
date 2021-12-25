package service

import (
	"context"
	"mmdm/internal/consts"
	"mmdm/internal/model/entity"
)

var Session = serviceSession{}

type serviceSession struct{}

// SetUser sets user into the session.
func (s *serviceSession) SetAdmin(ctx context.Context, admin *entity.Admin) error {
	return Context.Get(ctx).Session.Set(consts.AdminSessionKey, admin)
}

func (s *serviceSession) GetAdmin(ctx context.Context) *entity.Admin {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.AdminSessionKey); !v.IsNil() {
			var admin *entity.Admin
			_ = v.Struct(&admin)
			return admin
		}
	}
	return nil
}
