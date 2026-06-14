package aufgabe4

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// NodesWithDistance soll eine Liste aller Knoten liefern, die in genau d Schritten
// von n aus erreichbar sind.
// Die Liste soll keine Duplikate enthalten, sie muss aber nicht sortiert sein.
// Der Knoten n selbst ist in 0 Schritten von n aus erreichbar,
// seine direkten Nachbarn sind in 1 Schritt erreichbar, usw.
func (n *Node) NodesWithDistance(d int) []*Node {
	result := []*Node{n}
	// TODO
	return result
}
