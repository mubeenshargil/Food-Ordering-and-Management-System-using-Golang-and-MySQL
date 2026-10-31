package handlers

import (
	"food-ordering-system-project/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"food-ordering-system-project/helpers"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

// func CheckoutPage(c *gin.Context) {
func (h *OrderHandler) CheckoutPage(c *gin.Context) {

	menuID := c.Param("id")

	// menu, err := repository.GetMenuByID(menuID)
	menu, err := h.OrderService.GetMenuByID(menuID)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.HTML(http.StatusOK, "checkout.html", gin.H{
		"menu": menu,
	})
}

// func PlaceOrder(c *gin.Context) {
func (h *OrderHandler) PlaceOrder(c *gin.Context) {

	customerName := c.PostForm("customer_name")
	customerSurname := c.PostForm("customer_surname")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	email := c.PostForm("email")
    status := "PENDING"
	menuID := c.PostForm("menu_id")
	quantity := c.PostForm("quantity")
	price := c.PostForm("price")
	total := c.PostForm("total")


	// orderID, err := repository.CreateOrder(
	orderID, err := h.OrderService.CreateOrder(
	customerName,
	customerSurname,
	phone,
	address,
	email,
	menuID,
	quantity,
	status,
	price,
	total,
)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	err = h.OrderService.CreateOrderHistory(
		orderID,
		status,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}


	c.HTML(http.StatusOK, "success.html", gin.H{
	"name": customerName,
	"orderID": orderID,
})
}

// func AdminOrders(c *gin.Context) {
func (h *OrderHandler) AdminOrders(c *gin.Context) {

//    orders, err := repository.GetAllOrders()
    orders, err := h.OrderService.GetAllOrders()

	if err != nil {
		c.String(500, err.Error())
		return
	}

	user := helpers.GetSessionData(c)

	c.HTML(http.StatusOK, "ad_orders.html", gin.H{
		"title": "Orders",
		"orders": orders,
		"role": user.Role,
		"username": user.Username,
		"name":       user.Name,
		"Surname":       user.Surname,
	})
}

// func UpdateOrderStatus(c *gin.Context) {
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
    
	user := helpers.GetSessionData(c)

	id := c.Param("id")
  
	status := c.PostForm("status")
	reason := c.PostForm("cancel_reason")


	if status == "CANCELLED" && reason == "" {
	c.String(400, "Cancellation reason is required")
	return
   }

		// err := repository.UpdateOrderStatus(
		err := h.OrderService.UpdateOrderStatus(
			id,
			status,
			user.UserID,
			reason,
		)

	if err != nil {

		c.String(500, err.Error())
		return
	}
      


	if user.Role == "admin" {

		c.Redirect(http.StatusSeeOther,"/admin/orders",)

		} else if user.Role == "cashier" {

			c.Redirect(http.StatusSeeOther,"/cashier/orders",)

		} else {

			c.Redirect(http.StatusSeeOther,"/",)
		}

}

// func OrderDetails(c *gin.Context) {
func (h *OrderHandler) OrderDetails(c *gin.Context) {

	id := c.Param("id")

	// order, err := repository.GetOrderByID(id)
	order, err := h.OrderService.GetOrderByID(id)

	// history, err := repository.GetOrderHistory(id)
	history, err := h.OrderService.GetOrderHistory(id)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	user := helpers.GetSessionData(c)


	c.HTML(
		http.StatusOK,"ad_order_details.html",gin.H{
			"order": order,
			"history": history,
			"title": "Order Detail",
		    "role": user.Role,
			"username": user.Username,
			"name":       user.Name,
			"Surname":       user.Surname,
		},
	)
}

// func TrackOrder(c *gin.Context) {
func (h *OrderHandler) TrackOrder(c *gin.Context) {

	id := c.Param("id")

	// order, err := repository.GetOrderByID(id)
	order, err := h.OrderService.GetOrderByID(id)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.HTML(
		http.StatusOK,
		"track_order.html",
		gin.H{
			"order": order,
		},
	)
}