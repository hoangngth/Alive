package database

import (
	util "alive/Alive/ToDoGRPC/util"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	dsn := util.GetDatabaseString("3My3QTb653", "FGrwr22PUt", "remotemysql.com", 3306, "3My3QTb653")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open database: %v", err)
	}
	return db
}
