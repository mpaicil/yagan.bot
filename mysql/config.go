package mysql

import (
	// "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"yagan.bot/config"
)

var OpenDB *sql.DB

func Connection() {
	db, err := sql.Open("mysql", config.DBProperty())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	OpenDB = db
	// test()
}

type Ticker struct {
	id int `json: id`
	min_ask string `json:"min_ask"`
	max_bid string `json:"max_bid"`
}

// func test() {
// 	//Execute the query
// 	results, err := OpenDB.Query(config.RetrieveQuery("tickers"))
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}

// 	for results.Next() {
// 		var ticker Ticker

// 		err = results.Scan(&ticker.id, &ticker.min_ask, &ticker.max_bid)
// 		if err != nil {
// 			panic(err.Error()) // proper error handling instead of panic in your app
// 		}

// 		//fmt.Println(ticker.min_ask)
// 	}
// }
