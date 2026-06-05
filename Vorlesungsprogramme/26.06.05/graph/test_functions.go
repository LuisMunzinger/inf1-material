package graph

import (
	"testing"
)

// CheckAddEdge überprüft, ob die Methode AddEdge korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktion `GetEdges` muss bereits funktionieren.
func CheckAddEdge[G GraphRepr](emptyGraph func() G, t *testing.T) {
	// Gegeben ein leerer Graph (ohne Kanten).
	g := emptyGraph()

	// Füge zwei Kanten hinzu.
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 3)

	// Überprüfe, ob die Kantenliste die hinzugefügten Kanten enthält.
	edges := g.GetEdges()
	expected := [][3]int{{0, 1, 5}, {0, 2, 3}, {1, 2, 3}}

	AssertEdgesEqual(edges, expected, t)
}

// CheckRemoveEdge überprüft, ob die Methode RemoveEdge korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktionen `AddEdge` und `GetEdges` müssen bereits funktionieren.
func CheckRemoveEdge[G GraphRepr](emptyGraph func() G, t *testing.T) {
	g := emptyGraph()

	// Gegeben ein Graph mit zwei Knoten, die in beide Richtungen
	// durch Kanten verbunden sind.
	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 3)

	// Entferne eine der Kanten.
	g.RemoveEdge(0, 1)

	// Überprüfe, ob die Kantenliste nur noch die Kante enthält, die nicht entfernt wurde.
	edges := g.GetEdges()
	expected := [][3]int{{1, 2, 3}}

	AssertEdgesEqual(edges, expected, t)
}

// CheckGetNeighbors überprüft, ob die Methode GetNeighbors korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktion `AddEdge` muss bereits funktionieren.
func CheckGetNeighbors[G GraphRepr](emptyGraph func() G, t *testing.T) {
	g := emptyGraph()

	// Gegeben ein Graph mit Kanten von 0 zu 1, von 0 zu 2 und von 1 zu 2.
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)

	// Bestimme die Nachbar-Listen der Knoten 0 und 1.
	neighbors0 := g.GetNeighbors(0)
	neighbors1 := g.GetNeighbors(1)

	// Erstelle zwei Listen mit den erwarteten Nachbarlisten.
	expected0 := []int{1, 2}
	expected1 := []int{2}

	// Prüfe, ob die Nachbar-Listen die erwarteten Nachbarn enthalten.
	AssertNodesEqual(neighbors0, expected0, t)
	AssertNodesEqual(neighbors1, expected1, t)

}

// CheckDegree überprüft, ob die Funktion Degree korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktion `AddEdge` muss bereits funktionieren.
func CheckDegree[G GraphRepr](emptyGraph func() G, t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := emptyGraph()

	// Prüfe, ob der Grad eines Knotens in einem leeren Graphen 0 ist.
	d1 := Degree(g1, 0)
	if d1 != 0 {
		t.Errorf("Degree: Erwarteter Grad von Knoten 0 ist 0, tatsächlicher Wert %d", d1)
	}

	// Gegeben ein Graph mit einigen Kanten.
	g2 := emptyGraph()
	g2.AddEdge(0, 1, 5)
	g2.AddEdge(0, 2, 3)
	g2.AddEdge(1, 2, 1)

	// Prüfe, ob der Grad der Knoten die erwarteten Werte hat.
	expectedDegrees := map[int]int{0: 2, 1: 1, 2: 0}
	for node, expectedDegree := range expectedDegrees {
		if Degree(g2, node) != expectedDegree {
			t.Errorf("Degree: Erwarteter Grad von Knoten %d ist %d, tatsächlicher Wert %d", node, expectedDegree, Degree(g2, node))
		}
	}
}

