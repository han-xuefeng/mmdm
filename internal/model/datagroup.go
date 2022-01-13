package model

import "mmdm/internal/model/entity"

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

type DatagroupListItemOutput struct {
	*entity.Datagroup
	SourceName string `json:"sourceName"    description:"数据源名称"`
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
