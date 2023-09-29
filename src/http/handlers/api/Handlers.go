package api

import (
	"github.com/gin-gonic/gin"
	"menu/src/di"
)

const limit = 100

func GetDishes(c *gin.Context, container *di.Di) {
	dishes, _ := container.Repositories.Dish.GetMany(limit)
	c.JSONP(200, dishes)
}

func GetIngredients(c *gin.Context, container *di.Di) {
	ingredients, _ := container.Repositories.Ingredient.GetMany(limit)
	c.JSONP(200, ingredients)
}

func GetIngredientUnits(c *gin.Context, container *di.Di) {
	ingredientUnits, _ := container.Repositories.IngredientUnits.GetMany(limit)
	c.JSONP(200, ingredientUnits)
}
