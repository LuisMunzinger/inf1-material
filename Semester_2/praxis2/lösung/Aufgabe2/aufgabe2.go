package aufgabe2

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// HINWEIS: Zusätzlich zu den üblichen Feldern haben diese Listenelemente auch ihre Länge gespeichert.
//          Beim Entfernen eines Elements muss auch die Länge der Liste angepasst werden.
// MAX. PUNKTE: 10

// RemoveAll entfernt alle Elemente mit dem Wert value aus der Liste.
// Wenn value in der Liste nicht vorkommt, soll die Liste unverändert bleiben.
// Die Funktion liefert den neuen Kopf der Liste zurück.
func (n *Node) RemoveAll(value int) *Node {
	if n.IsEmpty() {
		return n
	}

	tail_removed := n.Next.RemoveAll(value)

	if n.Value == value {
		return tail_removed
	}

	n.Next = tail_removed
	n.Length = tail_removed.Length + 1

	return n
}
