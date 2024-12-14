package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Cari user berdasarkan username
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user bson.M
	err := config.GetCollection("users").FindOne(ctx, bson.M{"username": input.Username}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Username not found"})
	}

	// Verifikasi password
	storedPasswordHash, ok := user["pass"].(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid password format"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(input.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Ambil ID Role dari User
	idRole, ok := user["id_role"].(primitive.ObjectID)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Role ID not found"})
	}

	// Cari nama role berdasarkan ID role
	var role bson.M
	err = config.GetCollection("role").FindOne(ctx, bson.M{"_id": idRole}).Decode(&role)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Role not found"})
	}

	// Pastikan role memiliki nama
	roleName, ok := role["nm_role"].(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid role format"})
	}

	// Generate JWT token dengan role
	token, err := utils.GenerateJWT(input.Username, roleName, idRole.Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token, "role": roleName})
}

