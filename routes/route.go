package routes

import (
	"project-crud_baru/controllers"
	middleware "project-crud_baru/middlewares" // Pastikan untuk mengimpor middleware

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
	Users.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
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

	// Role routing
	Role := app.Group("/role")
	Role.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	Role.Post("/createRole", controllers.CreateRole) // Create a new role
	Role.Get("/getAllRole", controllers.GetRoles) // Get all roles
	Role.Get("/getRole/:id", controllers.GetRoleByID) // Get a role by ID
	Role.Put("/updateRole/:id", controllers.UpdateRole) // Update a role by ID
	Role.Delete("/deleteRole/:id", controllers.DeleteRole) // Delete a role by ID

	// Modul routing
	Modul := app.Group("/modul")
	Modul.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	Modul.Post("/createModul", controllers.CreateModul) // Create a new modul
	Modul.Get("/getAllModul", controllers.GetAllModul) // Get all moduls
	Modul.Get("/getModul/:id", controllers.GetModulByID) // Get a modul by ID
	Modul.Put("/updateModul/:id", controllers.UpdateModul) // Update a modul by ID
	Modul.Delete("/deleteModul/:id", controllers.DeleteModul) // Delete a modul by ID

	// Kategori routing
	Kategori := app.Group("/kategori")
	Kategori.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	Kategori.Post("/createKategori", controllers.CreateKategori) // Create a new kategori
	Kategori.Get("/getAllKategori", controllers.GetAllKategori) // Get all kategoris
	Kategori.Get("/getKategori/:id", controllers.GetKategoriByID) // Get a kategori by ID
	Kategori.Put("/updateKategori/:id", controllers.UpdateKategori) // Update a kategori by ID
	Kategori.Delete("/deleteKategori/:id", controllers.DeleteKategori) // Delete a kategori by ID

	// Jenis_user routing
	JenisUser := app.Group("/jenis_user")
	JenisUser.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	JenisUser.Post("/createJenisUser", controllers.CreateJenisUser) // Create a new jenis user
	JenisUser.Get("/getAllJenisUser", controllers.GetAllJenisUser) // Get all jenis user
	JenisUser.Get("/getJenisUser/:id", controllers.GetJenisUserByID) // Get jenis user by ID
	JenisUser.Put("/updateTemplates/:id", controllers.UpdateTemplates) // Update templates modul jenis user
	JenisUser.Delete("/deleteJenisUser/:id", controllers.DeleteJenisUserByID) // Delete jenis user by ID

}
