package handlers

import (
	"net/http"
	"food-ordering-system-project/services"
	"github.com/gin-gonic/gin"
	"food-ordering-system-project/helpers"
)


type AdminHandler struct {
	AdminService *services.AdminService
}


// func AdminDashboard(c *gin.Context) {
func (h *AdminHandler) AdminDashboard(c *gin.Context) {
	stats, err := h.AdminService.GetDashboardStats()
    
	if err != nil {
	c.String(500, err.Error())
	return
    }

	user := helpers.GetSessionData(c)

	

	c.HTML(http.StatusOK, "ad_dashboard.html", gin.H{

	  	"title":          "Dashboard",
		"role": user.Role,
		"username": user.Username,
		"name":       user.Name,
		"Surname":       user.Surname,
		"totalOrders":    stats.TotalOrders,
		"pendingOrders":  stats.PendingOrders,
		"closedOrders":   stats.ClosedOrders,
		"rejectedOrders": stats.RejectedOrders,
		"todayOrders":    stats.TodayOrders,
		"todayPending":   stats.TodayPending,
		"todayClosed":    stats.TodayClosed,
		"todayRejected":  stats.TodayRejected,
		"totalSales": stats.TotalSales,
        "todaySales": stats.TodaySales,
		"totalCancelledSales": stats.TotalCancelledSales,
        "todayCancelledSales": stats.TodayCancelledSales,
	})
	
}

// func AdminMenuPage(c *gin.Context) {
func (h *AdminHandler) AdminMenuPage(c *gin.Context) {

	
    // menus, err := repository.GetMenus()
	menus, err := h.AdminService.GetMenus()

	if err != nil {
		c.String(500, err.Error())
		return
	}

	
    

	user := helpers.GetSessionData(c)

	c.HTML(http.StatusOK, "ad_menu.html", gin.H{
	  	"title": "Menu",
		"menus": menus,
		"role": user.Role,
		"username": user.Username,
		"name":       user.Name,
		"Surname":       user.Surname,
	})
}

