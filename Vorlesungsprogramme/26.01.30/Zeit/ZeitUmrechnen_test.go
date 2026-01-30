package Zeit

type Zeit int

func (l Zeit) Sekunden() int {
	return int(l)
}

func (l Zeit) Minuten() int {
	return l.Sekunden() / 60
}

func (l Zeit) Stunden() int {
	return l.Minuten() / 60
}

func Example() {
	var a Zeit = 216000

	println(a.Sekunden())
	println(a.Minuten())
	println(a.Stunden())

	//Output:
	// 216000
	// 3600
	// 60

}
