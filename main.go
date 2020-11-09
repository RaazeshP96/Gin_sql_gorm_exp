package main

import (
	"fmt"

	"github.com/RaazeshP96/Gin_sql_gorm_exp/Config"
	"github.com/RaazeshP96/Gin_sql_gorm_exp/Models"
	"github.com/RaazeshP96/Gin_sql_gorm_exp/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
