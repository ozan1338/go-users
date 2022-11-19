package routes

import (
	"go-users/src/controllers"
	"go-users/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)
	api.Get("users", controllers.Users)
	api.Get("users/:id", controllers.GetUser)
	
	authenticated := api.Use(middlewares.IsAuthenticated)

	authenticated.Get("user/:scope", controllers.User)
	authenticated.Post("logout", controllers.Logout)
	authenticated.Put("users/info", controllers.UpdateInfo)
	authenticated.Put("users/password", controllers.UpdatePassword)

}
