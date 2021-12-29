// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mmdm/internal/service/internal/dao/internal"
)

// sqlTsakDao is the data access object for table mmdm_sql_tsak.
// You can define custom methods on it to extend its functionality as you wish.
type sqlTsakDao struct {
	*internal.SqlTsakDao
}

var (
	// SqlTsak is globally public accessible object for table mmdm_sql_tsak operations.
	SqlTsak = sqlTsakDao{
		internal.NewSqlTsakDao(),
	}
)

// Fill with you ideas below.
