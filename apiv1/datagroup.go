package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mmdm/internal/model/entity"
)

type DatagroupListReq struct {
	g.Meta   `path:"/datagroup/list" method:"get" tags:"Datagroup" summary:"数据源列表"`
	Page     int `v:"required" p:"page"`
	PageSize int `v:"required" p:"page_size"`
}

type DatagroupListRes struct {
	List  []*entity.Datagroup
	Total int
}

type DatagroupCreateReq struct {
	g.Meta   `path:"/datagroup/create" method:"post" tags:"Datagroup" summary:"数据源创建"`
	Name     string `v:"required" p:"name"`
	Prefix     string `v:"required" p:"prefix"`
	SourceId     int64 `v:"required" p:"source_id"`
	GroupConfig string `p:"group_config"`
}

type DatagroupCreateRes struct {
}

type DatagroupEditReq struct {
	g.Meta   `path:"/datagroup/edit" method:"post" tags:"Datagroup" summary:"数据源修改"`
	Id       int64  `v:"required" p:"id"`
	Name     string `v:"required" p:"name"`
	Prefix     string `v:"required" p:"prefix"`
	SourceId     int64 `v:"required" p:"source_id"`
	GroupConfig string `p:"group_config"`
}

type DatagroupEditRes struct {
}

type DatagroupDeleteReq struct {
	g.Meta   `path:"/datagroup/delete" method:"get" tags:"Datagroup" summary:"数据源删除"`
	Id       int64  `v:"required" p:"id"`
}

type DatagroupDeleteRes struct {
}

type DatagroupDetailReq struct {
	g.Meta   `path:"/datagroup/detail" method:"get" tags:"Datagroup" summary:"数据源删除"`
	Id       int64  `v:"required" p:"id"`
}

type DatagroupDetailRes struct {
	*entity.Datagroup
}

type DataGroupSqlExecReq struct {
	g.Meta   `path:"/datagroup/sql_exec" method:"post" tags:"Datagroup" summary:"sql执行"`
	Id       int64  `v:"required" p:"id"`
	Sql      string `v:"required" p:"sql"`
}

type DataGroupSqlExecRes struct {}
