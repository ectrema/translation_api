package keyrepo

import (
	"translations-api/database"
	"translations-api/model"
)

func Update(key *model.Key, updatedInfo *model.Key) (bool, error) {
	result := database.DB.Model(key).Updates(*updatedInfo)
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
