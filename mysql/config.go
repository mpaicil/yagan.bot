package mysql

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"yagan.bot/config"
)

var OpenDB *sql.DB

func Connection() {
	fmt.Println("Initialize mysql connection ...")
	db, err := sql.Open("mysql", config.DBProperty()+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	OpenDB = db

	RetrieveTickers()
}