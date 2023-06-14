package src

type station struct {
	name    string
	minutes int
	sector  int
}

func newStation(name string, minutes int, sector int) station {
	return station{name, minutes, sector}
}

var Stations = [20]station{
	station{
		name:    "Puerto",
		minutes: 0,
		sector:  1,
	},
	station{
		name:    "Bellavista",
		minutes: 1,
		sector:  1,
	},
	station{
		name:    "Francia",
		minutes: 2,
		sector:  1,
	},
	station{
		name:    "Baron",
		minutes: 5,
		sector:  1,
	},
	station{
		name:    "Portales",
		minutes: 8,
		sector:  1,
	},
	station{
		name:    "Recreo",
		minutes: 11,
		sector:  2,
	},
	station{
		name:    "Miramar",
		minutes: 13,
		sector:  2,
	},
	station{
		name:    "Viña del Mar",
		minutes: 14,
		sector:  2,
	},
	station{
		name:    "Hospital",
		minutes: 16,
		sector:  2,
	},
	station{
		name:    "Chorrillos",
		minutes: 17,
		sector:  2,
	},
	station{
		name:    "El Salto",
		minutes: 19,
		sector:  2,
	},
	station{
		name:    "Quilpue",
		minutes: 40,
		sector:  3,
	},
	station{
		name:    "El Sol",
		minutes: 42,
		sector:  3,
	},
	station{
		name:    "El Belloto",
		minutes: 45,
		sector:  3,
	},
	station{
		name:    "Las Americas",
		minutes: 47,
		sector:  4,
	},
	station{
		name:    "La Concepcion",
		minutes: 49,
		sector:  4,
	},
	station{
		name:    "Villa Alemana",
		minutes: 51,
		sector:  4,
	},
	station{
		name:    "Sargento Aldea",
		minutes: 53,
		sector:  4,
	},
	station{
		name:    "Peñablanca",
		minutes: 56,
		sector:  4,
	},
	station{
		name:    "Limache",
		minutes: 69,
		sector:  5,
	},
}
