package model

type DatagroupCreateInput struct {
	Name string
	SourceId int64
	Prefix string
	GroupConfig string
}

type DatagroupListInput struct {
	Page int
	PageSize int
}

type DatagroupEditInput struct {
	Id int64
	Name string
	SourceId int64
	Prefix string
	GroupConfig string
}

type DatagroupDetailInput struct {
	Id int64
}
