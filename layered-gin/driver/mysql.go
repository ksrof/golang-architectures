// layered-gin/driver/mysql.go

package driver

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB() starts the database connection
func InitDB(conf DBConfig) (*gorm.DB, error) {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to the database")
		return nil, err
	}

	return db, nil
}