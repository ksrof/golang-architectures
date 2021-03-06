// layered-architecture/driver/mysql.go

package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Host 			string
	User 			string
	Password 	string
	Port 			string
	DB 				string
}

// MySQLConnect takes mysql config, forms the connection string and connects to mysql.
func MySQLConnect(conf MySQLConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}