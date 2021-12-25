package service

import (
	"context"
	"mmdm/internal/model"
)

var (
	Admin adminService
)

type adminService struct {

}

func (s *adminService) AdminLogin(ctx context.Context, in model.AdminLoginInput) (err error) {
	return err
}