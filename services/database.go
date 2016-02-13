package services

import (
	. "app/common"
	_ "database/sql"
	"fmt"
	conf "github.com/andboson/configlog"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

func InitDb() {
	var err error
	config := conf.AppConfig
	database, _ := config.String("database.pg_base")
	user, _ := config.String("database.pg_user")
	password, _ := config.String("database.pg_pass")
	host, _ := config.String("database.pg_host")
	port, _ := config.String("database.db_port")
	sslmode, _ := config.String("database.sslmode")
	sslcert, _ := config.String("database.sslcert")
	sslkey, _ := config.String("database.sslkey")
	sslrootcert, _ := config.String("database.sslrootcert")
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s",
		host, port, user, password, database, sslmode, sslcert, sslkey, sslrootcert))
	db.DB()
	err = db.DB().Ping()
	if err != nil {
		Log.WithField("errod", err).Fatalf("\n Cannot connect to database")
	}

	DB = db
}
