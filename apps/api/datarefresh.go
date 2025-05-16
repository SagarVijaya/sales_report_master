package api

import (
	"encoding/csv"
	"net/http"
	"os"
	"salesproject/apps/models"
	"salesproject/apps/utils"
	"salesproject/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func DataRefreshAPI(r *gin.Context) {
	log := new(utils.LoggerId)
	log.SetSid(r.Request)
	log.Log("DataRefreshAPI (+)")
	defer log.Log("DataRefreshAPI (-)")
	err := DataFromCSV("csv/sample.csv")
	if err != nil {
		log.Log("DataReadError", err)
		r.JSON(http.StatusInternalServerError, gin.H{
			"status":  "E",
			"message": "Data Refresh error",
		})
		return
	}

	response := map[string]string{
		"status":  "S",
		"message": "Data refresh triggered successfully",
	}

	r.Writer.Header().Set("Content-Type", "application/json")
	r.JSON(http.StatusOK, response)

}

func DataFromCSV(filePath string) error {
	log := new(utils.LoggerId)
	log.SetRef("DataFromCSV")
	log.Log("DataFromCSV (+)")
	defer log.Log("DataFromCSV (-)")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	var lCustomer []models.Customer
	var lProduct []models.Product
	var lOrder []models.Order
	var lOrderItem []models.OrderDetails

	for _, row := range rows[1:] {
		orderID := row[0]
		productID := row[1]
		customerID := row[2]

		quantity, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)
		dateOfSale, _ := time.Parse("2006-01-02", row[6])

		lProductInfo := models.Product{
			ProductID: productID,
			Name:      row[3],
			Category:  row[4],
		}
		lOrderInfo := models.Order{
			OrderID:      orderID,
			CustomerID:   customerID,
			Region:       row[5],
			DateOfSale:   dateOfSale,
			PaymentType:  row[11],
			ShippingCost: shippingCost,
		}
		lCustomerInfo := models.Customer{
			CustomerID: customerID,
			Name:       row[12],
			Email:      row[13],
			Address:    row[14],
		}

		lOrderItemInfo := models.OrderDetails{
			OrderID:      orderID,
			ProductID:    productID,
			QuantitySold: quantity,
			UnitPrice:    unitPrice,
			Discount:     discount,
		}
		lProduct = append(lProduct, lProductInfo)
		lCustomer = append(lCustomer, lCustomerInfo)
		lOrderItem = append(lOrderItem, lOrderItemInfo)
		lOrder = append(lOrder, lOrderInfo)
	}
	lErr := database.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&lProduct).Error
	if lErr != nil {
		return lErr
	}
	lErr1 := database.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&lCustomer).Error
	if lErr1 != nil {
		return lErr1
	}
	lErr2 := database.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&lOrder).Error
	if lErr2 != nil {
		return lErr2
	}
	lErr3 := database.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&lOrderItem).Error
	if lErr3 != nil {
		return lErr3
	}

	return nil
}
