package model 

type Translation struct {
	//Belongs to FK
	KeyID uint `gorm:"not null; primaryKey; autoIncrement:false;" json:"key_id" validate:"required,number"` 
	Key *Key `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"omitempty,dive"` 
	LanguageID uint `gorm:"not null; primaryKey; autoIncrement:false;" json:"language_id" validate:"required,number"` 
	Language *Language `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"omitempty,dive"` 
	Value string `gorm:"not null;type:varchar(255)" json:"value" validate:"required,max=255"`
}
