package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" method:"post" tags:"Admin" summary:"Login in with exist account"`
	UserName string `v:"required" p:"username"`
	Password string `v:"required" p:"password"`
}

type AdminLoginRes struct{
	Token string `json:"token"`
}

type AdminLogoutReq struct {
	g.Meta `path:"/admin/logout" method:"get" tags:"Admin"`
}

type AdminLoutRes struct {}

type AdminInfoReq struct {
	g.Meta   `path:"/admin/info" method:"get" tags:"Admin" summary:"admin info"`
}
type AdminInfoRes struct {
	ID           int64       `json:"id"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type AdminChangePwdReq struct {
	g.Meta   `path:"/admin/change_pwd" method:"post" tags:"Admin" summary:"change password"`
	Password string `v:"required" p:"password"`
}

type AdminChangePwdRes struct {

}

