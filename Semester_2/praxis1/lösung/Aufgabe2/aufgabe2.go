package aufgabe2

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// HINWEIS: Zusätzlich zu den üblichen Feldern haben diese Listenelemente auch ihre Länge gespeichert.
//          Beim Entfernen eines Elements muss auch die Länge der Liste angepasst werden.
// MAX. PUNKTE: 10

// RemoveAt entfernt das Element an der Stelle index aus der Liste.
// Wenn index eine ungültige Position ist, soll die Liste unverändert bleiben.
func (n *Node) RemoveAt(index int) {
	if index < 0 || n.IsEmpty() {
		return
	}
	if index == 0 {
		n.Value = n.Next.Value
		n.Next = n.Next.Next
		n.Length--
		return
	}
	if index == 1 {
		n.Next = n.Next.Next
		n.Length--
		return
	}
	nextlen := n.Next.Length
	n.Next.RemoveAt(index - 1)
	if nextlen != n.Next.Length {
		n.Length--
	}
}
