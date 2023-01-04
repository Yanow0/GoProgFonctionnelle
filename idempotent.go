package main

/*
	La fonction addIngredient est une fonction pure idempotente qui ajoute
	un ingrédient à une recette sans modifier l'original.
*/
func addIngredient(r Recipe, ingredient string) Recipe {
	newIngredients := append(r.Ingredients, ingredient)
	return Recipe{
		Name:         r.Name,
		Ingredients:  newIngredients,
		Instructions: r.Instructions,
	}
}
