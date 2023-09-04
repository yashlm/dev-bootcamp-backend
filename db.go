package main

//
//import (
//	"database/sql"
//	"log"
//)
//
//func InitDB() {
//	connStr := "host=localhost user=postgres dbname=stock_data port=5432 sslmode=disable"
//	database, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = database
//}
//
//func InsertDataIntoDB(data StockData) {
//	// Prepare the INSERT INTO SQL statement
//	stmt, err := db.Prepare("INSERT INTO stock_data (symbol, date, open, high, low, close, volume) VALUES ($1, $2, $3, $4, $5, $6, $7)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//
//	// Execute the SQL statement with data from the API response
//	_, err = stmt.Exec(data.Symbol, data.Date, data.Open, data.High, data.Low, data.Close, data.Volume)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
