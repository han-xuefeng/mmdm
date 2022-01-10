package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service/internal/dao"
	"mmdm/internal/service/internal/dto"
	"mmdm/utility"
)

var (
	Admin adminService
)

type adminService struct {

}

func (s *adminService) AdminLogin(ctx context.Context, in model.AdminLoginInput) (err error) {
	var admin *entity.Admin
	err = dao.Admin.Ctx(ctx).Where(dto.Admin{
		UserName: in.UserName,
	}).Scan(&admin)
	if err != nil {
		return err
	}
	if admin == nil {
		return gerror.New(`用户不存在`)
	}
	saltPassword := utility.GenSaltPassword(admin.Salt, in.Password)
	if admin.Password != saltPassword {
		return gerror.New("密码错误，请重新输入")
	}
	if err = Session.SetAdmin(ctx,admin); err != nil {
		return err
	}
	return
}

func (s *adminService) AdminLogout (ctx context.Context) (err error) {
	err = Session.RemoveAdmin(ctx)
	return
}

func (s *adminService) SessionAdminInfo(ctx context.Context) (admin *entity.Admin) {
	admin = Session.GetAdmin(ctx)
	return
}

func (s *adminService) AdminChangePwd(ctx context.Context, in model.AdminChangePwd) (err error)  {
	var admin *entity.Admin
	admin = Session.GetAdmin(ctx)
	if admin == nil {
		return gerror.New(`用户不存在`)
	}
	saltPassword := utility.GenSaltPassword(admin.Salt, in.Password)
	admin.Password = saltPassword
	dao.Admin.Ctx(ctx).Save(admin)
	return
}