// CheckHasEdge überprüft, ob die Funktion HasEdge korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktion `AddEdge` muss bereits funktionieren.
func CheckHasEdge[G GraphRepr](emptyGraph func() G, t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := emptyGraph()

	// Prüfe, ob HasEdge für einen leeren Graphen immer false zurückgibt.
	if HasEdge(g1, 0, 1) {
		t.Errorf("HasEdge: Erwarte false für Kante (0, 1) in leerem Graphen, habe aber true erhalten")
	}

	// Gegeben ein Graph mit einigen Kanten.
	g2 := emptyGraph()
	g2.AddEdge(0, 1, 5)
	g2.AddEdge(0, 2, 3)
	g2.AddEdge(1, 2, 1)

	// Prüfe, ob HasEdge die erwarteten Werte zurückgibt.
	expectedEdges := [][2]int{{0, 1}, {0, 2}, {1, 2}}
	for _, edge := range expectedEdges {
		if !HasEdge(g2, edge[0], edge[1]) {
			t.Errorf("HasEdge: Erwarte true für Kante (%d, %d), habe aber false erhalten", edge[0], edge[1])
		}
	}

	// Prüfe, ob HasEdge für nicht vorhandene Kanten false zurückgibt.
	if HasEdge(g2, 1, 0) {
		t.Errorf("HasEdge: Erwarte false für Kante (1, 0), habe aber true erhalten")
	}
	if HasEdge(g2, 2, 0) {
		t.Errorf("HasEdge: Erwarte false für Kante (2, 0), habe aber true erhalten")
	}
	if HasEdge(g2, 2, 1) {
		t.Errorf("HasEdge: Erwarte false für Kante (2, 1), habe aber true erhalten")
	}
}

// CheckGetEdgeWeight überprüft, ob die Funktion GetEdgeWeight korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktion `AddEdge` muss bereits funktionieren.
func CheckGetEdgeWeight[G GraphRepr](emptyGraph func() G, t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := emptyGraph()

	// Prüfe, ob GetEdgeWeight für einen leeren Graphen immer 0 zurückgibt.
	if GetEdgeWeight(g1, 0, 1) != 0 {
		t.Errorf("GetEdgeWeight: Erwarte 0 für Kante (0, 1) in leerem Graphen, habe aber %d erhalten", GetEdgeWeight(g1, 0, 1))
	}

	// Gegeben ein Graph mit einigen Kanten.
	g2 := emptyGraph()
	g2.AddEdge(0, 1, 5)
	g2.AddEdge(0, 2, 3)
	g2.AddEdge(1, 2, 1)

	// Prüfe, ob GetEdgeWeight die erwarteten Werte zurückgibt.
	expectedWeights := map[[2]int]int{
		{0, 1}: 5,
		{0, 2}: 3,
		{1, 2}: 1,
	}
	for edge, expectedWeight := range expectedWeights {
		if GetEdgeWeight(g2, edge[0], edge[1]) != expectedWeight {
			t.Errorf("GetEdgeWeight: Erwartet %d für Kante (%d, %d), tatsächlicher Wert %d", expectedWeight, edge[0], edge[1], GetEdgeWeight(g2, edge[0], edge[1]))
		}
	}

	// Prüfe, ob GetEdgeWeight für nicht vorhandene Kanten 0 zurückgibt.
	if GetEdgeWeight(g2, 1, 0) != 0 {
		t.Errorf("GetEdgeWeight: Erwarte 0 für Kante (1, 0), habe aber %d erhalten", GetEdgeWeight(g2, 1, 0))
	}
	if GetEdgeWeight(g2, 2, 0) != 0 {
		t.Errorf("GetEdgeWeight: Erwarte 0 für Kante (2, 0), habe aber %d erhalten", GetEdgeWeight(g2, 2, 0))
	}
	if GetEdgeWeight(g2, 2, 1) != 0 {
		t.Errorf("GetEdgeWeight: Erwarte 0 für Kante (2, 1), habe aber %d erhalten", GetEdgeWeight(g2, 2, 1))
	}
}

// CheckDotString überprüft, ob die Funktion DotString korrekt implementiert ist.
// Voraussetzungen:
// - Die Funktionen `AddEdge` und `GetEdges` müssen bereits funktionieren.
func CheckDotString[G GraphRepr](emptyGraph func() G, t *testing.T) {
	// Gegeben ein leerer Graph.
	g1 := emptyGraph()

	// Prüfe, ob DotString für einen leeren Graphen den erwarteten String zurückgibt.
	expectedEmpty := "digraph G {\n}\n"
	if DotString(g1) != expectedEmpty {
		t.Errorf("DotString: Erwarte %q für leeren Graphen, habe aber %q erhalten", expectedEmpty, DotString(g1))
	}

	// Gegeben ein Graph mit einigen Kanten.
	g2 := emptyGraph()
	g2.AddEdge(0, 1, 5)
	g2.AddEdge(0, 2, 3)
	g2.AddEdge(1, 2, 1)

	// Prüfe, ob DotString den erwarteten String zurückgibt.
	expected := "digraph G {\n  0 -> 1 [label=5];\n  0 -> 2 [label=3];\n  1 -> 2 [label=1];\n}\n"
	if DotString(g2) != expected {
		t.Errorf("DotString: Erwarte %q, habe aber %q erhalten", expected, DotString(g2))
	}
}
