package router

import ( 
	"github.com/gofiber/fiber/v2" 
) 
func SetupRoutes(app *fiber.App) { 
	app.Use(recover.New()) 
	app.Use(logger.New()) 
	app.Use(middleware.Cors()) 

	app.Get("/", handler.Slash)
	 api := app.Group("/api") 
	 api.Get("/healthcheck", handler.Hello)
	//Internal
	internal := app.Group("/internal", middleware.ValidateInternal()) 
	
	internalKeys := internal.Group("/keys") 
	internalKeys.Post("/", handler.CreateKey) 
	internalKeys.Put("/:id", handler.UpdateKey) 
	internalKeys.Delete("/:id", handler.DeleteKey)
	
	//Translations
	internalTranslations := internal.Group("/translations") 
	internalTranslations.Post("/", handler.CreateTranslation) 
	internalTranslations.Put("/:keyId/:languageId", handler.UpdateTranslation) 
	internalTranslations.Delete("/:keyId/:languageId", handler.DeleteTranslation) 

	//Languages
	internalLanguages := internal.Group("/languages") 
	internalLanguages.Post("/", handler.CreateLanguage) 
	internalLanguages.Put("/:id", handler.UpdateLanguage)
	internalLanguages.Delete("/:id", handler.DeleteLanguage)
	//Translations
	translations := api.Group("/translations") 
	translations.Get("", handler.GetAllTranslations) 
	translations.Get("/:appId/:keyId/:languageId", handler.GetTranslation)
	//Languages
	languages := api.Group("/languages") 
	languages.Get("", handler.GetAllLanguages) 
	languages.Get("/:id", handler.GetLanguage)
	//Keys
	keys := api.Group("/keys") 
	keys.Get("", handler.GetAllKeys) 
	keys.Get("/:id", handler.GetKey)
}
