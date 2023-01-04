package main

import (
	"fmt"
	"strings"
)

type Recipe struct {
	Name         string
	Ingredients  []string
	Instructions string
}

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

//TODO: fonction compose

func main() {
	/*
		L'expression lambda uppercaseIngredients est une fonction qui met en majuscules
		une chaîne de caractères.
	*/
	uppercaseIngredients := func(s string) string {
		return strings.ToUpper(s)
	}

	r := Recipe{
		Name:         "Chocolate Cake",
		Ingredients:  []string{"flour", "sugar", "cocoa powder", "eggs"},
		Instructions: "Mix ingredients together and bake in a 350 degree oven for 30 minutes",
	}

	fmt.Println("Original recipe:")
	fmt.Println(r)

	fmt.Println("\nRecipe with uppercase ingredients:")
	uppercaseRecipe := applyToIngredients(r, uppercaseIngredients)
	fmt.Println(uppercaseRecipe)

	fmt.Println("\nModified recipe:")
	modifiedRecipe := addIngredient(r, "vanilla extract")
	fmt.Println(modifiedRecipe)

}
