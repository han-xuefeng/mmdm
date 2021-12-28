package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service/internal/dao"
	"mmdm/internal/service/internal/dto"
)

var (
	Datagroup datagroupService
)

type datagroupService struct {
}

func (s *datagroupService) Create(ctx context.Context, in model.DatagroupCreateInput) (err error) {
	return dao.Datagroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Datagroup.Ctx(ctx).Data(dto.Datagroup{
			Name:     in.Name,
			Prefix: in.Prefix,
			SourceId: in.SourceId,
			GroupConfig:     in.GroupConfig,
		}).Insert()
		return err
	})
}

func (s *datagroupService) Update(ctx context.Context, in model.DatagroupEditInput) (err error) {
	_, err = dao.Datagroup.Ctx(ctx).Data(dto.Datagroup{
		Id:       in.Id,
		Name:     in.Name,
		Prefix: in.Prefix,
		SourceId: in.SourceId,
		GroupConfig:     in.GroupConfig,
	}).Save()
	return err
}

func (s *datagroupService) DeleteOneById(ctx context.Context, in model.DatagroupDetailInput) (err error) {
	_, err = dao.Datagroup.Ctx(ctx).Where(dto.Datagroup{
		Id:       in.Id,
	}).Delete()
	return err
}

func (s *datagroupService) GetListPage(ctx context.Context, in model.DatagroupListInput) (list []*entity.Datagroup, total int, err error) {
	total, err = dao.Datagroup.Ctx(ctx).Count()
	if err != nil {
		return
	}
	err = dao.Datagroup.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list)
	return
}

func (s *datagroupService) GetOneById(ctx context.Context, in model.DatagroupDetailInput) (datagroup *entity.Datagroup, err error) {
	err = dao.Datagroup.Ctx(ctx).Where(dto.Datagroup{
		Id: in.Id,
	}).Scan(&datagroup)
	return
}
