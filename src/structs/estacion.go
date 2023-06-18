package structs

type Estacion struct {
	Nombre   string
	Duracion int
	Tramo    int
}

var Estaciones = [20]Estacion{
	{
		Nombre:   "Puerto",
		Duracion: 0,
		Tramo:    1,
	},
	{
		Nombre:   "Bellavista",
		Duracion: 1,
		Tramo:    1,
	},
	{
		Nombre:   "Francia",
		Duracion: 2,
		Tramo:    1,
	},
	{
		Nombre:   "Baron",
		Duracion: 5,
		Tramo:    1,
	},
	{
		Nombre:   "Portales",
		Duracion: 8,
		Tramo:    1,
	},
	{
		Nombre:   "Recreo",
		Duracion: 11,
		Tramo:    2,
	},
	{
		Nombre:   "Miramar",
		Duracion: 13,
		Tramo:    2,
	},
	{
		Nombre:   "Viña del Mar",
		Duracion: 14,
		Tramo:    2,
	},
	{
		Nombre:   "Hospital",
		Duracion: 16,
		Tramo:    2,
	},
	{
		Nombre:   "Chorrillos",
		Duracion: 17,
		Tramo:    2,
	},
	{
		Nombre:   "El Salto",
		Duracion: 19,
		Tramo:    2,
	},
	{
		Nombre:   "Quilpue",
		Duracion: 40,
		Tramo:    3,
	},
	{
		Nombre:   "El Sol",
		Duracion: 42,
		Tramo:    3,
	},
	{
		Nombre:   "El Belloto",
		Duracion: 45,
		Tramo:    3,
	},
	{
		Nombre:   "Las Americas",
		Duracion: 47,
		Tramo:    4,
	},
	{
		Nombre:   "La Concepcion",
		Duracion: 49,
		Tramo:    4,
	},
	{
		Nombre:   "Villa Alemana",
		Duracion: 51,
		Tramo:    4,
	},
	{
		Nombre:   "Sargento Aldea",
		Duracion: 53,
		Tramo:    4,
	},
	{
		Nombre:   "Peñablanca",
		Duracion: 56,
		Tramo:    4,
	},
	{
		Nombre:   "Limache",
		Duracion: 69,
		Tramo:    5,
	},
}
