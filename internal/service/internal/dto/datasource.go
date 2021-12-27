// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DatasourceForDao is the golang structure of table mmdm_datasource for DAO operations like Where/Data.
type Datasource struct {
	g.Meta   `orm:"dto:true"`
	Id       interface{} // 自增id
	Name     interface{} //
	Host     interface{} // host
	UserName interface{} // 用户名
	Port     interface{} // 端口
	Password interface{} // 密码
	CreateAt *gtime.Time // 新增时间
	UpdateAt *gtime.Time // 更新时间
	IsDelete interface{} // 是否删除
}