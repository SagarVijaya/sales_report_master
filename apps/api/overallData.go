package api

import (
	"net/http"
	"salesproject/apps/utils"
	"salesproject/database"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTopProductOverAll(r *gin.Context) {
	log := new(utils.LoggerId)
	log.SetSid(r.Request)
	log.Log("GetTopProductOverAll (+)")
	defer log.Log("GetTopProductOverAll (-)")

	var lResponse float64

	StartDateStr := r.Request.URL.Query().Get("start_date")
	EndDateStr := r.Request.URL.Query().Get("end_date")

	if StartDateStr == "" || EndDateStr == "" {
		log.Log("start date and end date are required")
		r.JSON(http.StatusBadRequest, gin.H{
			"status": "E",
			"errmsg": "Invalid start date and end data is required",
		})
		return
	}

	StartDate, err := time.Parse("2006-01-02", StartDateStr)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"status": "E",
			"errmsg": "Invalid start date",
		})
		return
	}
	EndDate, err := time.Parse("2006-01-02", EndDateStr)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"status": "E",
			"errmsg": "Invalid end date",
		})
		return
	}

	lErr := database.DB.Table("order_details od").Select("sum(od.quantity_sold) as quantity").Joins("join orders o on o.order_id =od.order_id").
		Where("o.date_of_sale between ? and ?", StartDate, EndDate).Order("od.quantity_sold desc").Find(&lResponse)
	if lErr.Error != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"status": "E",
			"errmsg": "Data Fetch Error",
		})
		return
	}

	response := map[string]interface{}{
		"status":  "S",
		"message": "Overall Sold Product",
		"data": map[string]interface{}{
			"start":                 StartDate.Format("2006-01-02"),
			"end":                   EndDate.Format("2006-01-02"),
			"overall_sold_quantity": lResponse,
		},
	}
	r.Writer.Header().Set("Content-Type", "application/json")
	r.JSON(http.StatusOK, response)

}
