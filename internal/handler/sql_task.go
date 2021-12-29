package handler

import (
	"context"
	"mmdm/apiv1"
	"mmdm/internal/model"
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
