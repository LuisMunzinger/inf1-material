package dict

import (
	"Dictionary/entry"
	"fmt"
)

func ExampleNew() {
	d := New()

	fmt.Println(d)

	// Output:
	// {[]}
}

func ExampleDict_Add() {
	d := New()
	d.Add(entry.New("Apfel", "apple"))
	d.Add(entry.New("Birne", "pear"))

	fmt.Println(d)

	// Output:
	// {[{Apfel apple} {Birne pear}]}
}

func ExampleDict_Size() {
	d := New()
	fmt.Println(d.Size())

	d.entries = append(d.entries, entry.New("Apfel", "apple"))
	fmt.Println(d.Size())

	// Output:
	// 0
	// 1
}

func ExampleDict_GetDe() {
	d := New()
	d.Add(entry.New("Apfel", "apple"))
	d.Add(entry.New("Birne", "pear"))

	fmt.Println(d.GetDe("Apfel"))
	fmt.Println(d.GetDe("Birne"))
	fmt.Println(d.GetDe("Banane"))

	// Output:
	// {Apfel apple}
	// {Birne pear}
	// { }
}

func ExampleDict_Lookup() {
	d := New()
	d.Add(entry.New("Apfel", "apple"))
	d.Add(entry.New("Birne", "pear"))

	fmt.Println(d.Lookup("Apfel"))
	fmt.Println(d.Lookup("Birne"))
	fmt.Println(d.Lookup("Banane"))

	// Output:
	// apple
	// pear
	//
}

func ExampleDict_GetAllDe() {
	d := New()
	d.Add(entry.New("Wald", "wood"))
	d.Add(entry.New("Wald", "forest"))
	d.Add(entry.New("Haus", "house"))
	d.Add(entry.New("Haus", "mansion"))
	d.Add(entry.New("Erdbeere", "strawberry"))

	fmt.Println(d.GetAllDe("Wald"))
	fmt.Println(d.GetAllDe("Haus"))
	fmt.Println(d.GetAllDe("Erdbeere"))
	fmt.Println(d.GetAllDe("Banane"))

	// Output:
	// [{Wald wood} {Wald forest}]
	// [{Haus house} {Haus mansion}]
	// [{Erdbeere strawberry}]
	// []
}

func ExampleDict_LookupAll() {
	d := New()
	d.Add(entry.New("Wald", "wood"))
	d.Add(entry.New("Wald", "forest"))
	d.Add(entry.New("Haus", "house"))
	d.Add(entry.New("Haus", "mansion"))
	d.Add(entry.New("Erdbeere", "strawberry"))

	fmt.Println(d.LookupAll("Wald"))
	fmt.Println(d.LookupAll("Haus"))
	fmt.Println(d.LookupAll("Erdbeere"))
	fmt.Println(d.LookupAll("Banane"))

	// Output:
	// [wood forest]
	// [house mansion]
	// [strawberry]
	// []
}
