package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mmdm/internal/model/entity"
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
	*entity.Admin
}

