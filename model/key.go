package model 

type Key struct { 
	ID uint `json:"id" validate:"omitempty,number"` //PK 
	Value string `gorm:"not null;type:varchar(255)" json:"value" validate:"required,max=255"`
	//Has many FK
	Translations []Translation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
