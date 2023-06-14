package src

type fee struct {
	name  string
	value [5]float64
}

func newFee(name string, value [5]float64) fee {
	return fee{name, value}
}

var horaBaja = fee{
	"Hora Baja",
	[5]float64{388.0, 534.0, 769.0, 815.0, 1012.0},
}

var horaMedia = fee{
	"Hora Media",
	[5]float64{409.0, 566.0, 811.0, 860.0, 1069.0},
}

var horaAlta = fee{
	"Hora Alta",
	[5]float64{430.0, 590.0, 854.0, 906.0, 1126.0},
}
