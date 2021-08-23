package translationrepo

import (
	"translations-api/database"
	"translations-api/model"
)

func CreateTranslations(apps *[]model.Translation) {
	database.DB.Create(*apps)
}

func CreateOneTranslation(app *model.Translation) {
	database.DB.Create(app)
}
