package handler

import (
	"context"
	"fmt"
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

	var ids = make([]int64,0)

	for _, item := range list {
		ids = append(ids, item.SourceId)
	}

	// 处理数据源
	datasourceList, err  := service.Datasource.GetListByIds(ctx, model.DatasourceIdsInput{
		Ids: ids,
	})
	fmt.Println(datasourceList)
	if err != nil {
		return
	}
	var datasourceListMap = make(map[int64]string, 0)
	for _, datasource := range datasourceList {
		datasourceListMap[datasource.Id] = datasource.Name
	}
	fmt.Println(datasourceListMap)
	var outputList = make([]*model.DatagroupListItemOutput,0)
	for _, datagroup := range list {
		item := &model.DatagroupListItemOutput{
			Datagroup:datagroup,
			SourceName: datasourceListMap[datagroup.SourceId],
		}
		outputList = append(outputList, item)
	}
	res = &apiv1.DatagroupListRes{
		List:  outputList,
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
