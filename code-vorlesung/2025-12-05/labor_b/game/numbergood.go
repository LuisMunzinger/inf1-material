package game

import "math/rand"

// Numbergood liefert true, falls die geratene Zahl richtig ist.
func NumberGood(g int) bool {
	return g == rand.Intn(100)+1
}
