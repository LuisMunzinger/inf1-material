package adjazenzmatrix

// Tests für die Implementierung der Adjazenzmatrix als Graph-Darstellung.
//
// Anmerkungen:
// - Der Datentyp `AdjacencyMatrixGraph` implementiert das Interface `Graph`
//   aus dem Package `graph`. D.h. die Methoden und ihre Rückgabetypen sind größtenteils
//   durch das Interface vorgegeben.
// - Die Tests sollten für alle Implementierungen des `Graph`-Interfaces funktionieren
//   bzw. zu großen Teilen wiederverwendbar sein.
// - Deshalb greifen die Tests auf Hilfsfunktionen zurück, die im Package `graph`
//   definiert sind.
// - Manche Tests werden sogar vollständig an solche Hilfsfunktionen delegiert.
// - Diese Vorgehensweise vermeidet Redundanz und stellt sicher, dass die
//   verschiedenen Implementierungen konsistent sind.

import (
	"inf1-material/Vorlesungsprogramme/26.06.05/graph"
	"testing"
)

func TestGetNodes(t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := EmptyAdjacencyMatrixGraph(0)

	// Prüfe, ob die Liste der Knoten leer ist.
	nodes := g1.GetNodes()
	graph.AssertNodesEmpty(nodes, t)

	// Gegeben ein Graph mit einigen Kanten.
	g2 := AdjacencyMatrixGraph{
		Edges: [][]int{
			{0, 5, 3, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	// Prüfe, ob die Liste der Knoten die erwarteten Knoten enthält.
	nodes = g2.GetNodes()
	expected := []int{0, 1, 2, 3, 4}
	graph.AssertNodesEqual(nodes, expected, t)
}

func TestGetEdges(t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := EmptyAdjacencyMatrixGraph(0)

	// Prüfe, ob die Liste der Kanten leer ist.
	edges := g1.GetEdges()
	graph.AssertEdgesEqual(edges, [][3]int{}, t)

	// Gegeben ein Graph mit einigen Kanten.
	g2 := AdjacencyMatrixGraph{
		Edges: [][]int{
			{0, 5, 3, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 2},
			{4, 0, 0, 0, 0},
		},
	}

	// Prüfe, ob die Liste der Kanten die erwarteten Kanten enthält.
	edges = g2.GetEdges()
	expected := [][3]int{{0, 1, 5}, {0, 2, 3}, {1, 2, 1}, {3, 4, 2}, {4, 0, 4}}
	graph.AssertEdgesEqual(edges, expected, t)
}

func EmptyGraph() graph.GraphRepr {
	return EmptyAdjacencyMatrixGraph(5)
}

func TestAddEdge(t *testing.T) {
	graph.CheckAddEdge(EmptyGraph, t)
}

func TestRemoveEdge(t *testing.T) {
	graph.CheckRemoveEdge(EmptyGraph, t)
}

func TestGetNeighbors(t *testing.T) {
	graph.CheckGetNeighbors(EmptyGraph, t)
}

func TestDegree(t *testing.T) {
	graph.CheckDegree(EmptyGraph, t)
}

func TestHasEdge(t *testing.T) {
	graph.CheckHasEdge(EmptyGraph, t)
}

func TestCheckGetEdgeWeight(t *testing.T) {
	graph.CheckGetEdgeWeight(EmptyGraph, t)
}

func TestDotString(t *testing.T) {
	graph.CheckDotString(EmptyGraph, t)
}
