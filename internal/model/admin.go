package model

type AdminLoginInput struct {
	UserName string
	Password string
}

type AdminChangePwd struct {
	Password string
}
