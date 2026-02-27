package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {
	result := 0
	if len(list) == 0 {
		return 0
	}
	wert := list[0]
	if wert%2 == 1 {
		result++
	}
	return result + CountOdd(list[1:])
}
