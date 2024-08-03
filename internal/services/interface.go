package services

import (
	"cafe-app-backend/internal/dtos/authDto"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(ctx *fiber.Ctx, request authDto.LoginRequest) (authDto.LoginResponse, int, error)
}
