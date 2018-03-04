package services

import "github.com/memclutter/gontacts/models"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (s *App) GetStatus() *models.Status {
	return models.NewStatus(true)
}
