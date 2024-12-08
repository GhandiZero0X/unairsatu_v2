package controllers

import (
	"context"
	"net/http"
	"time"
    "fmt"
    "os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"

	"project-crud_baru/config"
	"project-crud_baru/controllers/utils"
	"project-crud_baru/models"
)

var userCollection *mongo.Collection = config.GetCollection("users")

// Create User
func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Cek apakah username sudah ada
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
	}

	// Parse id_jenis_user from string to ObjectID
	idJenisUser, err := primitive.ObjectIDFromHex(user.Id_jenis_user.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_jenis_user format"})
	}
	user.Id_jenis_user = idJenisUser

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Hash password 2
	hashedPassword2, err := bcrypt.GenerateFromPassword([]byte(user.Pass_2), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password 2"})
	}

	// Generate random token
	token, err := utils.GenerateRandomString(32)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	now := time.Now().In(loc)
	user.Created_at = primitive.NewDateTimeFromTime(now)
	user.Updated_at = primitive.NewDateTimeFromTime(now)

	// Assume Created_by and Updated_by are set to a predefined user ID
	user.Created_by = primitive.NewObjectID() // Set to actual creator's ID
	user.Updated_by = user.Created_by       // Initially set as same as creator

	// Generate AuthKey for the user (can be a random string or token)
	user.AuthKey, err = utils.GenerateRandomString(32)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate auth key"})
	}

	newUser := models.User{
		ID:            primitive.NewObjectID(),
		Username:      user.Username,
		Nm_user:       user.Nm_user,
		Pass:          string(hashedPassword),
		Email:         user.Email,
		Role_aktif:    user.Role_aktif,
		Created_at:    user.Created_at,
		Updated_at:    user.Updated_at,
		Created_by:    user.Created_by,
		Updated_by:    user.Updated_by,
		AuthKey:       user.AuthKey,
		Jenis_kelamin: user.Jenis_kelamin,
		Photo:         user.Photo,
		Phone:         user.Phone,
		Token:         token,
		Id_jenis_user: user.Id_jenis_user,
		Pass_2:        string(hashedPassword2),
		Moduls:        user.Moduls,
	}

	_, errIns := userCollection.InsertOne(ctx, newUser)
	if errIns != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errIns.Error()})
	}

	return c.Status(http.StatusCreated).JSON(newUser)
}

// Get All Users
func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err = cursor.All(ctx, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(users)
}

// Get User by ID
func GetUserOne(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Update User by ID
func UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Receive updated user data from request body
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate Modul IDs
	for i, modul := range user.Moduls {
		modulID, err := primitive.ObjectIDFromHex(modul.ModulID.Hex())
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Invalid modul_id format at index %d", i),
			})
		}
		user.Moduls[i].ModulID = modulID
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Hash password 2
	hashedPassword2, err := bcrypt.GenerateFromPassword([]byte(user.Pass_2), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password 2"})
	}

	// Create update data
	update := bson.M{
		"username":      user.Username,
		"nm_user":       user.Nm_user,
		"pass":          string(hashedPassword),
		"email":         user.Email,
		"role_aktif":    user.Role_aktif,
		"jenis_kelamin": user.Jenis_kelamin,
		"photo":         user.Photo,
		"phone":         user.Phone,
		"pass_2":        string(hashedPassword2),
		"moduls":        user.Moduls, // Include modules if provided
		"id_jenis_user": user.Id_jenis_user,
		"updated_at":    primitive.NewDateTimeFromTime(time.Now()),
		"updated_by":    primitive.NewObjectID(), // Set updated_by to the current user
	}

	// Perform update
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// If no document is affected, user not found
	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Fetch updated user data
	var updatedUser models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&updatedUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}

// Change Password
func ChangePassword(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mendapatkan ID pengguna dari parameter
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Mendapatkan password lama dan baru dari request body
	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Mengambil data pengguna berdasarkan ID
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Verifikasi password lama dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(input.OldPassword)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Old password is incorrect"})
	}

	// Hash password baru
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	// Melakukan update password di database
	update := bson.M{"pass": string(hashedNewPassword)}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Password updated successfully"})
}

// Upload Photo
func UploadPhoto(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get user ID from params
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Retrieve uploaded file
	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to retrieve file"})
	}

	// Create storage/images directory if not exists
	if _, err := os.Stat("./storage/images"); os.IsNotExist(err) {
		err := os.MkdirAll("./storage/images", os.ModePerm)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create directory"})
		}
	}

	// Generate new filename based on timestamp
	timestamp := time.Now().Format("20060102150405.000")
	extension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", timestamp, extension)
	filePath := fmt.Sprintf("./storage/images/%s", newFileName)

	// Save the file to ./storage/images directory
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Update user document with the new file path in the `photo` field
	update := bson.M{"photo": filePath}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user photo"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Photo uploaded successfully", "photo_url": filePath})
}

// Delete User
func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

// Get All Modul by User ID
func GetAllModulByUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(user.Moduls)
}

// Add Modul to User
func AddModulToUser(c *fiber.Ctx) error {
    id := c.Params("id") // Ambil ID User dari parameter URL
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    // Parse modul baru dari body permintaan
    var newModul models.Moduls
    if err := c.BodyParser(&newModul); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid modul data"})
    }

    // Setkan ID untuk modul baru (jika belum memiliki ID)
    if newModul.ModulID.IsZero() {
        newModul.ModulID = primitive.NewObjectID()
    }

    // Tambahkan modul baru ke array moduls
    filter := bson.M{"_id": objID}
    update := bson.M{"$push": bson.M{"moduls": newModul}}

    _, err = userCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user moduls"})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Modul added successfully",
        "modul":   newModul,
    })
}

// Change Jenis User by User ID
func ChangeJenisUserByUserID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get user ID from params
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Parse new jenis_user ID from request body
	var input struct {
		IdJenisUser string `json:"id_jenis_user"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Fetch the new jenis_user data from the database
	var jenisUser models.JenisUser
	jenisUserCollection := config.GetCollection("jenis_user")
	jenisUserID, err := primitive.ObjectIDFromHex(input.IdJenisUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_jenis_user format"})
	}
	err = jenisUserCollection.FindOne(ctx, bson.M{"_id": jenisUserID}).Decode(&jenisUser)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "JenisUser not found"})
	}

	// Fetch moduls data based on templates
	modulCollection := config.GetCollection("modul")
	var newModuls []models.Moduls
	for _, template := range jenisUser.Templates {
		var modul models.Modul
		err := modulCollection.FindOne(ctx, bson.M{"_id": template.IDModul}).Decode(&modul)
		if err != nil {
			// Handle case where modul not found
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Modul not found for template ID"})
		}

		// Append fetched modul to newModuls
		newModuls = append(newModuls, models.Moduls{
			ModulID:  modul.ID,
			NmModul:  modul.NmModul,
			KetModul: modul.KetModul,
			Alamat:   modul.Alamat,
			GbrIcon:  modul.GbrIcon,
		})
	}

	// Update user's jenis_user and moduls
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"id_jenis_user": jenisUserID, // use ObjectID
			"moduls":        newModuls,
			"updated_at":    primitive.NewDateTimeFromTime(time.Now()),
			"updated_by":    primitive.NewObjectID(), // Replace with current user's ID
		},
	}

	userCollection := config.GetCollection("users")
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User jenis_user changed successfully",
		"moduls":  newModuls,
	})
}
