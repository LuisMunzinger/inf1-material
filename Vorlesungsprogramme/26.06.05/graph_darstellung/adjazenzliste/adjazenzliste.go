package adjazenzliste

// AdjacencyListGraph repräsentiert einen gerichteten Graphen als Adjazenzliste.
// Jeder Eintrag in der Liste enthält die Nachbarn eines Knotens
// und die Gewichteder Kanten.
type AdjacencyListGraph struct {
	AdjList map[int][][2]int // Knoten-ID -> Liste von (Nachbar-Knoten-ID, Gewicht)
}

// EmptyAdjacencyListGraph erstellt einen neuen leeren AdjacencyListGraph.
func EmptyAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		AdjList: make(map[int][][2]int),
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *AdjacencyListGraph) GetNodes() []int {
	nodes := make([]int, 0, len(g.AdjList))
	// TODO
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *AdjacencyListGraph) GetEdges() [][3]int {
	edges := make([][3]int, 0)
	// TODO
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *AdjacencyListGraph) AddEdge(start, end, weight int) {
	// TODO
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *AdjacencyListGraph) RemoveEdge(start, end int) {
	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *AdjacencyListGraph) GetNeighbors(node_id int) []int {
	neighbors := make([]int, 0)
	// TODO
	return neighbors
}
