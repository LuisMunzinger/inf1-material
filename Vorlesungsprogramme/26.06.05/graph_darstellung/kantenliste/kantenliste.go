package kantenliste

// EdgeListGraph repräsentiert einen gerichteten Graphen als Liste von Kanten.
// Jede Kante wird als Tripel (Startknoten, Zielknoten, Gewicht) dargestellt.
// Die Knoten sind dabei numerisch indiziert, z.B. 0, 1, 2, ...
type EdgeListGraph struct {
	Edges [][3]int
}

// EmptyEdgeListGraph erstellt einen neuen leeren EdgeListGraph.
func EmptyEdgeListGraph() *EdgeListGraph {
	return &EdgeListGraph{
		Edges: make([][3]int, 0),
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *EdgeListGraph) GetNodes() []int {
	// TODO
	return nil
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *EdgeListGraph) GetEdges() [][3]int {
	// TODO
	return nil
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *EdgeListGraph) AddEdge(start, end, weight int) {
	// TODO

}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *EdgeListGraph) RemoveEdge(start, end int) {
	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *EdgeListGraph) GetNeighbors(node_id int) []int {
	neighbors := make([]int, 0)
	// TODO
	return neighbors
}
