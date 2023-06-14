package src

type card struct {
	Name     string
	Discount int
}

var Cards = [3]card{
	{"General", 100},
	{"Estudiante", 66},
	{"Tercera Edad", 50},
}

func NewCard(name string, discount int) card {
	return card{name, discount}
}
