package mysql

import(
	"yagan.bot/config"
	"time"
	"container/list"
)

type Ticker struct {
	Id int `db:"id"`
	Min_ask string `db:"min_ask"`
	Max_bid string `db:"max_bid"`
	Date_time *time.Time `db:"date_time"`
}

func RetrieveTickers() *list.List{
	results, err := OpenDB.Query(config.RetrieveQuery("tickers"))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tickers := list.New()

	for results.Next() {
		var ticker Ticker

		err = results.Scan(&ticker.Id, &ticker.Min_ask, &ticker.Max_bid, &ticker.Date_time)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		tickers.PushBack(ticker)
	}

	return tickers
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