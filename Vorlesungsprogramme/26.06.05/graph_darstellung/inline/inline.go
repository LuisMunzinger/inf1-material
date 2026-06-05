package inline

// InlineGraph ist eine Implementierung des Graph-Interfaces, die
// die eine Liste von Knoten enthält, wobei jeder Knoten seine Nachbarn
// und die Gewichte der Kanten zu diesen Nachbarn direkt speichert.
type InlineGraph struct {
	// Nodes speichert die Knoten des Graphen. Jeder Eintrag enthält die ID des Knotens,
	// eine Liste der Nachbarn und die entsprechenden Gewichte der Kanten zu diesen Nachbarn.
	Nodes []struct {
		ID        int
		Neighbors []int
		Weights   []int
	}
}

// EmptyInlineGraph erstellt einen neuen leeren InlineGraph.
func EmptyInlineGraph() *InlineGraph {
	return &InlineGraph{
		Nodes: []struct {
			ID        int
			Neighbors []int
			Weights   []int
		}{},
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *InlineGraph) GetNodes() []int {
	var nodes []int
	// TODO
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *InlineGraph) GetEdges() [][3]int {
	var edges [][3]int
	// TODO
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *InlineGraph) AddEdge(start, end, weight int) {
	// TODO
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *InlineGraph) RemoveEdge(start, end int) {
	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *InlineGraph) GetNeighbors(node_id int) []int {
	// TODO
	return []int{}
}
