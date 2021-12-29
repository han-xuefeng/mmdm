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
	SqlTask sqlTaskService
)

type sqlTaskService struct {
}

func (s *sqlTaskService) GetListPage(ctx context.Context, in model.SqlTaskListInput) (list []*entity.SqlTsak, total int, err error) {
	total, err = dao.SqlTsak.Ctx(ctx).Count()
	if err != nil {
		return
	}
	err = dao.SqlTsak.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list)
	return
}

func (s *sqlTaskService) Create(ctx context.Context, in model.SqlTaskCreateInput) (err error) {
	var datagroup *entity.Datagroup
	datagroup, err = Datagroup.GetOneById(ctx,model.DatagroupDetailInput{
		Id: in.GroupId,
	})

	if err != nil {
		return err
	}

	if datagroup == nil {
		return gerror.New("数据组不存在")
	}

	return dao.SqlTsak.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.SqlTsak.Ctx(ctx).Data(dto.SqlTsak{
			Name:     in.Name,
			Desc: in.Desc,
			GroupId: in.GroupId,
			Sql: in.Sql,
		}).Insert()
		return err
	})
}
