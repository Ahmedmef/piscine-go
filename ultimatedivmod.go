package piscine

func UltimateDivMod(a *int, b *int) {
	g := *a % *b
	*a = *a / *b
	*b = g
}
