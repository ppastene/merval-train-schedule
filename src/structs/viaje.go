package structs

import (
	"src/src/helpers"
	"time"
)

type Viaje struct {
	Origen  Estacion
	Destino Estacion
}

func NewViaje(o Estacion, d Estacion) Viaje {
	return Viaje{o, d}
}

func (v Viaje) ObtenerTiempoViaje() int {
	return helpers.Abs(v.Origen.Duracion - v.Destino.Duracion)
}

func (v Viaje) ObtenerDireccionViaje() int {
	if v.Origen.Duracion < v.Destino.Duracion {
		return 0
	}
	return 19
}

func (v Viaje) ObtenerItinerario(t time.Time, d int) Itinerario {
	return ObtenerItinerario(t, d)
}
