package Duration

import (
	"fmt"
)

type Duration uint

// Duration als Sekunden festlegen
func Seconds(m int) Duration {
	return Duration(m)
}

//------------------------------------------------------------------------------------

// Minutes erwartet ein int, das als Minuten interpretiert wird.
// Liefert eine Duration für diese Minuten-Anzahl.(Anzahl der Sekunden)
func Minutes(m int) Duration {
	return Duration(m * 60)
}

func Houers(m int) Duration {
	return Duration(m * 60 * 60)
}

//-------------------------------------------------------------------------------------

// ToSeconds liefert die Duration als Secunden
func (s Duration) ToSeconds() int {
	return int(s)
}

// ToMinutes liefert die Duration als Minuten
func (s Duration) ToMinutes() int {
	return int(s) / 60
}

// ToHouers liefert die Duration als Stunden
func (s Duration) ToHouers() int {
	return int(s) / 60 / 60
}

//------------------Duration nicht mehr verwendet (müssen es nicht mehr wissen)----------------------------------

// SecondsToMinutes erwartet eine Sekunden-Anzahl
// und Liefert die entsprechende Minuten.
func SecondsToMinutes(s int) int {
	return int(Seconds(s).ToMinutes())
}

func SecondsToHouers(s int) int {
	return int(Seconds(s).ToHouers())
}

func Example() {
	fmt.Println(Seconds(60))
	fmt.Println(Minutes(60))
	fmt.Println(Houers(5))
	fmt.Println(SecondsToMinutes(120))
	fmt.Println(SecondsToHouers(3600))

	// Output:
	// 60
	// 3600
	// 18000
	// 2
	// 1

}
