package graph

import (
	"slices"
	"testing"
)

// AssertNodesEmpty überprüft, ob eine Liste von Knoten-IDs leer ist.
func AssertNodesEmpty(nodes []int, t *testing.T) {
	if !slices.Equal(nodes, []int{}) {
		t.Errorf("GetNodes: Erwartete leere Knotenliste, tatsächliche Knoten: %v", nodes)
	}
}

// AssertEdgesEmpty überprüft, ob eine Liste von Kanten leer ist.
func AssertEdgesEmpty(edges [][3]int, t *testing.T) {
	if !slices.Equal(edges, [][3]int{}) {
		t.Errorf("GetEdges: Erwartete leere Kantenliste, tatsächliche Kanten: %v", edges)
	}
}

// AssertNodesEqual ist eine Hilfsfunktion, die zwei Listen von Knoten-IDs vergleicht.
// Die Funktion kopiert und sortiert die Listen vorher, damit die Reihenfolge der
// Knoten keine Rolle spielt.
func AssertNodesEqual(a, b []int, t *testing.T) {
	if !NodesEqual(a, b) {
		t.Errorf("GetNodes: Erwartete Knoten %v, tatsächliche Knoten: %v", b, a)
	}
}

// AssertEdgesEqual ist eine Hilfsfunktion, die zwei Listen von Kanten vergleicht.
// Die Funktion kopiert und sortiert die Listen vorher, damit die Reihenfolge der
// Kanten keine Rolle spielt.
func AssertEdgesEqual(a, b [][3]int, t *testing.T) {
	if !EdgesEqual(a, b) {
		t.Errorf("GetEdges: Erwartete Kanten %v, tatsächliche Kanten: %v", b, a)
	}
}

// NodesEqual ist eine Hilfsfunktion, die zwei Listen von Knoten-IDs vergleicht.
// Die Funktion kopiert und sortiert die Listen vorher, damit die Reihenfolge der
// Knoten keine Rolle spielt.
func NodesEqual(a, b []int) bool {
	aCopy := make([]int, len(a))
	copy(aCopy, a)
	slices.Sort(aCopy)

	bCopy := make([]int, len(b))
	copy(bCopy, b)
	slices.Sort(bCopy)

	return slices.Equal(aCopy, bCopy)
}

// EdgesEqual ist eine Hilfsfunktion, die zwei Listen von Kanten vergleicht.
// Die Funktion kopiert und sortiert die Listen vorher, damit die Reihenfolge der
// Kanten keine Rolle spielt.
func EdgesEqual(a, b [][3]int) bool {
	aCopy := make([][3]int, len(a))
	copy(aCopy, a)
	cmp := func(x, y [3]int) int {
		return slices.Compare(x[:], y[:])
	}
	slices.SortFunc(aCopy, cmp)

	bCopy := make([][3]int, len(b))
	copy(bCopy, b)
	slices.SortFunc(bCopy, cmp)

	return slices.Equal(aCopy, bCopy)
}
