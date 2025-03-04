package rest

import (
	"second-spot-backend/internal/app/user/usecase"
	"second-spot-backend/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecaseItf
}

func NewUserHandler(userUsecase usecase.UserUsecaseItf) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) SetupRoutes(router fiber.Router) {
	router.Post("/register", h.Register)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request format"})
	}

	response, err := h.userUsecase.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"payload": response,
	})
}
