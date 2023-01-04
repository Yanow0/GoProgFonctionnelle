package main

import (
	"fmt"
	"strings"
)

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

	addSalt := func(r Recipe) Recipe {
		r.Ingredients = append(r.Ingredients, "SALT")
		return r
	}
	addYeast := func(r Recipe) Recipe {
		r.Ingredients = append(r.Ingredients, "YEAST")
		return r
	}
	addSaltAndYeast := compose(addSalt, addYeast)
	fmt.Println("\nRecipe with salt and yeast:")
	fmt.Println(addSaltAndYeast(r))
}
