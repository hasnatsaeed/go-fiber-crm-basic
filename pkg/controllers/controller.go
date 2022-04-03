package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-crm-basic/pkg/models"
)

func GetLeads(ctx *fiber.Ctx) error {

	leads := models.GetLeads()
	return ctx.Status(fiber.StatusOK).JSON(leads)

}

func GetLeadById(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	lead, err := models.GetLeadById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Lead not found!!")
	}
	return ctx.Status(fiber.StatusOK).JSON(lead)
}

func CreateLead(ctx *fiber.Ctx) error {

	lead := models.Lead{}
	if err := ctx.BodyParser(&lead); err != nil {
		panic(err)
	}

	lead.Create()
	return ctx.Status(fiber.StatusOK).JSON(lead)
}

func DeleteLeadById(ctx *fiber.Ctx) error {

	id, _ := ctx.ParamsInt("id")
	lead, err := models.DeleteLeadById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Lead not found!!")
	}
	return ctx.Status(fiber.StatusOK).JSON(lead)
}
