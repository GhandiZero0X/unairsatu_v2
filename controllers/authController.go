package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"project-crud_baru/config"
	"project-crud_baru/controllers/utils"
)

func Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	// Cek username di database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user bson.M
	err := config.GetCollection("users").FindOne(ctx, bson.M{"username": input.Username}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Username not found"})
	}

	// Ambil password hash dari database
	storedPasswordHash, ok := user["pass"].(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid password format"})
	}

	// Verifikasi password menggunakan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(input.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate token JWT
	token, err := utils.GenerateJWT(input.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}
