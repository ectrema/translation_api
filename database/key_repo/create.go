package keyrepo

import (
	"translations-api/database"
	"translations-api/model"
)

func CreateKeys(keys *[]model.Key) {
	database.DB.Create(*keys)
}

func CreateOneKey(key *model.Key) {
	database.DB.Create(&key)
}
