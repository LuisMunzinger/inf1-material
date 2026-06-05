package graph

// Degree gibt die Anzahl der ausgehenden Kanten eines Knotens zurück.
func Degree(g GraphRepr, node_id int) int {
	// TODO
	return 0
}

// HasEdge prüft, ob eine gerichtete Kante von start zu end existiert.
func HasEdge(g GraphRepr, start, end int) bool {
	// TODO
	return false
}

// GetEdgeWeight gibt das Gewicht der gerichteten Kante von start zu end zurück,
// oder 0, wenn die Kante nicht existiert.
func GetEdgeWeight(g GraphRepr, start, end int) int {
	// TODO
	return 0
}

// DotString gibt eine String-Darstellung des Graphen im DOT-Format zurück,
// die zur Visualisierung mit Graphviz verwendet werden kann.
func DotString(g GraphRepr) string {
	var result string
	// TODO
	return result
}
