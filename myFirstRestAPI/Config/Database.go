package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// this is for the configuration of the database which tells us which database will be used and configured
var DB *gorm.DB // this is the gorm variable

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

func BuildDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "ayush123",
		DBName:   "api_db",
	}
}

func DBURL(db *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.DBName)
}
