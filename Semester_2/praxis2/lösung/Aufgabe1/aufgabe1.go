package aufgabe1

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// Sum gibt die Summe aller Werte in der Liste zurück.
// Wenn die Liste leer ist, soll 0 zurückgegeben werden.
func (n *Node) Sum() int {
	sum := 0
	for !n.IsEmpty() {
		sum += n.Value
		n = n.Next
	}
	return sum
}
