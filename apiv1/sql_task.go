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
	List  []*entity.SqlTask
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


type SqlTaskEditReq struct {
	g.Meta   `path:"/sql_task/edit" method:"post" tags:"SqlTask" summary:"预执行的任务"`
	Id int64 `v:"required" p:"id"`
	Name string  `v:"required" p:"name"`
	Desc string  `v:"required" p:"desc"`
	GroupId int64  `v:"required" p:"group_id"`
	Sql string  `v:"required" p:"sql"`
}

type SqlTaskEditRes struct {}

type SqlTaskDetailReq struct {
	g.Meta   `path:"/sql_task/detail" method:"get" tags:"SqlTask" summary:"预执行的任务详情"`
	Id int64 `v:"required" p:"id"`
}

type SqlTaskDetailRes struct {
	*entity.SqlTask
}

type SqlTaskExecReq struct {
	g.Meta   `path:"/sql_task/exec" method:"get" tags:"SqlTask" summary:"预执行的任务详情"`
	Id int64 `v:"required" p:"id"`
}

type SqlTaskExecRes struct {

}

type SqlTaskDeleteReq struct {
	g.Meta   `path:"/sql_task/delete" method:"get" tags:"SqlTask" summary:"预执行的任务详情"`
	Id int64 `v:"required" p:"id"`
}

type SqlTaskDeleteRes struct {

}