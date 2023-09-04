package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type StockData struct {
	Symbol string    `json:"symbol"`
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}

func main() {
	// Initialize the database connection.
	initDB()

	// Create a router using Gin.
	r := gin.Default()

	// Define API routes.
	r.GET("/stock/:symbol/:date", getStockData)
	r.GET("/stock/:symbol/latest", getLatestStockData)
	r.GET("/stocks", getAllStockData)

	// Start the server.
	r.Run(":8080")
}

func initDB() {
	// Replace with your PostgreSQL database connection details.
	connStr := "user=username password=postgrres dbname=dbname sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func getStockData(c *gin.Context) {
	symbol := c.Param("symbol")
	date := c.Param("date")

	// Query the database to fetch stock data for a specific symbol and date.
	rows, err := db.Query("SELECT * FROM stock_data WHERE symbol=$1 AND date=$2", symbol, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Process the rows and return JSON response.
	var stockData StockData
	if rows.Next() {
		err := rows.Scan(
			&stockData.Symbol,
			&stockData.Date,
			&stockData.Open,
			&stockData.High,
			&stockData.Low,
			&stockData.Close,
			&stockData.Volume,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stockData)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Stock data not found"})
	}
}

func getLatestStockData(c *gin.Context) {
	symbol := c.Param("symbol")

	// Query the database to fetch the latest stock data for a specific symbol.
	rows, err := db.Query("SELECT * FROM stock_data WHERE symbol=$1 ORDER BY date DESC LIMIT 1", symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Process the rows and return JSON response.
	var stockData StockData
	if rows.Next() {
		err := rows.Scan(
			&stockData.Symbol,
			&stockData.Date,
			&stockData.Open,
			&stockData.High,
			&stockData.Low,
			&stockData.Close,
			&stockData.Volume,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stockData)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Stock data not found"})
	}
}

func getAllStockData(c *gin.Context) {
	// Query the database to fetch all stock data.
	rows, err := db.Query("SELECT * FROM stock_data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Process the rows and return JSON response.
	var stockData []StockData
	for rows.Next() {
		var data StockData
		err := rows.Scan(
			&data.Symbol,
			&data.Date,
			&data.Open,
			&data.High,
			&data.Low,
			&data.Close,
			&data.Volume,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stockData = append(stockData, data)
	}
	c.JSON(http.StatusOK, stockData)
}

//
//func fetchAndStoreStockData() {
//	// Make an HTTP GET request to the API to fetch NIFTY50 stock data
//	apiUrl := "https://api.example.com/nifty50-stock-data" // Replace with the actual API URL
//	response, err := resty.New().R().Get(apiUrl)           // Create a new Resty instance
//	if err != nil {
//		log.Println("Failed to fetch stock data:", err)
//		return
//	}
//	defer response.RawResponse.Body.Close() // Close the response body when done
//
//	// Check if the response status code is not 200 OK
//	if response.StatusCode() != http.StatusOK {
//		log.Println("API returned a non-OK status code:", response.Status())
//		return
//	}
//
//	// Create an io.Reader from the response body
//	responseReader := response.Body()
//
//	// Parse the API response into your StockData struct
//	var stockData StockData
//	if err := json.NewDecoder(responseReader).Decode(&stockData); err != nil {
//		log.Println("Failed to parse API response:", err)
//		return
//	}
//
//	// Insert the data into the database
//	insertDataIntoDB(stockData)
//}
//
//func scheduleDataUpdates() {
//	// Use a ticker to schedule periodic updates (e.g., every minute)
//	ticker := time.NewTicker(1 * time.Minute)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			// Fetch and store stock data at regular intervals
//			fetchAndStoreStockData()
//		}
//	}
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
