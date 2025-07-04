package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateUserHandler handles POST /users
func CreateUserHandler(ctx *fiber.Ctx) error {
	log.Println("📩 CreateUserHandler called")

	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := repository.CreateUser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// GetUserByIDHandler handles GET /users/:id
func GetUserByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := repository.GetUserByID(id)

	//if returned err variable from GetUserByID is not nil, it means user was not found
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return ctx.JSON(user)
}

// GetAllUsersHandler handles GET /users
func GetAllUsersHandler(c *fiber.Ctx) error {

	log.Println("📄 GetAllUsersHandler called")
	users, err := repository.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// UpdateUserHandler handles PUT /users/:id
func UpdateUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateUser model.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updateData := bson.M{}
	if updateUser.Name != "" {
		updateData["name"] = updateUser.Name
	}
	if updateUser.Email != "" {
		updateData["email"] = updateUser.Email
	}

	result, err := repository.UpdateUser(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}



// GetAllUsersHandler handles GET /users
// func GetAllUsersHandler(ctx *fiber.Ctx) error {
// 	users, err := repository.GetAllUsers()
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to fetch users",
// 		})
// 	}

// 	return ctx.JSON(users)
// }

// UpdateUserHandler handles PUT /users/:id
// func UpdateUserHandler(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")

// 	var updateData map[string]interface{}
// 	if err := ctx.BodyParser(&updateData); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	result, err := repository.UpdateUser(id, updateData)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to update user",
// 		})
// 	}

// 	return ctx.JSON(result)
// }

// DeleteUserHandler handles DELETE /users/:id
// func DeleteUserHandler(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")

// 	result, err := repository.DeleteUser(id)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to delete user",
// 		})
// 	}

// 	return ctx.JSON(result)
// }







// package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"log"

// 	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
// 	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"

// 	"github.com/gorilla/mux"
// )

// // CreateUserHandler handles POST /users
// func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("📩 CreateUserHandler called")

// 	var user model.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	result, err := repository.CreateUser(&user)
// 	if err != nil {
// 		http.Error(w, "Failed to create user", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }

// // GetUserByIDHandler handles GET /users/{id}
// func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	user, err := repository.GetUserByID(id)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// // GetAllUsersHandler handles GET /users
// func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	users, err := repository.GetAllUsers()
// 	if err != nil {
// 		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(users)
// }

// // UpdateUserHandler handles PUT /users/{id}
// func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	var updateData map[string]interface{}
// 	err := json.NewDecoder(r.Body).Decode(&updateData)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	result, err := repository.UpdateUser(id, updateData)
// 	if err != nil {
// 		http.Error(w, "Failed to update user", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }

// // DeleteUserHandler handles DELETE /users/{id}
// func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	result, err := repository.DeleteUser(id)
// 	if err != nil {
// 		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }
