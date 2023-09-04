package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"gopkg.in/resty.v1"
//	"log"
//	"net/http"
//	"time"
//)
//
//func FetchAndStoreStockData() {
//	// Make an HTTP GET request to the API to fetch NIFTY50 stock data
//	apiUrl := "https://api.example.com/nifty50-stock-data" // Replace with the actual API URL
//	response, err := resty.New().R().Get(apiUrl)           // Create a new Resty instance
//	if err != nil {
//		log.Println("Failed to fetch stock data:", err)
//		return
//	}
//
//	// Close the response body when done
//	defer func() {
//		if cerr := response.RawResponse.Body.Close(); cerr != nil {
//			log.Println("Error closing response body:", cerr)
//		}
//	}()
//
//	// Check if the response status code is not 200 OK
//	if response.StatusCode() != http.StatusOK {
//		log.Println("API returned a non-OK status code:", response.Status())
//		return
//	}
//
//	// Read the response body into a byte slice
//	responseBody := response.Body()
//
//	// Create an io.Reader from the byte slice
//	responseReader := bytes.NewBuffer(responseBody)
//
//	// Parse the API response into your StockData struct
//	var stockData StockData
//	if err := json.NewDecoder(responseReader).Decode(&stockData); err != nil {
//		log.Println("Failed to parse API response:", err)
//		return
//	}
//
//	// Insert the data into the database
//	InsertDataIntoDB(stockData)
//}
//
//func ScheduleDataUpdates() {
//	// Use a ticker to schedule periodic updates (e.g., every minute)
//	ticker := time.NewTicker(1 * time.Minute)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			// Fetch and store stock data at regular intervals
//			FetchAndStoreStockData()
//		}
//	}
//}
//
//func GetStockData(c *gin.Context) {
//	symbol := c.Param("symbol")
//	date := c.Param("date")
//
//	// Query the database to fetch stock data for a specific symbol and date.
//	// Implement proper error handling here.
//	rows, err := db.Query("SELECT * FROM stock_data WHERE symbol=$1 AND date=$2", symbol, date)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer rows.Close()
//
//	// Process the rows and return JSON response.
//	// Implement proper data processing and error handling here.
//	var stockData StockData
//	if rows.Next() {
//		err := rows.Scan(
//			&stockData.Symbol,
//			&stockData.Date,
//			&stockData.Open,
//			&stockData.High,
//			&stockData.Low,
//			&stockData.Close,
//			&stockData.Volume,
//		)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, stockData)
//	} else {
//		c.JSON(http.StatusNotFound, gin.H{"message": "Stock data not found"})
//	}
//}
//
//func GetLatestStockData(c *gin.Context) {
//	symbol := c.Param("symbol")
//
//	// Query the database to fetch the latest stock data for a specific symbol.
//	// Implement proper error handling here.
//	rows, err := db.Query("SELECT * FROM stock_data WHERE symbol=$1 ORDER BY date DESC LIMIT 1", symbol)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer rows.Close()
//
//	// Process the rows and return JSON response.
//	// Implement proper data processing and error handling here.
//	var stockData StockData
//	if rows.Next() {
//		err := rows.Scan(
//			&stockData.Symbol,
//			&stockData.Date,
//			&stockData.Open,
//			&stockData.High,
//			&stockData.Low,
//			&stockData.Close,
//			&stockData.Volume,
//		)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, stockData)
//	} else {
//		c.JSON(http.StatusNotFound, gin.H{"message": "Stock data not found"})
//	}
//}
//
//func GetAllStockData(c *gin.Context) {
//	// Query the database to fetch all stock data.
//	// Implement proper error handling here.
//	rows, err := db.Query("SELECT * FROM stock_data")
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer rows.Close()
//
//	// Process the rows and return JSON response.
//	// Implement proper data processing and error handling here.
//	var stockData []StockData
//	for rows.Next() {
//		var data StockData
//		err := rows.Scan(
//			&data.Symbol,
//			&data.Date,
//			&data.Open,
//			&data.High,
//			&data.Low,
//			&data.Close,
//			&data.Volume,
//		)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		stockData = append(stockData, data)
//	}
//	c.JSON(http.StatusOK, stockData)
//}
