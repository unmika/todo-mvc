package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unmika/todo-mvc/config"
	"github.com/unmika/todo-mvc/routes"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	config.InitDB()
	defer config.CloseDB()

	r := gin.Default()
	routes.Serve(r)

	r.Run()
}
