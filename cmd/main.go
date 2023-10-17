package main

import (
	"fmt"
	"todo-app/pkg/controllers"
	db "todo-app/pkg/database"
	"todo-app/pkg/models"
	"todo-app/pkg/routes"
	"todo-app/pkg/services"
)

func main() {
	fmt.Println("server is starting...")
	r := routes.SetupRouter()

	// Initialize the database
	db, err := db.InitDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.Todo{})

	todoService := services.NewTodoService(db)
	controllers.NewTodoController(todoService)

	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
