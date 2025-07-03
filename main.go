package main

import (
	"fmt"
	"net/http"

	"github.com/pulasthiBuddikaGit/go_fiber_app/config"
	"github.com/pulasthiBuddikaGit/go_fiber_app/storage"
)

func main() {
	cfg := config.LoadConfig()
	storage.InitMongo(cfg)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
