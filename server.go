package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alfiancikoa/graphql-gorm/graph"
	"github.com/alfiancikoa/graphql-gorm/graph/generated"
	"github.com/alfiancikoa/graphql-gorm/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	dbconnect := "root:root@tcp(172.17.0.2:3306)/?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dbconnect)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("DATABASE IS CONNECTED")
	}
	db.LogMode(true)
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE IF NOT EXISTS db_graphql_gorm")
	db.Exec("USE db_graphql_gorm")

	// Migrate the schema
	// Perintah untuk membuat tabel secara otomatis pada database
	db.AutoMigrate(&model.Star{})
	db.AutoMigrate(&model.Movie{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
