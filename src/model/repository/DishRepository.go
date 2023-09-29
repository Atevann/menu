package repository

import (
	"gorm.io/gorm"
	"menu/src/model/entity"
)

type DishRepository struct {
	Db *gorm.DB
}

func (repo *DishRepository) GetMany(limit int) ([]entity.Dish, error) {
	var dishes []entity.Dish

	result := repo.Db.
		Preload("Ingredients.IngredientUnit").
		Limit(limit).
		Find(&dishes)

	return dishes, result.Error
}
