package main

/*
	La fonction applyToIngredients est une fonction d'ordre supérieur idempotente qui prend une recette et
	une fonction en entrée, et applique la fonction à chaque ingrédient de la recette.
	Cette fonction est pure, car elle ne modifie pas l'objet Recipe original et renvoie
	une nouvelle instance modifiée.
*/
func applyToIngredients(r Recipe, f func(string) string) Recipe {
	for i, ingredient := range r.Ingredients {
		r.Ingredients[i] = f(ingredient)
	}
	return r
}
