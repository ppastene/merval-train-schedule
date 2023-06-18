package structs

import (
	"src/src/helpers"
	"time"
)

type Itinerario struct {
	PrimerTren time.Time
	UltimoTren time.Time
}

func NewItinerario(primer time.Time, ultimo time.Time) Itinerario {
	return Itinerario{primer, ultimo}
}

func ObtenerItinerario(t time.Time, direction int) Itinerario {
	var itinerario Itinerario
	dia := t.Weekday()
	if int(dia) > 6 {
		dia = helpers.DayOverflow(dia)
	}
	if direction == 0 { // Si tren sale de Puerto
		switch dia {
		case 0:
			itinerario = Itinerario{
				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 12, 0, 0, t.Location()),
			}
		case 6:
			itinerario = Itinerario{

				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 8, 30, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 24, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:
			itinerario = Itinerario{

				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 30, 0, 0, t.Location()),
			}
		}
	} else if direction == 19 { // Si tren sale de Limache
		switch dia {
		case 0:
			itinerario = Itinerario{
				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 6, 0, 0, t.Location()),
			}
		case 6:
			itinerario = Itinerario{
				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 7, 30, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 26, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:

			itinerario = Itinerario{
				PrimerTren: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				UltimoTren: time.Date(t.Year(), t.Month(), t.Day(), 22, 15, 0, 0, t.Location()),
			}
		}
	}
	return itinerario
}
