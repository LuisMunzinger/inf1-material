package hanoi

import "fmt"

/*
s = Start
m = Mitte
z = Ziel
*/

func Move(s, z string) {
	fmt.Printf("%s -> % s\n", s, z)
}

//Bewegt einen Turm der höhe 1 von s nach z
func Hanoi1(s, z string) {
	Move(s, z)
}

//Bewegt einen Turm der höhe 2 von s auf z
func Hanoi2(s, m, z string) {
	Move(s, m) //Parke die oberste Platte auf m
	Move(s, z) //Bewege die unterste platte auf z
	Move(m, z) //Bewege die gepakte Platte von m auf z
}

//Bewegt einen Turm der höhe 3 von s auf z
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m) //Parke einen Turm der höhe 2 auf m
	Move(s, z)      //Bewege die unterste platte auf z
	Hanoi2(m, s, z) //Bewege die gepakte Platte von m auf z
}

//Bewegt einen Turm der höhe 4 von s auf z
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m) //Parke einen Turm der höhe 3 auf m
	Move(s, z)      //Bewege die unterste platte auf z
	Hanoi3(m, s, z) //Bewege die gepakte Platte von m auf z
}

//Bewegt einen Turm der höhe 4 von s auf z
func Hanoi5(s, m, z string) {
	Hanoi4(s, z, m) //Parke einen Turm der höhe 4 auf m
	Move(s, z)      //Bewege die unterste platte auf z
	Hanoi4(m, s, z) //Bewege die gepakte Platte von m auf z
}
