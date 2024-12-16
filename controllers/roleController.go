package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"project-crud_baru/config"
	"project-crud_baru/models"
)

var roleCollection *mongo.Collection = config.GetCollection("role")

// GetRoles retrieves all roles from the database
func GetRoles(c *fiber.Ctx) error {
	cursor, err := roleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch roles",
		})
	}
	defer cursor.Close(context.Background())

	var roles []models.Role
	if err := cursor.All(context.Background(), &roles); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse roles",
		})
	}

	return c.JSON(roles)
}

// GetRoleByID retrieves a role by its ID from the database
func GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role ID",
		})
	}

	var role models.Role
	if err := roleCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&role); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Role not found",
		})
	}

	return c.JSON(role)
}

// CreateRole creates a new role in the database
func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	role.ID = primitive.NewObjectID()
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	_, err := roleCollection.InsertOne(context.Background(), role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create role",
		})
	}

	return c.Status(http.StatusCreated).JSON(role)
}

// UpdateRole updates an existing role in the database
func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role ID",
		})
	}

	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	role.UpdatedAt = time.Now()

	update := bson.M{
		"$set": role,
	}

	_, err = roleCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update role",
		})
	}

	return c.JSON(role)
}

// DeleteRole deletes a role from the database
func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role ID",
		})
	}

	_, err = roleCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete role",
		})
	}

	return c.SendStatus(http.StatusNoContent)
}