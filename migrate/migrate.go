package main

import (
	"fmt"

	"github.com/1206yaya/go-echo-jwt-noteapp-api/db"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Note{})
}
