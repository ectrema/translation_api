package model

type Language struct { 
	ID uint `json:"id" validate:"omitempty,number"` //PK 
	Name string `gorm:"not null;type:varchar(255)" json:"name" validate:"required,max=255"` 
	Tag string `gorm:"not null;type:varchar(255)" json:"tag" validate:"required,max=255"` 
	FlagUrl string `gorm:"not null;type:varchar(255)" json:"flag_url" validate:"required,max=255"`
	//Has many FK
	Translations []Translation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
