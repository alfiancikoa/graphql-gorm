package dbconf

import (
	"fmt"

	"github.com/alfiancikoa/graphql-gorm/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbconnection := "root:root@tcp(172.17.0.2:3306)/?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dbconnection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("DATABASE IS CONNECTED")
	}
	DB.LogMode(true)
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	DB.Exec("CREATE DATABASE IF NOT EXISTS db_graphql_gorm ")
	DB.Exec("USE db_graphql_gorm")

	// Migrate the schema
	// Perintah untuk membuat tabel secara otomatis pada database
	DB.AutoMigrate(&model.Movie{})
	DB.AutoMigrate(&model.Star{})
}
