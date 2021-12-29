package handler

import (
	"context"
	"mmdm/apiv1"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service"
)

var (
	SqlTask = handlerSqlTask{}
)

type handlerSqlTask struct{}

func (a *handlerSqlTask) SqlTaskList (ctx context.Context, req *apiv1.SqlTaskListReq) (res *apiv1.SqlTaskListRes, err error) {
	list, total, err := service.SqlTask.GetListPage(ctx, model.SqlTaskListInput{
		Page: req.Page,
		PageSize: req.PageSize,
	})
	res = &apiv1.SqlTaskListRes{
		List: list,
		Total: total,
	}
	return
}
func (a *handlerSqlTask) SqlTaskCreate (ctx context.Context, req *apiv1.SqlTaskCreateReq) (res *apiv1.SqlTaskCreateRes, err error) {
	err = service.SqlTask.Create(ctx, model.SqlTaskCreateInput{
		Name: req.Name,
		Desc: req.Desc,
		GroupId: req.GroupId,
		Sql: req.Sql,
	})
	return
}

func (a *handlerSqlTask) SqlTaskEdit (ctx context.Context, req *apiv1.SqlTaskEditReq) (res *apiv1.SqlTaskEditRes, err error) {
	err = service.SqlTask.UpdateById(ctx, model.SqlTaskEditInput{
		Name: req.Name,
		Desc: req.Desc,
		GroupId: req.GroupId,
		Sql: req.Sql,
		Id: req.Id,
	})
	return
}

func (a *handlerSqlTask) SqlTaskDetail (ctx context.Context, req *apiv1.SqlTaskDetailReq) (res *apiv1.SqlTaskDetailRes, err error) {
	var sqlTask *entity.SqlTask
	sqlTask, err = service.SqlTask.GetOneById(ctx, model.SqlTaskDetailInput{
		Id: req.Id,
	})
	res = &apiv1.SqlTaskDetailRes{
		SqlTask: sqlTask,
	}
	return
}

func (a *handlerDatasource) SqlTaskDelete (ctx context.Context, req *apiv1.SqlTaskDeleteReq) (res *apiv1.SqlTaskDeleteRes, err error) {
	err = service.SqlTask.DeleteOneById(ctx, model.SqlTaskDetailInput{
		Id:req.Id,
	})
	return
}

func (a *handlerDatasource) SqlTaskExec(ctx context.Context, req *apiv1.SqlTaskExecReq) (res *apiv1.SqlTaskExecRes, err error) {
	err = service.SqlTask.SqlTaskExec(ctx, model.SqlTaskDetailInput{
		Id:req.Id,
	})
	return
}