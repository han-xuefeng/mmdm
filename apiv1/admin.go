package apiv1

import "github.com/gogf/gf/v2/frame/g"

type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" method:"post" tags:"Admin" summary:"Login in with exist account"`
	UserName string `v:"required" p:"username"`
	Password string `v:"required" p:"password"`
}

type AdminLoginRes struct{
	Token string `json:"token"`
}

