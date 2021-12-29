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

func (s *sqlTaskService) GetListPage(ctx context.Context, in model.SqlTaskListInput) (list []*entity.SqlTask, total int, err error) {
	total, err = dao.SqlTask.Ctx(ctx).Count()
	if err != nil {
		return
	}
	err = dao.SqlTask.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list)
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

	return dao.SqlTask.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.SqlTask.Ctx(ctx).Data(dto.SqlTask{
			Name:     in.Name,
			Desc: in.Desc,
			GroupId: in.GroupId,
			Sql: in.Sql,
		}).Insert()
		return err
	})
}

func (s *sqlTaskService) GetOneById(ctx context.Context, in model.SqlTaskDetailInput) (sqlTask *entity.SqlTask, err error) {
	err = dao.SqlTask.Ctx(ctx).Where(dto.SqlTask{
		Id: in.Id,
	}).Scan(&sqlTask)
	return
}

func (s *sqlTaskService) UpdateById(ctx context.Context, in model.SqlTaskEditInput) (err error) {
	var sqlTask *entity.SqlTask
	sqlTask, err = s.GetOneById(ctx, model.SqlTaskDetailInput{
		Id: in.Id,
	})

	if err != nil {
		return err
	}

	if sqlTask == nil {
		return gerror.New("sql任务不存在")
	}

	if sqlTask.Status != 0 {
		return gerror.New("sql任务已执行或执行结束,不允许修改")
	}

	_, err = dao.SqlTask.Ctx(ctx).Data(dto.SqlTask{
		Id:       in.Id,
		Name:     in.Name,
		Desc: in.Desc,
		GroupId: in.GroupId,
		Sql: in.Sql,
	}).Save()
	return err
}

func (s *sqlTaskService) SqlTaskDetail(ctx context.Context, in model.SqlTaskDetailInput) (sqlTask *entity.SqlTask,err error){
	sqlTask, err = s.GetOneById(ctx, in)
	if err != nil {
		return
	}

	if sqlTask == nil {
		return nil,gerror.New("sql任务不存在")
	}

	return
}


func (s *sqlTaskService) SqlTaskExec(ctx context.Context, in model.SqlTaskDetailInput) (err error){
	var sqlTask *entity.SqlTask
	sqlTask, err = s.GetOneById(ctx, in)

	if err != nil {
		return err
	}

	if sqlTask == nil {
		return gerror.New("sql任务不存在")
	}

	if sqlTask.Status != 0 {
		return gerror.New("sql任务已执行或执行结束,不允许重复执行")
	}

	// todo: 这里写入队列 然后队列执行

	_, err = dao.SqlTask.Ctx(ctx).Where(dto.SqlTask{
		Id: in.Id,
	}).Update(dto.SqlTask{
		Status: 1,
	})
	return err
}

func (s *sqlTaskService) SqlTaskDelete(ctx context.Context, in model.SqlTaskDetailInput) (err error){
	var sqlTask *entity.SqlTask
	sqlTask, err = s.GetOneById(ctx, in)

	if err != nil {
		return err
	}

	if sqlTask == nil {
		return gerror.New("sql任务不存在")
	}

	if sqlTask.Status != 0 {
		return gerror.New("sql任务已执行或执行结束,不允许删除")
	}

	_, err = dao.SqlTask.Ctx(ctx).Where(dto.SqlTask{
		Id:       in.Id,
	}).Delete()
	return err
}

func (s *sqlTaskService) DeleteOneById(ctx context.Context, in model.SqlTaskDetailInput) (err error) {
	_, err = dao.SqlTask.Ctx(ctx).Where(dto.SqlTask{
		Id:       in.Id,
	}).Delete()
	return err
}

