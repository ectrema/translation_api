package languagerepo

import (
	"translations-api/database"
	"translations-api/model"
)

func Update(language *model.Language, updatedInfo *model.Language) (bool, error) {
	result := database.DB.Model(language).Updates(*updatedInfo)
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
