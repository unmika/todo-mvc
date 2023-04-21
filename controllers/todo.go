package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unmika/todo-mvc/models"
	"gorm.io/gorm"
)

type Todos struct {
	DB *gorm.DB
}

func (t *Todos) FindAll(ctx *gin.Context) {
	var todos []models.Todo
	searchText := ctx.Query("search")
	sText := strings.Replace(searchText, " ", "+", -1)

	if searchText != "" {
		t.DB.Where("title like (?) ", "%"+sText+"%").Find(&todos)
	} else {
		t.DB.Find(&todos)
	}

	ctx.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (t *Todos) Create(ctx *gin.Context) {
	var input models.Todo
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.DB.Create(&input).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo": input})
}

func (t *Todos) Delete(ctx *gin.Context) {
	todo, err := t.findTodoByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	t.DB.Delete(&todo)

	ctx.Status(http.StatusNoContent)
}

func (t *Todos) findTodoByID(ctx *gin.Context) (*models.Todo, error) {
	var todo models.Todo
	id := ctx.Param("id")

	if err := t.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
