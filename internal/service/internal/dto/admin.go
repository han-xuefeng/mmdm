// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminForDao is the golang structure of table mmdm_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta   `orm:"dto:true"`
	Id       interface{} // 自增id
	UserName interface{} // 用户名
	Salt     interface{} // 盐
	Password interface{} // 密码
	CreateAt *gtime.Time // 新增时间
	UpdateAt *gtime.Time // 更新时间
	IsDelete interface{} // 是否删除
}
