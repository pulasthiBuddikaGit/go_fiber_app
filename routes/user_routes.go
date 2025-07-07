


package routes

import (

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/handler"
)

// RegisterUserRoutes defines routes related to users
func RegisterUserRoutes(app fiber.Router) {
	userGroup := app.Group("/users")

	userGroup.Post("/", handler.CreateUserHandler)
	userGroup.Get("/", handler.GetAllUsersHandler)
	userGroup.Get("/:id", handler.GetUserByIDHandler)
	userGroup.Put("/:id", handler.UpdateUserHandler)
	userGroup.Delete("/:id", handler.DeleteUserHandler)
}







// package routes

// import (
// 	"net/http"

// 	"github.com/pulasthiBuddikaGit/go_fiber_app/handler"

// 	"github.com/gorilla/mux"
// )

// func RegisterUserRoutes(router *mux.Router) {
// 	userRouter := router.PathPrefix("/users").Subrouter()

// 	userRouter.HandleFunc("", handler.CreateUserHandler).Methods(http.MethodPost)
// 	userRouter.HandleFunc("", handler.GetAllUsersHandler).Methods(http.MethodGet)
// 	userRouter.HandleFunc("/{id}", handler.GetUserByIDHandler).Methods(http.MethodGet)
// 	userRouter.HandleFunc("/{id}", handler.UpdateUserHandler).Methods(http.MethodPut)
// 	userRouter.HandleFunc("/{id}", handler.DeleteUserHandler).Methods(http.MethodDelete)
// }
