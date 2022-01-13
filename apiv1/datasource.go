package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mmdm/internal/model/entity"
)

type DatasourceListReq struct {
	g.Meta   `path:"/datasource/list" method:"get" tags:"Datasource" summary:"数据源列表"`
	Page     int `v:"required" p:"page"`
	PageSize int `v:"required" p:"page_size"`
}

type DatasourceListRes struct {
	List  []*entity.Datasource `json:"list"`
	Total int `json:"total"`
}

type DatasourceListAllReq struct {
	g.Meta   `path:"/datasource/listAll" method:"get" tags:"Datasource" summary:"数据源列表"`
}

type DatasourceListAllRes struct {
	List  []*entity.Datasource `json:"list"`
}

type DatasourceCreateReq struct {
	g.Meta   `path:"/datasource/create" method:"post" tags:"Datasource" summary:"数据源创建"`
	Name     string `v:"required" p:"name"`
	Host     string `v:"required" p:"host"`
	Port     string `v:"required" p:"port"`
	UserName string `v:"required" p:"username"`
	Password string `v:"required" p:"password"`
}

type DatasourceCreateRes struct {
}

type DatasourceEditReq struct {
	g.Meta   `path:"/datasource/edit" method:"post" tags:"Datasource" summary:"数据源修改"`
	Id       int64  `v:"required" p:"id"`
	Name     string `v:"required" p:"name"`
	Host     string `v:"required" p:"host"`
	Port     string `v:"required" p:"port"`
	UserName string `v:"required" p:"username"`
	Password string `v:"required" p:"password"`
}

type DatasourceEditRes struct {
}

type DatasourceDeleteReq struct {
	g.Meta   `path:"/datasource/delete" method:"get" tags:"Datasource" summary:"数据源删除"`
	Id       int64  `v:"required" p:"id"`
}

type DatasourceDeleteRes struct {
}

type DatasourceDetailReq struct {
	g.Meta   `path:"/datasource/detail" method:"get" tags:"Datasource" summary:"数据源删除"`
	Id       int64  `v:"required" p:"id"`
}

type DatasourceDetailRes struct {
	*entity.Datasource
}
