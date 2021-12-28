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

type DatasourceEditInput struct {
	Id int64
	Name string
	Passport string
	Port string
	Host string
	UserName string
}

type DatasourceDetailInput struct {
	Id int64
}
