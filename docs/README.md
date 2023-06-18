# Explicacion del problema

## Introduccion
El Metro de Valparaiso (de ahora en adelante, MERVAL) consiste de una linea de tren con 20 estaciones que va desde Plaza Altamirano en Valparaiso hasta el centro de la ciudad de Limache separados por casi 42 kilometros de linea ferroviaria, ademas de un servicio de acercamiento de buses desde Estacion Limache para los sectores de Olmue, Limache Viejo y Quillota.

## Problema
Si bien hay información con respecto al servicio de MERVAL en realidad no existe un itinerario de trenes que muestre las salidas de los trenes dada una estación de origen y destino, el tiempo de viaje y el valor de la tarifa, solo existe un pequeño formulario que solo muestra los precios segun las estaciones y la hora de servicio. Esto genera un problema para un usuario que desea información mas detallada sobre los el itinerario del día del servicio de MERVAL considerando que la frecuencia de trenes es muy baja un usuario puede esperar hasta 15 minutos por la llegada de un tren, una frencuencia muy del servicio del Metro de Santiago

## Solucion
Esta aplicación en base a los datos dados por el usuario como la estacion de origen y destino y el tipo de tarjeta genera una tabla con todas las salidas de trenes de una estacion, el tiempo de viaje, hora de llegada y valor de la tarifa. Dicha información es generada en base a la existente dada por MERVAL y permite mostrar un itinerario de trenes mas completo. En base a la siguiente información:

- Apertura y cierre de estaciones
- Hora de salida del primer y ultimo tren
- Rango de horas para que el servicio sea en hora baja, media o alta
- Tiempo de viaje entre estaciones
- Descuentos y beneficios para los usuarios 
- Frecuencia de salida de trenes segun la hora del servicio
- Valores del viaje en base a la hora de servicio y los tramos de estaciones

Se puede calcular lo siguiente para el usuario
- Hora de salida y hora de llegada en base a la distancia de las estaciones y la frencuencia de salida de los trenes
- Hora de servicio y el valor de la tarifa junto con el beneficio de usuario si aplica

Todo en base al día en que se ocupa la aplicación ya que la frecuencia y los valores cambian segun si el servicio es en dia de semana o si el servicio ocurre durante hora punta

## Diagrama de Clases
-- Insertar Diagrama Clases --<br>
Un problema con este proyecto es que si bien hay clases que no tienen una relación directa con otras clases igual dependen de estas para obtener la información que se debe mostrar al usuario. Si bien un sistema como este requiere de muchas mas clases como "Historial" para ver todo el registro de viajes del usuario se hizo lo posible para que esto quedara lo mas simple posible. Además, toda la informacion para generar al usuario esta incrustada en el codigo asi que el usuario no tiene que crear absolutamente nada ya que solo debe ingresar unos valores para obtener la data requerida

## Caso de Uso
-- Insertar Diagrama Flujo --<br>
Un usuario quiere viajar un dia viernes desde Estacion Quilpue a Estacion Puerto a las 10:00am. No sabe las horas de salida de los trenes, la distancia que puede tomar el viaje y el valor de la tarifa. Decide usar esta aplicacion para generar el itinerario en base a dicho día, como la aplicacion muestra el itinerario en base al dia actual el usuario tiene que usarla el mismo dia para generar la informacion del dia.

Al abrir la aplicacion el usuario debe ingresar tres valores enteros: Los ID de las estaciones de origen y destino ademas del ID del tipo de tarjeta por si tiene una con beneficio Quilpue tiene ID 11, Puerto tiene ID 0 y la tarjeta general tiene ID 0, por tanto el usuario debe ingresar en pantalla: 11 0 0

El sistema busca y retorna dos objetos de tipo Estacion y los asigna a un objeto de tipo Viaje. Se ejecuta los metodos ++++++ para obtener la direccion del viaje (si es hacia Puerto o Limache) y +++++++ para obtener la cantidad de tramos entre ambas estaciones. Se consulta la funcion +++++++++ del cual en base a la direccion y la fecha actual se retorna un objeto Itinerario con las horas de salida del primer y ultimo tren. Con la hora del primer viaje se setea un contador de tiempo del cual se le va sumando la frecuencia de salida de los trenes (por el momento esta seteado a 12 minutos) y por cada actualización se consulta la funcion ++++++++ que retorna un objeto de tipo Tarifa del cual contiene el nombre de la hora de servicio y un conjunto de tarifas. Se obtiene la tarifa exacta en base al calculo de tramos de las estaciones de Viaje, y antes de retornar el valor al usuario se retorna un objeto ya sea de tipo Tarjeta o TarjetaEspecial dependiendo del valor de tarjeta que ingreso el usuario, si es TarjetaEspecial se verificará si la tarjeta aun esta vigente y en caso de estarlo se aplica un descuento con la funcion ++++++++. Finalmente se muestra en pantalla la siguiente informacion:

Salida: Hora de salida del primer tren en base a Itinerario.PrimerTren
Llegada: Hora de llegada del tren en base a la hora de salida + el tiempo entre las estaciones
Tarifa: El tipo de hora de servicio del cual se basa en la hora de llegada del tren. 
Valor: El valor de la tarifa en base a los tramos de las estaciones y la hora de servicio

Para mostrar la siguiente hora se repite todo el proceso pero esta vez en la hora de salida se le suma la frecuencia que indica la hora de salida del proximo tren. Todo esto se repite en un bucle  hasta que la hora de salida sea mayor a la hora de Itinerario.UltimoTren en donde se muestra una ultima hora de salida en base a dicha hora.