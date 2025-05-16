package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"salesproject/apps/api"
	"salesproject/apps/utils"
	"salesproject/config"
	"salesproject/database"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create the log directory if it doesn't exist
	err := os.MkdirAll("log", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	logFolderName := "./log/log" + time.Now().Format("02012006.15.04.05") + ".log"
	file, err := os.OpenFile(logFolderName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)

	err = database.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	go router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "This page is not available,Please check the url",
		})
	})
	go Schedular()
	router.GET("/trigger", api.DataRefreshAPI)
	routerGroup := router.Group("/topProduct")
	routerGroup.GET("/categoryData", api.GetCategoryData)
	routerGroup.GET("/overallData", api.GetTopProductOverAll)
	routerGroup.GET("/regionData", api.GetRegionData)
	port := config.GetConfig().Service.Port
	log.Println("Server running on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func init() {
	config.LoadGlobalConfig("toml/config.toml")
}

func Schedular() {
	log := new(utils.LoggerId)

	var runScheduler any
	runScheduler = func() {
		log.Log("Running Scheduler For Data Refresh")

		// Run the data refresh operation
		err := api.DataFromCSV("csv/sample.csv")
		if err != nil {
			log.Log("Scheduler error:", err.Error())
		}
		_, lOk := runScheduler.(func())
		if lOk {
			// After completing, reschedule the task to run again after 24 hours
			time.AfterFunc(24*time.Hour, runScheduler.(func()))
			log.Log("Scheduler Triggered")
		}
	}

	_, lOk := runScheduler.(func())
	if lOk {
		// For Initial Trigger
		time.AfterFunc(0, runScheduler.(func()))
	}
}
