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

func dibujarTabla(viaje structs.Viaje, itinerario structs.Itinerario, tiempoViaje time.Duration, tramo int, tarjeta int) {
	var origen, destino structs.Estacion = viaje.Origen, viaje.Destino
	var direccion int = viaje.ObtenerDireccionViaje()
	var frecuencia time.Duration = time.Duration(time.Minute * 12)
	var salidaTerminal, salidaProximoTren time.Time = itinerario.PrimerTren, itinerario.PrimerTren
	var ultimoTren time.Time = itinerario.UltimoTren
	var salidaOrigen, llegadaDestino time.Time
	var tarifa structs.Tarifa
	var valorViaje int
	fmt.Printf("Itinerario de trenes MERVAL\n")
	fmt.Printf("Origen: %v\nDestino: %v\nTiempo de Viaje: %v\n", origen.Nombre, destino.Nombre, tiempoViaje)
	for itinerario.UltimoTren.After(salidaTerminal) {
		if salidaTerminal.Equal(salidaProximoTren) {
			salidaOrigen = salidaTerminal.Add(time.Duration(helpers.Abs(viaje.Origen.Duracion-structs.Estaciones[direccion].Duracion)) * time.Minute)
			llegadaDestino = salidaOrigen.Add(tiempoViaje)
			tarifa = structs.ObtenerTarifaSegunFecha(llegadaDestino)
			valorViaje = obtenerValorViaje(tarifa, tramo, tarjeta)
			fmt.Printf("| Salida: %-7v | Llegada: %-7v | Tarifa: %-10v | Valor: %-3v |\n", salidaOrigen.Format(time.Kitchen), llegadaDestino.Format(time.Kitchen), tarifa.Nombre, valorViaje)
			salidaTerminal = salidaTerminal.Add(frecuencia)
		}
		salidaProximoTren = salidaProximoTren.Add(time.Minute * 1)
	}
	salidaOrigen = ultimoTren.Add(time.Duration(helpers.Abs(viaje.Origen.Duracion-structs.Estaciones[direccion].Duracion)) * time.Minute)
	fmt.Printf("| Salida: %-7v | Llegada: %-7v | Tarifa: %-10v | Valor: %-3v |\n", salidaOrigen.Format(time.Kitchen), salidaOrigen.Add(tiempoViaje).Format(time.Kitchen), tarifa.Nombre, valorViaje)
}

func main() {
	t := time.Now()
	fecha := time.Date(2023, 06, 19, 8, 0, 0, 0, t.Location())
	var origen, destino, tarjeta = obtenerInputs()
	var viaje, itinerario, tramo, tiempoViaje = obtenerInfoBase(origen, destino, fecha)
	dibujarTabla(viaje, itinerario, tiempoViaje, tramo, tarjeta)
}
