package kantenliste

import "slices"

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
	ids := []int{}
	for _, i := range g.Edges {
		ids = append(ids, i[0])
		ids = append(ids, i[1])
	}
	slices.Sort(ids)
	ids = slices.Compact(ids)

	return ids
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *EdgeListGraph) GetEdges() [][3]int {
	return g.Edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *EdgeListGraph) AddEdge(start, end, weight int) {
	a := [3]int{start, end, weight}
	g.Edges = append(g.Edges, a)
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *EdgeListGraph) RemoveEdge(start, end int) {
	for i, edge := range g.Edges {
		if edge[0] == start && edge[1] == end {
			g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
		}
	}
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *EdgeListGraph) GetNeighbors(node_id int) []int {
	neighbors := make([]int, 0)
	for a, heinz := range g.Edges {
		if heinz[0] == node_id {
			neighbors = append(neighbors, g.Edges[a][1])
		}
	}
	return neighbors
}
