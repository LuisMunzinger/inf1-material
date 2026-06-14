package aufgabe5

import "fmt"

func ExampleNode_PathStrings() {
	n := NewNode()

	n.Add(100)
	n.Add(50)
	n.Add(25)
	n.Add(75)
	n.Add(150)

	paths := n.PathStrings()
	for _, path := range paths {
		fmt.Println(path)
	}

	// Output:
	// ll
	// lr
	// r
}

/* Anmerkung: Der obige Baum sieht folgendermaßen aus:

		    100
		   /   \
		  /     \
	     50     150
	    /  \
	   25  75

Daher gibt es drei Pfade, von der Wurzel zu je einem Blatt.
*/
