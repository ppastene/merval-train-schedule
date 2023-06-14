package src

import "errors"

type Station struct {
	Name    string
	Minutes int
	Sector  int
}

var Stations = [20]Station{
	{
		Name:    "Puerto",
		Minutes: 0,
		Sector:  1,
	},
	{
		Name:    "Bellavista",
		Minutes: 1,
		Sector:  1,
	},
	{
		Name:    "Francia",
		Minutes: 2,
		Sector:  1,
	},
	{
		Name:    "Baron",
		Minutes: 5,
		Sector:  1,
	},
	{
		Name:    "Portales",
		Minutes: 8,
		Sector:  1,
	},
	{
		Name:    "Recreo",
		Minutes: 11,
		Sector:  2,
	},
	{
		Name:    "Miramar",
		Minutes: 13,
		Sector:  2,
	},
	{
		Name:    "Viña del Mar",
		Minutes: 14,
		Sector:  2,
	},
	{
		Name:    "Hospital",
		Minutes: 16,
		Sector:  2,
	},
	{
		Name:    "Chorrillos",
		Minutes: 17,
		Sector:  2,
	},
	{
		Name:    "El Salto",
		Minutes: 19,
		Sector:  2,
	},
	{
		Name:    "Quilpue",
		Minutes: 40,
		Sector:  3,
	},
	{
		Name:    "El Sol",
		Minutes: 42,
		Sector:  3,
	},
	{
		Name:    "El Belloto",
		Minutes: 45,
		Sector:  3,
	},
	{
		Name:    "Las Americas",
		Minutes: 47,
		Sector:  4,
	},
	{
		Name:    "La Concepcion",
		Minutes: 49,
		Sector:  4,
	},
	{
		Name:    "Villa Alemana",
		Minutes: 51,
		Sector:  4,
	},
	{
		Name:    "Sargento Aldea",
		Minutes: 53,
		Sector:  4,
	},
	{
		Name:    "Peñablanca",
		Minutes: 56,
		Sector:  4,
	},
	{
		Name:    "Limache",
		Minutes: 69,
		Sector:  5,
	},
}

func NewStation(Name string, Minutes int, Sector int) Station {
	return Station{Name, Minutes, Sector}
}

func GetStation(index int) (Station, error) {
	if index < 0 || index > len(Stations) {
		return Station{}, errors.New("The station doesn't exist")
	}
	return Stations[index], nil
}

func (origin Station) GetTravelTime(destination Station) int {
	return Abs(origin.Minutes - destination.Minutes)
}

func (origin Station) GetTravelDirection(destination Station) int {
	if origin.Minutes < destination.Minutes {
		return 0 // Tren sale de Estacion Puerto
	}
	return 19 // Tren sale de Estacion Limache
}
