package main

import (
	"fmt"
	"src/src/helpers"
	"src/src/structs"
	"time"
)

func obtenerInputs() (int, int, int) {
	var origen, destino, tarjeta int
	fmt.Println("Tramo 1: 1: Puerto - 2: Bellavista - 3: Francia - 3: Baron - 4: Portales")
	fmt.Println("Tramo 2: 5: Recreo - 6 Miramar - 7: Viña del Mar - 8: Hospital - 9: Chorrillos - 10: El Salto")
	fmt.Println("Tramo 3: 11: Quilpue - 12: El Sol - 13: El Belloto")
	fmt.Println("Tramo 4: 14: Las Americas - 15: La Concepcion - 16: Villa Alemana - 17: Sargento Aldea - 18: Peñablanca")
	fmt.Println("Tramo 5: 19: Limache")
	fmt.Println("Tipos de usuario: 0: General (sin %) - 1: Estudiante (66%) - 2: Tercera Edad (50%)")
	fmt.Println("Porfavor ingrese los datos separados por un espacio y en el siguiente orden: ")
	fmt.Println("  Estacion Origen  Estacion Destino  Tipo de Usuario  ")
	_, err := fmt.Scanf("%v %v %v", &origen, &destino, &tarjeta)
	if err != nil {
		fmt.Printf("Your input could not be read%v", err)
	}

	return origen, destino, tarjeta
}

func obtenerInfoBase(origen int, destino int, hora time.Time) (structs.Viaje, structs.Itinerario, int, time.Duration) {
	var viaje = structs.NewViaje(structs.Estaciones[origen], structs.Estaciones[destino])
	var itinerario = viaje.ObtenerItinerario(hora, viaje.ObtenerDireccionViaje())
	var tramo = helpers.Abs(viaje.Origen.Tramo - viaje.Destino.Tramo)
	var tiempoViaje = time.Duration(helpers.Abs(viaje.Origen.Duracion-viaje.Destino.Duracion)) * time.Minute

	return viaje, itinerario, tramo, tiempoViaje
}

func obtenerHoraSalida(itinerario structs.Itinerario, frecuencia time.Duration, fecha time.Time) (time.Time, time.Time) {
	var salida = fecha
	var llegada time.Time
	if fecha.Before(itinerario.PrimerTren) {
		salida = itinerario.PrimerTren
	}
	if salida.After(itinerario.UltimoTren) {
		salida = itinerario.UltimoTren
	}
	llegada = salida.Add(frecuencia)
	return salida, llegada
}

func obtenerValorViaje(tarifa structs.Tarifa, tramo int, usuario int) int {
	var valor int = tarifa.Valores[tramo]
	var tarjeta structs.TarjetaEspecial
	switch usuario {
	case 0:
		return valor
	case 1:
		tarjeta = *structs.Estudiante
	case 2:
		tarjeta = *structs.TerceraEdad
	}

	valor = tarjeta.AplicarDescuento(valor)
	return valor
}

func dibujarTabla(viaje structs.Viaje, itinerario structs.Itinerario, tiempoViaje time.Duration, frecuencia time.Duration, tramo int, tarjeta int) {
	origen := viaje.Origen
	destino := viaje.Destino
	var primerTren, ultimoTren time.Time = itinerario.PrimerTren, itinerario.UltimoTren
	var horaSalida time.Time = primerTren
	var horaLlegada time.Time
	var tarifa structs.Tarifa
	var valorViaje int
	fmt.Printf("Origen: %v Destino: %v\n", origen.Nombre, destino.Nombre)
	for itinerario.UltimoTren.After(horaSalida) {
		horaLlegada = horaSalida.Add(tiempoViaje)
		tarifa = structs.ObtenerTarifaSegunFecha(horaLlegada)
		valorViaje = obtenerValorViaje(tarifa, tramo, tarjeta)
		fmt.Printf("Salida: %v Llegada: %v Tarifa: %v Valor: %v\n", horaSalida.Format(time.Kitchen), horaLlegada.Format(time.Kitchen), tarifa.Nombre, valorViaje)
		horaSalida = horaSalida.Add(frecuencia)
	}
	horaSalida = ultimoTren
	fmt.Printf("Salida: %v Llegada: %v Hora: %v Valor: %v\n", horaSalida.Format(time.Kitchen), horaSalida.Add(tiempoViaje).Format(time.Kitchen), tarifa.Nombre, valorViaje)
}

func main() {
	fecha := time.Now()
	frecuencia := time.Minute * 12

	var origen, destino, tarjeta = obtenerInputs()
	var viaje, itinerario, tramo, tiempoViaje = obtenerInfoBase(origen, destino, fecha)
	dibujarTabla(viaje, itinerario, tiempoViaje, frecuencia, tramo, tarjeta)
}
