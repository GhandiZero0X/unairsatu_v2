package controllers

// update templates modul jenis user
import (
	"context"
	"net/http"
	"project-crud_baru/config"
	"project-crud_baru/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jenisUserCollection *mongo.Collection = config.GetCollection("jenis_user")

// CreateJenisUser baru
func CreateJenisUser(c *fiber.Ctx) error {
	jenisUser := new(models.JenisUser)

	if err := c.BodyParser(jenisUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	jenisUser.ID = primitive.NewObjectID()
	_, err := jenisUserCollection.InsertOne(context.Background(), jenisUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create JenisUser"})
	}

	return c.Status(http.StatusCreated).JSON(jenisUser)
}

// getAll jenis user
func GetAllJenisUser(c *fiber.Ctx) error {
	cursor, err := jenisUserCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get JenisUser"})
	}

	var jenisUsers []models.JenisUser = make([]models.JenisUser, 0)
	for cursor.Next(context.Background()) {
		var jenisUser models.JenisUser
		cursor.Decode(&jenisUser)
		jenisUsers = append(jenisUsers, jenisUser)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": jenisUsers})
}

// get JenisUser by ID
func GetJenisUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objID}
	jenisUser := jenisUserCollection.FindOne(context.Background(), filter)

	var result models.JenisUser
	if err := jenisUser.Decode(&result); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get JenisUser"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": result})
}

// UpdateTemplates updates the templates of a JenisUser by ID.
func UpdateTemplates(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var newTemplates []models.Template
	if err := c.BodyParser(&newTemplates); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"templates": newTemplates}}

	_, err = jenisUserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update templates"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Templates updated successfully"})
}

// DeleteJenisUserByID deletes a JenisUser by ID.
func DeleteJenisUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objID}
	_, err = jenisUserCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete JenisUser"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "JenisUser deleted successfully"})
}