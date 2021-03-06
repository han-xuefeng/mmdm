// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mmdm/internal/service/internal/dao/internal"
)

// datasourceDao is the data access object for table mmdm_datasource.
// You can define custom methods on it to extend its functionality as you wish.
type datasourceDao struct {
	*internal.DatasourceDao
}

var (
	// Datasource is globally public accessible object for table mmdm_datasource operations.
	Datasource = datasourceDao{
		internal.NewDatasourceDao(),
	}
)

// Fill with you ideas below.
