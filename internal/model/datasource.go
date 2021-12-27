package model

type DatasourceCreateInput struct {
	Name string
	Passport string
	Port string
	Host string
	UserName string
}

type DatasourceListInput struct {
	Page int
	PageSize int
}
