package inline

// Tests für die Implementierung der Inline-Graph-Darstellung.
//
// Anmerkungen:
// - Der Datentyp `InlineGraph` implementiert das Interface `Graph`
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
	g1 := EmptyInlineGraph()

	// Prüfe, ob die Liste der Knoten leer ist.
	nodes := g1.GetNodes()
	graph.AssertNodesEmpty(nodes, t)

	// Gegeben ein Graph mit einigen Kanten.
	g2 := InlineGraph{
		Nodes: []struct {
			ID        int
			Neighbors []int
			Weights   []int
		}{
			{ID: 0, Neighbors: []int{1, 2}, Weights: []int{5, 3}},
			{ID: 1, Neighbors: []int{0}, Weights: []int{5}},
			{ID: 2, Neighbors: []int{0}, Weights: []int{3}},
		},
	}

	// Prüfe, ob die Liste der Knoten die erwarteten Knoten enthält.
	nodes = g2.GetNodes()
	expected := []int{0, 1, 2}
	graph.AssertNodesEqual(nodes, expected, t)
}

func TestGetEdges(t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := EmptyInlineGraph()

	// Prüfe, ob die Liste der Kanten leer ist.
	edges := g1.GetEdges()
	graph.AssertEdgesEqual(edges, [][3]int{}, t)

	// Gegeben ein Graph mit einigen Kanten.
	g2 := InlineGraph{
		Nodes: []struct {
			ID        int
			Neighbors []int
			Weights   []int
		}{
			{ID: 0, Neighbors: []int{1, 2}, Weights: []int{5, 3}},
			{ID: 1, Neighbors: []int{0}, Weights: []int{5}},
			{ID: 2, Neighbors: []int{0}, Weights: []int{3}},
		},
	}

	// Prüfe, ob die Liste der Kanten die erwarteten Kanten enthält.
	edges = g2.GetEdges()
	expected := [][3]int{{0, 1, 5}, {0, 2, 3}, {1, 0, 5}, {2, 0, 3}}
	graph.AssertEdgesEqual(edges, expected, t)
}

func TestAddEdge(t *testing.T) {
	graph.CheckAddEdge(EmptyInlineGraph, t)
}

func TestRemoveEdge(t *testing.T) {
	graph.CheckRemoveEdge(EmptyInlineGraph, t)
}

func TestGetNeighbors(t *testing.T) {
	graph.CheckGetNeighbors(EmptyInlineGraph, t)
}

func TestDegree(t *testing.T) {
	graph.CheckDegree(EmptyInlineGraph, t)
}

func TestHasEdge(t *testing.T) {
	graph.CheckHasEdge(EmptyInlineGraph, t)
}

func TestCheckGetEdgeWeight(t *testing.T) {
	graph.CheckGetEdgeWeight(EmptyInlineGraph, t)
}

func TestDotString(t *testing.T) {
	graph.CheckDotString(EmptyInlineGraph, t)
}
