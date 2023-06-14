package src

type card struct {
	name     string
	discount int
}

func newCard(name string, discount int) card {
	return card{name, discount}
}
