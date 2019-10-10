package sqldb

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	mariaDriver = "mysql"
)

//MariaConfig maria db configuration
type MariaConfig struct {
	DBName, Host, User, Pass string
	Port                     int
}

func (mc *MariaConfig) build() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		mc.User,
		mc.Pass,
		mc.Host,
		mc.Port,
		mc.DBName,
		"UTF8",
	)
}

//NewMaria creates new maria db client
func NewMaria(conf *MariaConfig) (*sqlx.DB, error) {
	cs := conf.build()
	return sqlx.Open(mariaDriver, cs)
}
