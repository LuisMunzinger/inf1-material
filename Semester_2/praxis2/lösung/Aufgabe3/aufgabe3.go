package aufgabe3

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// HasNeighbour prüft, ob n einen Nachbarn mit dem angegebenen Label hat.
func (n *Node) HasNeighbour(label string) bool {
	for _, neighbour := range n.neighbours {
		if neighbour.Label == label {
			return true
		}
	}
	return false
}
