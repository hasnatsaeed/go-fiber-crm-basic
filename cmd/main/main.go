package main

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-crm-basic/pkg/routes"
	"log"
)

func main() {

	application := fiber.New()
	routes.RegisterRoutes(application)
	log.Fatal(application.Listen(":9010"))
}
