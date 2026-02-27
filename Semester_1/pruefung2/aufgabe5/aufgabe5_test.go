package aufgabe5

import "fmt"

func ExampleReplaceEn() {
	dict1 := []Entry{{"Haus", "house"}, {"Holz", "wood"}, {"Fahrrad", "bicycle"}, {"Holz", "wood"}, {"Wald", "wood"}}
	dict2 := []Entry{{"Haus", "house"}, {"Holz", "wood"}, {"Fahrrad", "bicycle"}, {"Holz", "wood"}, {"Wald", "wood"}}

	ReplaceEn(dict1, "Holz", "timber")
	ReplaceEn(dict1, "Maus", "mouse")

	fmt.Println(dict1)
	fmt.Println(dict2)

	// Output:
	// [{Haus house} {Holz timber} {Fahrrad bicycle} {Holz wood} {Wald wood}]
	// [{Haus house} {Holz wood} {Fahrrad bicycle} {Holz wood} {Wald wood}]

}
