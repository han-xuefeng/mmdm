package handler

import (
	"context"
	"mmdm/apiv1"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service"
)

var (
	Datagroup = handlerDatagroup{}
)

type handlerDatagroup struct{}

func (a *handlerDatagroup) DatagroupList (ctx context.Context, req *apiv1.DatagroupListReq) (res *apiv1.DatagroupListRes, err error) {
	list, total, err := service.Datagroup.GetListPage(ctx, model.DatagroupListInput{
		Page: req.Page,
		PageSize: req.PageSize,
	})
	res = &apiv1.DatagroupListRes{
		List: list,
		Total: total,
	}
	return
}

func (a *handlerDatagroup) DatagroupCreate (ctx context.Context, req *apiv1.DatagroupCreateReq) (res *apiv1.DatagroupCreateRes, err error) {
	err = service.Datagroup.Create(ctx, model.DatagroupCreateInput{
		Name: req.Name,
		Prefix: req.Prefix,
		SourceId: req.SourceId,
		GroupConfig: req.GroupConfig,
	})
	return
}

func (a *handlerDatagroup) DatagroupEdit (ctx context.Context, req *apiv1.DatagroupEditReq) (res *apiv1.DatagroupEditRes, err error) {
	err = service.Datagroup.Update(ctx, model.DatagroupEditInput{
		Id:req.Id,
		Name: req.Name,
		Prefix: req.Prefix,
		SourceId: req.SourceId,
		GroupConfig: req.GroupConfig,
	})
	return
}

func (a *handlerDatagroup) DatagroupDelete (ctx context.Context, req *apiv1.DatagroupDeleteReq) (res *apiv1.DatagroupDeleteRes, err error) {
	err = service.Datagroup.DeleteOneById(ctx, model.DatagroupDetailInput{
		Id:req.Id,
	})
	return
}

func (a *handlerDatagroup) DatagroupDetail (ctx context.Context, req *apiv1.DatagroupDetailReq) (res *apiv1.DatagroupDetailRes, err error) {
	var datagroup *entity.Datagroup
	datagroup,err = service.Datagroup.GetOneById(ctx, model.DatagroupDetailInput{
		Id:req.Id,
	})
	res = &apiv1.DatagroupDetailRes{
		Datagroup:datagroup,
	}
	return
}

func (a *handlerDatagroup) SqlExec(ctx context.Context, req *apiv1.DataGroupSqlExecReq) (res *apiv1.DataGroupSqlExecRes, err error) {

	

	return
}
