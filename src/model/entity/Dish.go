package entity

type Dish struct {
	Id          uint `gorm:"primarykey"`
	Name        string
	Recipe      string
	Ingredients []Ingredient `gorm:"many2many:dish_ingredients;"`
}
