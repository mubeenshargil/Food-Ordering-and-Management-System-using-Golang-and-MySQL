package routes

import (
	"food-ordering-system-project/handlers"
	"food-ordering-system-project/middleware"

	"github.com/gin-gonic/gin"
	"food-ordering-system-project/config"
	"food-ordering-system-project/repository"
	"food-ordering-system-project/services"
)

func SetupRoutes(r *gin.Engine) {
     
	menuRepo := &repository.MenuRepository{
	DB: config.DB,
	}

	orderRepo := &repository.OrderRepository{
		DB: config.DB,
	}

	menuService := &services.MenuService{
		MenuRepo: menuRepo,
	}

	orderService := &services.OrderService{
		OrderRepo: orderRepo,
		MenuRepo:  menuRepo,
	}

	adminService := &services.AdminService{
		MenuRepo: menuRepo,
	}

	menuHandler := &handlers.MenuHandler{
		MenuService: menuService,
	}

	orderHandler := &handlers.OrderHandler{
		OrderService: orderService,
	}

	adminHandler := &handlers.AdminHandler{
		AdminService: adminService,
	}
	reportRepo := &repository.ReportRepository{
	DB: config.DB,
   }

	reportService := &services.ReportService{
		ReportRepo: reportRepo,
	}

	reportHandler := &handlers.ReportHandler{
		ReportService: reportService,
	}
	userRepo := &repository.UserRepository{
	DB: config.DB,
	}

	userService := &services.UserService{
		UserRepo: userRepo,
	}

	userHandler := &handlers.UserHandler{
		UserService: userService,
	}

	authHandler := &handlers.AuthHandler{
	UserService: userService,
   }

	// PUBLIC ROUTES

	r.GET("/", menuHandler.MenuPage)

	r.GET("/checkout/:id", orderHandler.CheckoutPage)

	r.POST("/place-order", orderHandler.PlaceOrder)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	r.POST("/register", authHandler.Register)

	r.POST("/login", authHandler.Login)

	r.GET("/logout", authHandler.Logout)
	
	r.GET("/track-order/:id",orderHandler.TrackOrder,)

	// ADMIN ROUTES

	admin := r.Group("/admin")

	admin.Use(middleware.AdminRequired())

	{
		admin.GET("/", adminHandler.AdminDashboard)
        
		admin.GET("/dashboard", adminHandler.AdminDashboard)
		
		admin.GET("/menu", adminHandler.AdminMenuPage)

		admin.GET("/orders", orderHandler.AdminOrders)

		// admin.GET("/users", handlers.AdminUsersPage)

		admin.POST("/menu", menuHandler.AddMenu)

		admin.POST("/update-menu/:id", menuHandler.UpdateMenu)

		admin.POST("/update-order-status/:id", orderHandler.UpdateOrderStatus)
        
		admin.GET("/users", userHandler.AdminUsersPage)

		admin.POST("/add-user", userHandler.AddUser)

		admin.POST("/update-user/:id", userHandler.UpdateUser)

		// admin.GET("/delete-user/:id", handlers.DeleteUser)
		admin.GET("/order-details/:id",orderHandler.OrderDetails)

		admin.GET("/reports/sales",reportHandler.SalesReport)
		
		admin.GET("/reports/menu-performance",reportHandler.MenuPerformanceReport)

		admin.GET("/reports/customer-orders",reportHandler.CustomerOrdersReport)

		admin.GET("/reports/all-orders",reportHandler.OrdersReport)

	}


	// CASHIER ROUTES

			cashier := r.Group("/cashier")

			cashier.Use(middleware.CashierRequired())

			{
                cashier.GET("/dashboard", adminHandler.AdminDashboard)

				cashier.GET("/", orderHandler.AdminOrders)

				cashier.GET("/orders", orderHandler.AdminOrders)

				cashier.POST("/update-order-status/:id",orderHandler.UpdateOrderStatus)

				cashier.GET("/order-details/:id",orderHandler.OrderDetails)
			}
			
}