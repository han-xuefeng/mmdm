package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service/internal/dao"
	"mmdm/internal/service/internal/dto"
)

var (
	Datasource datasourceService
)

type datasourceService struct {

}

func (s *datasourceService) Create(ctx context.Context, in model.DatasourceCreateInput) (err error) {

	if in.UserName == "" {
		return gerror.New("UserName Empty Is Not Allow")
	}

	if in.Port == "" {
		return gerror.New("Port Empty Is Not Allow")
	}

	if in.Passport == "" {
		return gerror.New("Passport Empty Is Not Allow")
	}

	if in.Host == "" {
		return gerror.New("Host Empty Is Not Allow")
	}

	return dao.Datasource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Datasource.Ctx(ctx).Data(dto.Datasource{
			Name: in.Name,
			UserName: in.UserName,
			Password: in.Passport,
			Host: in.Host,
			Port: in.Port,
		}).Insert()
		return err
	})
}

func (s *datasourceService) GetListPage(ctx context.Context, in model.DatasourceListInput) (list []*entity.Datasource, total int,err error) {
	total, err = dao.Datasource.Ctx(ctx).Count()
	if err != nil {
		return
	}
	err = dao.Datasource.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list)
	return
}
