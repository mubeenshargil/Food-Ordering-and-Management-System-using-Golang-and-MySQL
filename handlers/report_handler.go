package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"food-ordering-system-project/helpers"
	"food-ordering-system-project/services"

)

type ReportHandler struct {
	ReportService *services.ReportService
}

// func SalesReport(c *gin.Context) {
func (h *ReportHandler) SalesReport(c *gin.Context) {

	fromDate := c.Query("from")
	toDate := c.Query("to")

	// report, err := repository.GetSalesReport(
	report, err := h.ReportService.GetSalesReport(
		fromDate,
		toDate,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}
  
	user := helpers.GetSessionData(c)

	c.HTML(
		http.StatusOK,
		"sales_report.html",
		gin.H{
			"title":  "Sales Report",
			"report": report,
			"from":   fromDate,
			"to":     toDate,
		    "role": user.Role,
			"username": user.Username,
			"name":       user.Name,
			"Surname":       user.Surname,
		},
	)
}

// func MenuPerformanceReport(c *gin.Context) {
func (h *ReportHandler) MenuPerformanceReport(c *gin.Context) {

	// report, err := repository.GetMenuPerformanceReport()
	report, err := h.ReportService.GetMenuPerformanceReport()

	if err != nil {
		c.String(500, err.Error())
		return
	}
    
	user := helpers.GetSessionData(c)
	
	c.HTML(
		http.StatusOK,
		"menu_performance_report.html",
		gin.H{
			"title":  "Menu Performance Report",
			"report": report,
			"role": user.Role,
			"username": user.Username,
			"name":       user.Name,
			"Surname":       user.Surname,
		},
	)
}

// func CustomerOrdersReport(c *gin.Context) {
func (h *ReportHandler) CustomerOrdersReport(c *gin.Context) {

	// report, err := repository.GetCustomerOrdersReport()
	report, err := h.ReportService.GetCustomerOrdersReport()

	if err != nil {
		c.String(500, err.Error())
		return
	}

	user := helpers.GetSessionData(c)

	c.HTML(
		http.StatusOK,
		"customer_orders_report.html",
		gin.H{
			"title":   "Customer Orders Report",
			"report":  report,
			"role":    user.Role,
			"name":    user.Name,
			"surname": user.Surname,
		},
	)
}

// func OrdersReport(c *gin.Context) {
func (h *ReportHandler) OrdersReport(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	status := c.Query("status")

	// report, err := repository.GetOrdersReport(
	report, err := h.ReportService.GetOrdersReport(

		startDate,
		endDate,
		status,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	user := helpers.GetSessionData(c)

	c.HTML(
		http.StatusOK,
		"orders_report.html",
		gin.H{
			"title":     "Orders Report",
			"report":    report,
			"startDate": startDate,
			"endDate":   endDate,
			"status":    status,
			"role":      user.Role,
			"name":      user.Name,
			"surname":   user.Surname,
		},
	)
}