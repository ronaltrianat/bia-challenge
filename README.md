# bia-challenge
Prueba tecnica para Bia.

## Descripcion de los Servicios

Este servicio permite consultar los consumos de energia de medidores mediante su id y un periodo de tiempo dado.

Los parametros para filtrar se envian mediante Query Parameters en una peticion http de tipo GET. A continuacion la descripcion de los parametros permitidos y su formato:

* **start_date:** Fecha desde donde se debe realizar la consulta. Formato **yyyy-MM-dd**
* **end_date:** Fecha hasta donde se debe realizar la consulta. Formato **yyyy-MM-dd**
* **kind_period:** Especifica la periodicidad de entrega de resultados. Valores permitidos: **monthly**, **weekly** y **daily**.
* **meters_ids:** Lista de id de los medidores. Tipo Integer. (Se envian varios query parameters si se requiere la consulta de varios medidores. Ejemplo: meters_ids=1&meters_ids=2)

[Mas Detalle](docs/service_description.md)

## Descripcion del componente

Se implementa una arquitectura hexagonal en la cual se busca encapsular la logica de negocio.

El microservicio esta diseñado para que se pueda crear una implementacion nueva de los puertos de entrada y de salida mediante nuevos adaptadores sin necesidad de afectar la logica de negocio ya construida.

![Diagrama 3D](/docs/bia-diagram.jpg)


## Postman con ejemplos de consumo

Con el siguiente archivo postman, puede realizar test de los servicios desplegado en AWS.

* [Export Postman Test Cloud](docs/bia-challenge.postman_collection.json)
* [Export Postman Test Local](docs/bia-challenge.postman_collection.json)