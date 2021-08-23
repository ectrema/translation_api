package languagerepo

import (
	"translations-api/database"
	"translations-api/model"
)

func Delete(id string) (bool, error) {
	result := database.DB.Delete(&model.Language{}, id)
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
