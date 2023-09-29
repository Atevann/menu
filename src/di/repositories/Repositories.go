package repositories

import (
	"gorm.io/gorm"
	"menu/src/model/repository"
)

type Repositories struct {
	Dish            repository.DishRepository
	Ingredient      repository.IngredientRepository
	IngredientUnits repository.IngredientUnitsRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	dish := initDish(db)
	ingredient := initIngredient(db)
	ingredientUnits := initIngredientUnits(db)

	return Repositories{
		Dish:            dish,
		Ingredient:      ingredient,
		IngredientUnits: ingredientUnits,
	}
}

func initDish(db *gorm.DB) repository.DishRepository {
	return repository.DishRepository{
		Db: db,
	}
}

func initIngredient(db *gorm.DB) repository.IngredientRepository {
	return repository.IngredientRepository{
		Db: db,
	}
}

func initIngredientUnits(db *gorm.DB) repository.IngredientUnitsRepository {
	return repository.IngredientUnitsRepository{
		Db: db,
	}
}
