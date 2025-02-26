package app

import (
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App,
	categoryController controller.CategoryController,
	customerController controller.CustomerController,
	employeeController controller.EmployeeController,
	productController controller.ProductController) {

	authMiddleware := middleware.NewAuthMiddleware()

	api := app.Group("/api", authMiddleware)

	// Routes untuk Category
	categories := api.Group("/categories")
	categories.Get("/", categoryController.FindAll)
	categories.Get("/:categoryId", categoryController.FindById)
	categories.Post("/", categoryController.Create)
	categories.Put("/:categoryId", categoryController.Update)
	categories.Delete("/:categoryId", categoryController.Delete)

	// Routes untuk Customer
	customers := api.Group("/customers")
	customers.Get("/", customerController.FindAll)
	customers.Get("/:customerId", customerController.FindById)
	customers.Post("/", customerController.Create)
	customers.Put("/:customerId", customerController.Update)
	customers.Delete("/:customerId", customerController.Delete)

	// Routes untuk Employee
	employees := api.Group("/employees")
	employees.Get("/", employeeController.FindAll)
	employees.Get("/:employeeId", employeeController.FindById)
	employees.Post("/", employeeController.Create)
	employees.Put("/:employeeId", employeeController.Update)
	employees.Delete("/:employeeId", employeeController.Delete)

	// Routes untuk Product
	products := api.Group("/products")
	products.Get("/", productController.FindAll)
	products.Get("/:productId", productController.FindById)
	products.Post("/", productController.Create)
	products.Put("/:productId", productController.Update)
	products.Delete("/:productId", productController.Delete)
}
