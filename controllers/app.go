package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/services"
)

type App struct {
	service *services.App
}

func NewApp() *App {
	return &App{
		service: services.NewApp(),
	}
}

func (c *App) Status(ctx *gin.Context) {
	model := c.service.GetStatus()
	ctx.JSON(http.StatusOK, model)
}
