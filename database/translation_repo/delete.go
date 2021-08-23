package translationrepo

import (
	"translations-api/database"
	"translations-api/model"
)

func Delete(keyId string, languageId string) (bool, error) {
	result := database.DB.Where("key_id = " + keyId + " AND language_id = " + languageId).Delete(&model.Translation{})
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
