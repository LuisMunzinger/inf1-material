package aufgabe4

import "slices"

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// ReachableNodes soll eine Liste aller von n aus erreichbaren Knoten liefern.
func (n *Node) ReachableNodes() []*Node {
	visitedNodes := []*Node{}
	newNodes := []*Node{n}

	for len(newNodes) > 0 {
		current := newNodes[len(newNodes)-1]
		visitedNodes = append(visitedNodes, current)
		newNodes = newNodes[:len(newNodes)-1]
		for _, neighbour := range current.neighbours {
			if !slices.Contains(visitedNodes, neighbour) {
				newNodes = append(newNodes, neighbour)
			}
		}
	}
	return visitedNodes
}
