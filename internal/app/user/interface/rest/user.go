package rest

import (
	"second-spot-backend/internal/app/user/usecase"
	"second-spot-backend/internal/domain/dto"

<<<<<<< HEAD
=======
	"github.com/go-playground/validator/v10"
>>>>>>> bb86e19 (commit add generate token login)
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
<<<<<<< HEAD
	userUsecase usecase.UserUsecaseItf
}

func NewUserHandler(userUsecase usecase.UserUsecaseItf) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
=======
	usecase   usecase.UserUsecaseItf
	validator *validator.Validate
}

func NewUserHandler(usecase usecase.UserUsecaseItf, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		usecase:   usecase,
		validator: validator,
	}
>>>>>>> bb86e19 (commit add generate token login)
}

func (h *UserHandler) SetupRoutes(router fiber.Router) {
	router.Post("/register", h.Register)
<<<<<<< HEAD
=======
	router.Post("/login", h.Login)
>>>>>>> bb86e19 (commit add generate token login)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
<<<<<<< HEAD
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
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.usecase.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.usecase.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
>>>>>>> bb86e19 (commit add generate token login)
}
