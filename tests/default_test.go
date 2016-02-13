package test

import (
	"app/services"
	"app/models"
)

func init() {
	services.InitDb()
	services.DB.AutoMigrate(&models.Model{})
}
