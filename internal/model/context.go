package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session *ghttp.Session // Session in context.
	Admin    *ContextAdmin   // User in context.
}

type ContextAdmin struct {
	Id       int64   // User ID.
	Passport string // User passport.
	UserName string // User nickname.
}
