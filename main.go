package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unmika/todo-mvc/config"
	"github.com/unmika/todo-mvc/routes"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	r := gin.Default()
	routes.Serve(r)

	r.Run()
}
