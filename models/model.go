package models

import (
	. "app/common"
	"app/services"
	"time"
)

type Model struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Date      time.Time `json:"date"`
}

func (m *Model) GetByName(name string) *Model {
	error := services.DB.Find(m, "name = ?", name).Error
	if error != nil {
		Log.Printf("Unable to get model!  %s", error, name)
	}

	return m
}

func (m Model) TableName() string {
	return "models"
}
