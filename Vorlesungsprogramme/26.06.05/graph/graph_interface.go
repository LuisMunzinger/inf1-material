package graph

// GraphRepr ist ein Interface, das die grundlegenden Operationen definiert,
// die auf einem Graphen ausgeführt werden können.
type GraphRepr interface {

	// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
	GetNodes() []int

	// GetEdges gibt die Liste aller Kanten im Graphen zurück.
	// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
	GetEdges() [][3]int

	// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
	AddEdge(start, end, weight int)

	// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
	RemoveEdge(start, end int)

	// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
	// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
	GetNeighbors(node_id int) []int
}
