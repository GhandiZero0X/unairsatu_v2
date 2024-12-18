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
	// api.Post("/createUser", controllers.CreateUser) // api untuk create user

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
	Role.Post("/createRole", controllers.CreateRole) // api untuk Create a new role
	Role.Get("/getAllRole", controllers.GetRoles) // api untuk Get all roles
	Role.Get("/getRole/:id", controllers.GetRoleByID) // api untuk Get a role by ID
	Role.Put("/updateRole/:id", controllers.UpdateRole) // api untuk Update a role by ID
	Role.Delete("/deleteRole/:id", controllers.DeleteRole) // api untuk Delete a role by ID

	// Modul routing
	Modul := app.Group("/modul")
	Modul.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	Modul.Post("/createModul", controllers.CreateModul) // api untuk Create a new modul
	Modul.Get("/getAllModul", controllers.GetAllModul) // api untuk Get all moduls
	Modul.Get("/getModul/:id", controllers.GetModulByID) // api untuk Get a modul by ID
	Modul.Put("/updateModul/:id", controllers.UpdateModul) // api untuk Update a modul by ID
	Modul.Delete("/deleteModul/:id", controllers.DeleteModul) // api untuk Delete a modul by ID

	// Kategori routing
	Kategori := app.Group("/kategori")
	Kategori.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	Kategori.Post("/createKategori", controllers.CreateKategori) // api untuk Create a new kategori
	Kategori.Get("/getAllKategori", controllers.GetAllKategori) // api untuk Get all kategoris
	Kategori.Get("/getKategori/:id", controllers.GetKategoriByID) // api untuk Get a kategori by ID
	Kategori.Put("/updateKategori/:id", controllers.UpdateKategori) // api untuk Update a kategori by ID
	Kategori.Delete("/deleteKategori/:id", controllers.DeleteKategori) // api untuk Delete a kategori by ID

	// Jenis_user routing
	JenisUser := app.Group("/jenis_user")
	JenisUser.Use(middleware.AuthMiddleware, middleware.RoleMiddleware("674039236461fc1488d67fec"))
	JenisUser.Post("/createJenisUser", controllers.CreateJenisUser) // api untuk Create a new jenis user
	JenisUser.Get("/getAllJenisUser", controllers.GetAllJenisUser) // api untuk Get all jenis user
	JenisUser.Get("/getJenisUser/:id", controllers.GetJenisUserByID) // api untuk Get jenis user by ID
	JenisUser.Put("/updateTemplates/:id", controllers.UpdateTemplates) // api untuk Update templates modul jenis user
	JenisUser.Delete("/deleteJenisUser/:id", controllers.DeleteJenisUserByID) // api untuk Delete jenis user by ID

}
