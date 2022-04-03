package routes

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-crm-basic/pkg/controllers"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/api/v1/lead", controllers.GetLeads)
	app.Get("/api/v1/lead/:id", controllers.GetLeadById)
	app.Post("/api/v1/lead", controllers.CreateLead)
	app.Delete("/api/v1/lead/:id", controllers.DeleteLeadById)
}
