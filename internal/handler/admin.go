package handler

import (
	"context"
	"mmdm/apiv1"
	"mmdm/internal/model"
	"mmdm/internal/service"
)

var (
	Admin = handlerAdmin{}
)

type handlerAdmin struct{}

func (a *handlerAdmin) AdminLogin(ctx context.Context, req *apiv1.AdminLoginReq) (res *apiv1.AdminLoginRes, err error) {
	err = service.Admin.AdminLogin(ctx, model.AdminLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil,err
	}
	res = &apiv1.AdminLoginRes{
		Token: req.UserName,
	}
	return
}

func (a *handlerAdmin) AdminLogout(ctx context.Context, req *apiv1.AdminLogoutReq)(res *apiv1.AdminLoutRes,err error){
	err = service.Admin.AdminLogout(ctx)
	return
}

func (a *handlerAdmin) AdminInfo(ctx context.Context, req *apiv1.AdminInfoReq) (res *apiv1.AdminInfoRes, err error){
	admin := service.Admin.SessionAdminInfo(ctx)
	res = &apiv1.AdminInfoRes{
		ID: admin.Id,
		Name: admin.UserName,
		Introduction: "I am a super administrator",
		Roles:[]string{"admin"},
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	return
}

func (a *handlerAdmin) AdminChangePwd(ctx context.Context, req *apiv1.AdminChangePwdReq)(res *apiv1.AdminChangePwdRes,err error){
	err = service.Admin.AdminChangePwd(ctx, model.AdminChangePwd{
		Password: req.Password,
	})
	return
}