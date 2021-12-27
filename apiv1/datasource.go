package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mmdm/internal/model/entity"
)

type DatasourceListReq struct {
	g.Meta   `path:"/datasource/list" method:"get" tags:"Datasource" summary:"数据源列表"`
	Page int `v:"required" p:"page"`
	PageSize int `v:"required" p:"page_size"`
}

type DatasourceListRes struct {
	List []*entity.Datasource
	Total int
}


type DatasourceCreateReq struct {
	g.Meta `path:"/datasource/create" method:"post" tags:"Datasource" summary:"数据源创建"`
	Name string `v:"required" p:"name"`
	Host string `v:"required" p:"host"`
	Port string `v:"required" p:"port"`
	UserName string `v:"required" p:"username"`
	Password string `v:"required" p:"password"`

}

type DatasourceCreateRes struct {

}