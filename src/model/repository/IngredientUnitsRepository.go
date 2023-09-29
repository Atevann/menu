package repository

import (
	"gorm.io/gorm"
	"menu/src/model/entity"
)

type IngredientUnitsRepository struct {
	Db *gorm.DB
}

func (repo *IngredientUnitsRepository) GetMany(limit int) ([]entity.IngredientUnit, error) {
	var ingredientUnits []entity.IngredientUnit
	result := repo.Db.
		Limit(limit).
		Find(&ingredientUnits)

	return ingredientUnits, result.Error
}
