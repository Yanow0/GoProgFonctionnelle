package main

/*
	La fonction de composition de fonctions permet de combiner plusieurs fonctions en une seule.
	Dans cet exemple, la fonction de composition de fonctions est utilisée pour appliquer deux
	ingrédients à une recette.
*/
func compose(f, g func(Recipe) Recipe) func(Recipe) Recipe {
	return func(r Recipe) Recipe {
		return f(g(r))
	}
}
