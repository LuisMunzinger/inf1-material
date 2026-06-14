package aufgabe5

import "fmt"

func ExampleNode_Size() {
	n := NewNode()

	n.Add(100)
	n.Add(50)
	n.Add(25)
	n.Add(75)
	n.Add(150)

	fmt.Println(n.Size())

	// Output:
	// 5
}

/* Anmerkung: Der obige Baum sieht folgendermaßen aus:

	    100
	   /   \
	  /     \
     50     150
    /  \
   25  75
*/
