package entity

type Ingredient struct {
	Id             uint `gorm:"primarykey"`
	Name           string
	UnitId         uint
	IngredientUnit IngredientUnit `gorm:"foreignKey:Id;references:UnitId;"`
}
