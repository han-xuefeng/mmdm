package handler

import (
	"context"
	"mmdm/apiv1"
	"mmdm/internal/model"
	"mmdm/internal/model/entity"
	"mmdm/internal/service"
)

var (
	Datasource = handlerDatasource{}
)

type handlerDatasource struct{}

func (a *handlerDatasource) DatasourceList (ctx context.Context, req *apiv1.DatasourceListReq) (res *apiv1.DatasourceListRes, err error) {
	list, total, err := service.Datasource.GetListPage(ctx, model.DatasourceListInput{
		Page: req.Page,
		PageSize: req.PageSize,
	})
	res = &apiv1.DatasourceListRes{
		List: list,
		Total: total,
	}
	return
}

func (a *handlerDatasource) DatasourceCreate (ctx context.Context, req *apiv1.DatasourceCreateReq) (res *apiv1.DatasourceCreateRes, err error) {
	err = service.Datasource.Create(ctx, model.DatasourceCreateInput{
		Name: req.Name,
		UserName: req.UserName,
		Passport: req.Password,
		Port: req.Port,
		Host: req.Host,
	})
	return
}

func (a *handlerDatasource) DatasourceEdit (ctx context.Context, req *apiv1.DatasourceEditReq) (res *apiv1.DatasourceEditRes, err error) {
	err = service.Datasource.Update(ctx, model.DatasourceEditInput{
		Id:req.Id,
		Name: req.Name,
		UserName: req.UserName,
		Passport: req.Password,
		Port: req.Port,
		Host: req.Host,
	})
	return
}

func (a *handlerDatasource) DatasourceDelete (ctx context.Context, req *apiv1.DatasourceDeleteReq) (res *apiv1.DatasourceDeleteRes, err error) {
	err = service.Datasource.DeleteOneById(ctx, model.DatasourceDetailInput{
		Id:req.Id,
	})
	return
}

func (a *handlerDatasource) DatasourceDetail (ctx context.Context, req *apiv1.DatasourceDetailReq) (res *apiv1.DatasourceDetailRes, err error) {
	var datasource *entity.Datasource
	datasource,err = service.Datasource.GetOneById(ctx, model.DatasourceDetailInput{
		Id:req.Id,
	})
	res = &apiv1.DatasourceDetailRes{
		Datasource:datasource,
	}
	return
}
