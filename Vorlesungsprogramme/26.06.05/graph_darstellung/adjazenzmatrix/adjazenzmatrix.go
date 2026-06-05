package adjazenzmatrix

// AdjacencyMatrixGraph ist eine Implementierung des Graph-Interfaces, die eine Adjazenzmatrix verwendet, um die Kanten und ihre Gewichte zu speichern.
type AdjacencyMatrixGraph struct {
	// Edges speichert die Gewichte der Kanten zwischen den Knoten.
	// matrix[i][j] enthält das Gewicht der Kante von Knoten i zu Knoten j, oder 0, wenn keine Kante existiert.
	Edges [][]int
}

// EmptyAdjacencyMatrixGraph erstellt einen neuen leeren AdjacencyMatrixGraph mit einer vorgegebenen Anzahl von Knoten.
func EmptyAdjacencyMatrixGraph(numNodes int) *AdjacencyMatrixGraph {
	edges := make([][]int, numNodes)
	for i := range edges {
		edges[i] = make([]int, numNodes)
	}
	return &AdjacencyMatrixGraph{
		Edges: edges,
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *AdjacencyMatrixGraph) GetNodes() []int {
	nodes := make([]int, len(g.Edges))
	// TODO
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *AdjacencyMatrixGraph) GetEdges() [][3]int {
	edges := make([][3]int, 0)
	// TODO
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *AdjacencyMatrixGraph) AddEdge(start, end, weight int) {
	// TODO
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *AdjacencyMatrixGraph) RemoveEdge(start, end int) {
	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *AdjacencyMatrixGraph) GetNeighbors(node_id int) []int {
	neighbors := make([]int, 0)
	// TODO
	return neighbors
}
