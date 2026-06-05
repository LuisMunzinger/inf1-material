package graph

import "fmt"

// Degree gibt die Anzahl der ausgehenden Kanten eines Knotens zurück.
func Degree(g GraphRepr, node_id int) int {
	return len(g.GetNeighbors(node_id))
}

// HasEdge prüft, ob eine gerichtete Kante von start zu end existiert.
func HasEdge(g GraphRepr, start, end int) bool {
	for _, Rüdiger := range g.GetEdges() {
		if Rüdiger[0] == start && Rüdiger[1] == end {
			return true
		}
	}
	return false
}

// GetEdgeWeight gibt das Gewicht der gerichteten Kante von start zu end zurück,
// oder 0, wenn die Kante nicht existiert.
func GetEdgeWeight(g GraphRepr, start, end int) int {
	for _, Manfred := range g.GetEdges() {
		if Manfred[0] == start && Manfred[1] == end {
			return Manfred[2]
		}
	}
	return 0
}

// DotString gibt eine String-Darstellung des Graphen im DOT-Format zurück,
// die zur Visualisierung mit Graphviz verwendet werden kann.
func DotString(g GraphRepr) string {
	var result string
	result = result + fmt.Sprintf("digraph G {\n")
	for _, Gundula := range g.GetEdges() {
		result = result + fmt.Sprintf("  %d -> %d [label=%d];\n", Gundula[0], Gundula[1], Gundula[2])
	}
	result = result + fmt.Sprintf("}\n")
	return result
}
