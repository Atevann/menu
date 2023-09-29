package repository

import (
	"gorm.io/gorm"
	"menu/src/model/entity"
)

type IngredientRepository struct {
	Db *gorm.DB
}

func (repo *IngredientRepository) GetMany(limit int) ([]entity.Ingredient, error) {
	var ingredients []entity.Ingredient
	result := repo.Db.
		Preload("IngredientUnit").
		Limit(limit).
		Find(&ingredients)

	return ingredients, result.Error
}
