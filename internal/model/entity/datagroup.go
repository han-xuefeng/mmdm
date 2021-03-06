// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Datagroup is the golang structure for table datagroup.
type Datagroup struct {
	Id          int64       `json:"id"          description:"自增id"`
	SourceId    int64       `json:"sourceId"    description:"数据源id"`
	Name        string      `json:"name"        description:"组名"`
	Prefix      string      `json:"prefix"      description:"组前缀规则"`
	GroupConfig string      `json:"groupConfig" description:"组命令执行前的sql"`
	CreateAt    *gtime.Time `json:"createAt"    description:"新增时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    description:"更新时间"`
	IsDelete    int         `json:"isDelete"    description:"是否删除"`
}
