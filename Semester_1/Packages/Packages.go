package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// Ein Slice mit Zahlen
	numbers := []int{5, 2, 8, 1, 9}

	fmt.Println("")
	fmt.Println("Original Slice:", numbers)
	fmt.Println("")

	// Sortieren
	slices.Sort(numbers)
	fmt.Println("Sortiert:", numbers)

	// Prüfen, ob ein Wert enthalten ist
	contains := slices.Contains(numbers, 8)
	fmt.Println("Enthält 8?", contains)

	// Index eines Wertes finden
	index := slices.Index(numbers, 8)
	fmt.Println("wo steht die 8 in der liste:", index)

	// Element löschen (Index 1 entfernen)
	numbers = slices.Delete(numbers, 1, 2)
	fmt.Println("Nach Löschen des Elements an Index 1:", numbers)

	// Vergleich mit einem anderen Slice
	other := []int{1, 5, 8, 9}
	equal := slices.Equal(numbers, other)
	fmt.Println("Ist gleich mit", other, "?", equal)

	cloned := slices.Clone(numbers)
	fmt.Println("Clone:", cloned)

	// Min / Max
	fmt.Println("Min:", slices.Min(numbers))
	fmt.Println("Max:", slices.Max(numbers))

	// Repeat (Slice 3x wiederholen)
	repeated := slices.Repeat(numbers, 3)
	fmt.Println("Repeat x3:", repeated)

	// Replace (Index 1–3 ersetzen durch neue Werte)
	replaced := slices.Replace(numbers, 1, 3, 100, 200)
	fmt.Println("Replace Index 1-3 mit 100,200:", replaced)

	// Concat (zwei Slices verbinden)
	other3 := []int{50, 60}
	concatenated := slices.Concat(numbers, other3)
	fmt.Println("Concat mit", other3, ":", concatenated)

	// Insert (Elemente einfügen an Position 2)
	inserted := slices.Insert(numbers, 2, 999)
	fmt.Println("Insert 999 an Index 2:", inserted)

	fmt.Println("")
	fmt.Println("")

	//------------------------------------------------------------------------------------------------
	// String-Slice mit unsauberen Eingaben
	words := []string{"  Apfel  ", "BaNaNE", "  Kirsche", "mango  "}

	fmt.Println("Original:", words)
	fmt.Println("")

	// Trim anwenden (Leerzeichen entfernen)
	for i, w := range words {
		words[i] = strings.Trim(w, " ")
	}
	fmt.Println("Getrimmt:", words)

	// Prüfen auf Prefix und Suffix
	fmt.Println("\n--- Prefix / Suffix Prüfung ---")
	for _, w := range words {
		fmt.Printf("%s -> HasPrefix 'A'? %v | HasSuffix 'E'? %v\n",
			w,
			strings.HasPrefix(w, "A"),
			strings.HasSuffix(w, "E"),
		)
	}

	text := "go ist toll, go ist schnell, go macht spaß"
	fmt.Println("Text:", text)

	count := strings.Count(text, "go")
	fmt.Println("\nCount 'go':", count)

	// Index
	index1 := strings.Index(text, "schnell")
	fmt.Println("Index von 'schnell':", index1)

	// Cut
	before, after, found := strings.Cut(text, ",")
	fmt.Println("\nCut bei erstem Komma:")
	fmt.Println("Vorne:", before)
	fmt.Println("Hinten:", after)
	fmt.Println("Found:", found)

	// CutPrefix
	prefixText := "Hallo Welt"
	withoutPrefix, ok := strings.CutPrefix(prefixText, "Hallo ")
	fmt.Println("\nCutPrefix 'Hallo ' aus 'Hallo Welt':")
	fmt.Println("Ergebnis:", withoutPrefix)
	fmt.Println("Gefunden:", ok)

	// Split
	words2 := strings.Split(text, "e")
	fmt.Println("\nSplit nach Leerzeichen:", words2)

	// Join
	joined := strings.Join(words, "-")
	fmt.Println("Join mit '-':", joined)
}
