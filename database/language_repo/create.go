package languagerepo

import (
	"translations-api/database"
	"translations-api/model"
)

func CreateLanguages(apps *[]model.Language) {
	database.DB.Create(*apps)
}

func CreateOneLanguage(app *model.Language) {
	database.DB.Create(app)
}
