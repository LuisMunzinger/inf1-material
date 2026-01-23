package Quicksorte

import "fmt"

//soriere die Liste l mit dem Verfahren Quicksorte
func Qsort(l []int){

	//sonderfall: Die Liste ist Leer oder hat nur eine Element
	if len(l) <= 1{
		return
	}
	
	pivot := l[0]

	left := []int{}
	right := []int{}

	//Pationieren der Liste:
	//Kleinere Elemente als das Pivot nach Links, größere nach rechts.
	for _, el := range l[1:] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	Qsort(left)
	Qsort(right)

	//Elemente in die ursprüngliche Liste zurückkopieren
	if i, el := range left {
		l[i] = el
	}
	l[len(left)] = pivot
	if i, el := range right {
		l[i+len(left)+1] = el
	}
}

func ExampleQsort() {
	l1 := []int{17,25,22,3,15,4,35,105,42,1}

	Qsort(l1)

	fmt.Println(l1)
	
	// Output:
	// [1 3 4 15 17 22 25 35 42 105]
}