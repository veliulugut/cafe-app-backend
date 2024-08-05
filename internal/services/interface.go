package services

import (
	"github.com/gofiber/fiber/v2"

	"cafe-app-backend/internal/dtos/authDto"
)

type AuthService interface {
	Login(ctx *fiber.Ctx, request authDto.LoginRequest) (authDto.LoginResponse, int, error)
	ResetPassword(ctx *fiber.Ctx, request authDto.ForgetPasswordRequest) (int, error)
	ForgetPassword(ctx *fiber.Ctx, request authDto.ForgetPasswordRequest) (int, error)
}
