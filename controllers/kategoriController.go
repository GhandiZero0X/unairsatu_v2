package controllers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"project-crud_baru/config"
	"project-crud_baru/models"
)

var kategoriCollection *mongo.Collection = config.GetCollection("kategori")

// CreateKategori creates a new kategori
func CreateKategori(c *fiber.Ctx) error {
	var kategori models.Kategori

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	kategori.ID = primitive.NewObjectID()
	_, err := kategoriCollection.InsertOne(context.Background(), kategori)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(kategori)
}

// GetAllKategori retrieves all kategori
func GetAllKategori(c *fiber.Ctx) error {
	var kategoris []models.Kategori

	cursor, err := kategoriCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var kategori models.Kategori
		cursor.Decode(&kategori)
		kategoris = append(kategoris, kategori)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(kategoris)
}

// GetKategoriByID retrieves a kategori by its ID
func GetKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var kategori models.Kategori
	err = kategoriCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&kategori)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Kategori not found"})
	}

	return c.Status(http.StatusOK).JSON(kategori)
}

// UpdateKategori updates an existing kategori
func UpdateKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var kategori models.Kategori
	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	update := bson.M{
		"$set": kategori,
	}

	_, err = kategoriCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(kategori)
}

// DeleteKategori deletes a kategori by its ID
func DeleteKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, err = kategoriCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Kategori deleted"})
}