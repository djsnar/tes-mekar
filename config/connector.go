package config

import (
	"database/sql"
	"fmt"
	"log"
	"mekar/utils"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

	dbUser := utils.ReadEnv("dbUser", "root")
	dbPass := utils.ReadEnv("dbPass", "password")
	dbHost := utils.ReadEnv("dbHost", "localhost")
	dbPort := utils.ReadEnv("dbPort", "3306")
	dbName := utils.ReadEnv("dbName", "schema")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	//fmt.Println("Environment has been activated!")

	return db, nil
}
