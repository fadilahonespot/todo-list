package response

import (
	"github.com/fadilahonespot/todo-list/utils/model"
	"github.com/gofiber/fiber/v2"
)

func HandleSuccess(c *fiber.Ctx, data interface{}) error {
	resp := model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
