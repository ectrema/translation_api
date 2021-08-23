package translationrepo

import (
	"gorm.io/gorm/clause"
	"time"
	"translations-api/database"
	"translations-api/model"
)

type BasicTranslationInfo struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}

type Key struct {
	Key         string
	Translation string
}

func GetAll() *[]model.Translation {
	var translations []model.Translation
	result := database.DB.Find(&translations)
	if result.RowsAffected == 0 {
		return nil
	}
	return &translations
}