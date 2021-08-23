package keyrepo

import (
	"translations-api/database"
	"translations-api/model"
)

type BasicKeyInfo struct {
	ID        uint      `json:"id"`
	Value     string    `json:"slug"`
}

func GetAll() *[]model.Key {
	var keys []model.Key
	result := database.DB.Find(&apps)
	if result.RowsAffected == 0 {
		return nil
	}
	return &keys
}