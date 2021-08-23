package translationrepo

import (
	"translations-api/database"
	"translations-api/model"
)

func Update(translation *model.Translation, updatedInfo *model.Translation) (bool, error) {
	result := database.DB.Model(translation).Updates(*updatedInfo)
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
