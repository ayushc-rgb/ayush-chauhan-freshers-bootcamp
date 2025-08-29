package main

import (
	"firstAPI/Config"
	"firstAPI/Models"
	"firstAPI/Routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DBURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status of Config: ", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetUpRouter()
	r.Run()
}
