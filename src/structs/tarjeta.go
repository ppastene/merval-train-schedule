package structs

type Tarjeta struct {
	Nombre string
}

func NewTarjeta(n string) Tarjeta {
	return Tarjeta{n}
}

var General = NewTarjeta("General")
