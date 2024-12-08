package routes

import (
	"project-crud_baru/controllers"
	// middleware "project-crud_baru/middlewares" // Pastikan untuk mengimpor middleware

	"github.com/gofiber/fiber/v2"
)

func RouteApp(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", controllers.HomeFunc)

	// User tanpa autentikasi
	// api.Post("/users/login", controllers.Login) // Pastikan login berada di luar grup yang menggunakan middleware

	// User routing
	Users := app.Group("/users")
	// login
	Users.Post("/login", controllers.Login)
	// middlewar
	// Users.Use(middleware.AuthMiddleware)
	Users.Post("/createUser", controllers.CreateUser) // api untuk create user
	Users.Get("/getAllUser", controllers.GetUsers) // api untuk get all user
	Users.Get("/getUser/:id", controllers.GetUserOne) // api untuk get user by id
	Users.Put("/updateUser/:id", controllers.UpdateUser) // api untuk update user
	Users.Put("/changePassword/:id", controllers.ChangePassword) // api untuk change password
	Users.Put("/uploadPhoto/:id", controllers.UploadPhoto) // api untuk upload photo
	Users.Delete("/deleteUser/:id", controllers.DeleteUser) // api untuk delete user
	Users.Get("/getUserModul/:id", controllers.GetAllModulByUserID) // api untuk get all modul by user id
	Users.Post("/addModul/:id", controllers.AddModulToUser) // api untuk add modul to user
	Users.Put("/change-jenis-user/:id", controllers.ChangeJenisUserByUserID) // api untuk change jenis user

	// Modul routing
	Modul := app.Group("/modul")
	Modul.Post("/createModul", controllers.CreateModul) // Create a new modul
	Modul.Get("/getAllModul", controllers.GetAllModul) // Get all moduls
	Modul.Get("/getModul/:id", controllers.GetModulByID) // Get a modul by ID
	Modul.Put("/updateModul/:id", controllers.UpdateModul) // Update a modul by ID
	Modul.Delete("/deleteModul/:id", controllers.DeleteModul) // Delete a modul by ID

	// Jenis_user routing
	JenisUser := app.Group("/jenis_user")
	JenisUser.Get("/getAllJenisUser", controllers.GetAllJenisUser) // Get all jenis user
	JenisUser.Get("/getJenisUser/:id", controllers.GetJenisUserByID) // Get jenis user by ID
	JenisUser.Put("/updateTemplates/:id", controllers.UpdateTemplates) // Update templates modul jenis user

}
