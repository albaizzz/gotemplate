package sqldb

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	mssqlDriver = "mssql"
)

//Mssql db configuration
type MssqlConfig struct {
	DBName, Host, User, Pass string
	Port                     int
}

func (mc *MssqlConfig) build() string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s",
		mc.User,
		mc.Pass,
		mc.Host,
		mc.Port,
		mc.DBName,
	)
}

//NewMssql create new connection sqlserver client
func NewMssql(conf *MssqlConfig) (*sqlx.DB, error) {
	cs := conf.build()
	return sqlx.Open(mssqlDriver, cs)
}
