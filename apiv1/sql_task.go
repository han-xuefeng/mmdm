package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mmdm/internal/model/entity"
)

type SqlTaskListReq struct {
	g.Meta   `path:"/sql_task/list" method:"get" tags:"SqlTask" summary:"执行的任务列表"`
	Page     int `v:"required" p:"page"`
	PageSize int `v:"required" p:"page_size"`
}

type SqlTaskListRes struct {
	List  []*entity.SqlTsak
	Total int
}

type SqlTaskCreateReq struct {
	g.Meta   `path:"/sql_task/create" method:"post" tags:"SqlTask" summary:"执行的任务"`
	Name string  `v:"required" p:"name"`
	Desc string  `v:"required" p:"desc"`
	GroupId int64  `v:"required" p:"group_id"`
	Sql string  `v:"required" p:"sql"`
}

type SqlTaskCreateRes struct {

}

type SqlTaskPreCreateReq struct {
	g.Meta   `path:"/sql_task/pre_create" method:"post" tags:"SqlTask" summary:"预执行的任务"`
	Name string  `v:"required" p:"name"`
	Desc string  `v:"required" p:"desc"`
	GroupId int64  `v:"required" p:"group_id"`
	Sql string  `v:"required" p:"sql"`
}

type SqlTaskPreCreateRes struct {

}