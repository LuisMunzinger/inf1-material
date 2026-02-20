package hanoi_v2

import "fmt"

/*
s = Start
m = Mitte
z = Ziel
*/

var Counter int = 0

func Move(s, z string) {
	Counter++
	fmt.Printf("%s -> % s\n", s, z)
}

//Löst die Türme von Hanoi
func Hanoi(h int, s, m, z string) {
	if h == 0 {
		return
	}
	Hanoi(h-1, s, z, m)
	Move(s, z)
	Hanoi(h-1, m, s, z)
}
