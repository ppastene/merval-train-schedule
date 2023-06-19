package structs

import (
	"src/src/helpers"
	"time"
)

type Tarifa struct {
	Id      int
	Nombre  string
	Valores []int
}

func NewTarifa(id int, n string, v []int) Tarifa {
	return Tarifa{id, n, v}
}

var TarifaBaja = NewTarifa(0, "Hora Baja", []int{388.0, 534.0, 769.0, 815.0, 1012.0})
var TarifaMedia = NewTarifa(1, "Hora Media", []int{409.0, 566.0, 811.0, 860.0, 1069.0})
var TarifaAlta = NewTarifa(2, "Hora Alta", []int{430.0, 590.0, 854.0, 906.0, 1126.0})

func ObtenerTarifaSegunFecha(t time.Time) Tarifa {
	var tarifa Tarifa
	switch t.Weekday() {
	case 6, 0:
		tarifa = TarifaBaja
	case 1, 2, 3, 4, 5:
		if helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 10, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 12, 59, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 14, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 15, 59, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())) {
			tarifa = TarifaBaja
		} else if helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 6, 29, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 10, 29, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 13, 00, 00, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 13, 59, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 16, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 16, 59, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 20, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 20, 59, 59, 0, t.Location())) {
			tarifa = TarifaMedia
		} else if helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 9, 29, 59, 0, t.Location())) ||
			helpers.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 17, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 19, 59, 59, 0, t.Location())) {
			tarifa = TarifaAlta
		}
	}
	return tarifa
}
