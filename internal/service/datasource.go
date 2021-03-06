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
			Name:     in.Name,
			UserName: in.UserName,
			Password: in.Passport,
			Host:     in.Host,
			Port:     in.Port,
		}).Insert()
		return err
	})
}

func (s *datasourceService) Update(ctx context.Context, in model.DatasourceEditInput) (err error) {
	_, err = dao.Datasource.Ctx(ctx).Data(dto.Datasource{
		Id:       in.Id,
		Name:     in.Name,
		UserName: in.UserName,
		Password: in.Passport,
		Host:     in.Host,
		Port:     in.Port,
	}).Save()
	return err
}

func (s *datasourceService) DeleteOneById(ctx context.Context, in model.DatasourceDetailInput) (err error) {
	_, err = dao.Datasource.Ctx(ctx).Where(dto.Datasource{
		Id:       in.Id,
	}).Delete()
	return err
}

func (s *datasourceService) GetListPage(ctx context.Context, in model.DatasourceListInput) (list []*entity.Datasource, total int, err error) {
	total, err = dao.Datasource.Ctx(ctx).Count()
	if err != nil {
		return
	}
	err = dao.Datasource.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list)
	return
}

func (s *datasourceService) GetListAll(ctx context.Context) (list []*entity.Datasource, err error) {
	err = dao.Datasource.Ctx(ctx).Scan(&list)
	return
}

func (s *datasourceService) GetOneById(ctx context.Context, in model.DatasourceDetailInput) (datasource *entity.Datasource, err error) {
	err = dao.Datasource.Ctx(ctx).Where(dto.Datasource{
		Id: in.Id,
	}).Scan(&datasource)
	return
}

func (s *datasourceService) GetListByIds(ctx context.Context, in model.DatasourceIdsInput) (list []*entity.Datasource, err error) {
	err = dao.Datasource.Ctx(ctx).WhereIn("id", in.Ids).Scan(&list)
	return
}
