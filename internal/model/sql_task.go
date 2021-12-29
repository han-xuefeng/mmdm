package model

type SqlTaskListInput struct {
	Page int
	PageSize int
}


type SqlTaskCreateInput struct {
	Name string
	Desc string
	GroupId int64
	Sql string
}