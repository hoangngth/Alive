package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	util "Alive/ToDoGRPCv2/util"
)

func ConnectDatabaseToDo() *sql.DB {
	dsn := util.GetDatabaseString("3My3QTb653", "FGrwr22PUt", "remotemysql.com", 3306, "3My3QTb653")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open database: %v", err)
	}
	return db
}

func ConnectDatabaseProfile() *sql.DB {
	dsn := util.GetDatabaseString("sql12341781", "VXT6f6URGw", "sql12.freemysqlhosting.net", 3306, "sql12341781")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open database: %v", err)
	}
	return db
}
