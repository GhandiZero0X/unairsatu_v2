package controllers

import (
	"context"
	"project-crud_baru/config"
	"project-crud_baru/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection reference for modul
var modulCollection *mongo.Collection = config.GetCollection("modul")

// create modul
func CreateModul(c *fiber.Ctx) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the modul input
	var modul models.Modul
	if err := c.BodyParser(&modul); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate and parse IDKategori, CreatedBy, and UpdatedBy as ObjectID
	idKategori, err := primitive.ObjectIDFromHex(modul.IDKategori.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_kategori"})
	}
	createdBy, err := primitive.ObjectIDFromHex(modul.CreatedBy.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid created_by"})
	}
	updatedBy, err := primitive.ObjectIDFromHex(modul.UpdatedBy.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid updated_by"})
	}

	modul.IDKategori = idKategori
	modul.CreatedBy = createdBy
	modul.UpdatedBy = updatedBy

	// Set created and updated timestamps
	modul.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	modul.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Insert the modul into the collection
	result, err := modulCollection.InsertOne(ctx, modul)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create modul"})
	}

	// Return success response with the inserted ID
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Modul created successfully",
		"id":      result.InsertedID,
	})
}

// GetAllModul 
func GetAllModul(c *fiber.Ctx) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the modul input
	var moduls []models.Modul

	// Find all moduls in the collection
	cursor, err := modulCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch moduls"})
	}

	// Iterate over the cursor and decode each modul
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var modul models.Modul
		cursor.Decode(&modul)
		moduls = append(moduls, modul)
	}

	// Return the moduls
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"moduls": moduls,
	})
}

// GetModulByID retrieves a modul document by ID from the database
func GetModulByID(c *fiber.Ctx) error {
	// Get the modul ID from the URL parameters
	id := c.Params("id")
	modulID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid modul ID"})
	}

	// Create a filter to search by ID
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var modul models.Modul
	err = modulCollection.FindOne(ctx, bson.M{"_id": modulID}).Decode(&modul)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Modul not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch modul data"})
	}

	// Return the modul data
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"modul": modul,
	})
}

// UpdateModul updates an existing modul based on the provided ID
func UpdateModul(c *fiber.Ctx) error {
	// Get the modul ID from the URL parameters
	id := c.Params("id")
	modulID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid modul ID"})
	}

	// Parse the new modul data from the request body
	var updatedModul models.Modul
	if err := c.BodyParser(&updatedModul); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Set the updated time
	updatedModul.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Create the update filter and update fields
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updatedModul,
	}

	// Perform the update operation
	result, err := modulCollection.UpdateOne(ctx, bson.M{"_id": modulID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update modul"})
	}

	// If no documents were modified, return an error
	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Modul not found"})
	}

	// Return success response
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Modul updated successfully",
	})
}

// DeleteModul deletes a modul by its ID
func DeleteModul(c *fiber.Ctx) error {
	// Get the modul ID from the URL parameters
	id := c.Params("id")
	modulID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid modul ID"})
	}

	// Create the delete filter
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Perform the delete operation
	result, err := modulCollection.DeleteOne(ctx, bson.M{"_id": modulID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete modul"})
	}

	// If no documents were deleted, return an error
	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Modul not found"})
	}

	// Return success response
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Modul deleted successfully",
	})
}
