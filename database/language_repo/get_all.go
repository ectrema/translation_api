package languagerepo

import (
	"translations-api/database"
	"translations-api/model"
)

type BasicKeyInfo struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Tag       string    `json:"slug"`
}

func GetAll() *[]model.Key {
	var languages []model.Language
	result := database.DB.Find(&languages)
	if result.RowsAffected == 0 {
		return nil
	}
	return &languages
}