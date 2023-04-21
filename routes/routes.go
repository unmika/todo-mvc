package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unmika/todo-mvc/config"
	"github.com/unmika/todo-mvc/controllers"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	todosGroup := r.Group("/api/v1/todos")
	todoController := controllers.Todos{DB: db}
	{
		todosGroup.GET("", todoController.FindAll)
		todosGroup.POST("", todoController.Create)
		todosGroup.DELETE("/:id", todoController.Delete)
	}
}
