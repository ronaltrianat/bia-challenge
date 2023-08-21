# Descripcion de los Servicios

Este servicio permite consultar los consumos de energia de medidores mediante su id y un periodo de tiempo dado.

Los parametros para filtrar se envian mediante Query Parameters en una peticion http de tipo GET. A continuacion la descripcion de los parametros permitidos y su formato:

* **start_date:** Fecha desde donde se debe realizar la consulta. Formato **yyyy-MM-dd**
* **end_date:** Fecha hasta donde se debe realizar la consulta. Formato **yyyy-MM-dd**
* **kind_period:** Especifica la periodicidad de entrega de resultados. Valores permitidos: **monthly**, **weekly** y **daily**.
* **meters_ids:** Lista de id de los medidores. Tipo Integer. (Se envian varios query parameters si se requiere la consulta de varios medidores. Ejemplo: meters_ids=1&meters_ids=2)


**Method** : `GET`
**URL** : `https://$host/consumption`

## Ejemplos de consumo

**Solicitud de ejemplo consumo mensual para 2 medidores**

`http://$host/consumption?start_date=2023-06-01&end_date=2023-07-10&kind_period=monthly&meters_ids=1&meters_ids=2`

**HTTP Code** : `200 OK`
```json
{
    "data_graph": [
        {
            "meter_id": 1,
            "address": "some address mock for meter id 1",
            "active": [
                "1563573.517239999999",
                "4401580.218639999992"
            ],
            "reactive_inductive": [
                "166221.7478299999999",
                "475529.5926400000007"
            ],
            "reactive_capacitive": [
                "3.36599999999999997",
                "10.988000000000000073"
            ],
            "exported": [
                "2017.56832141194714",
                "5943.037877272413334"
            ],
            "period": [
                "JULY 2023",
                "JUNE 2023"
            ]
        },
        {
            "meter_id": 2,
            "address": "some address mock for meter id 2",
            "active": [
                "4008410.949719999992",
                "10898488.745770000001"
            ],
            "reactive_inductive": [
                "2443377.975379999998",
                "6605406.771729999995"
            ],
            "reactive_capacitive": [
                "0",
                "0"
            ],
            "exported": [
                "153.72857775836155",
                "468.7381462616538297"
            ],
            "period": [
                "JULY 2023",
                "JUNE 2023"
            ]
        }
    ]
}
```

**Solicitud de ejemplo consumo semanal para 2 medidores**

`http://$host/consumption?start_date=2023-06-01&end_date=2023-06-26&kind_period=weekly&meters_ids=1&meters_ids=2`

**HTTP Code** : `200 OK`
```json
{
    "data_graph": [
        {
            "meter_id": 1,
            "address": "some address mock for meter id 1",
            "active": [
                "559762.1560099999985",
                "1000319.3849099999995",
                "1023929.8746999999975",
                "1051010.9760799999955",
                "152389.4036"
            ],
            "reactive_inductive": [
                "61442.7725499999999",
                "109130.3795200000008",
                "110791.2690199999999",
                "112498.61108",
                "16235.31716"
            ],
            "reactive_capacitive": [
                "2.200000000000000084",
                "3.472000000000000018",
                "3.090999999999999934",
                "1.884000000000000021",
                "0.094000000000000016"
            ],
            "exported": [
                "778.583815582390492",
                "1371.924436419699618",
                "1384.641015249025038",
                "1401.516028353179608",
                "201.271199747180567"
            ],
            "period": [
                "2023-06-01 - 2023-06-04",
                "2023-06-05 - 2023-06-11",
                "2023-06-12 - 2023-06-18",
                "2023-06-19 - 2023-06-25",
                "2023-06-26 - 2023-06-26"
            ]
        },
        {
            "meter_id": 2,
            "address": "some address mock for meter id 2",
            "active": [
                "1355213.303340000001",
                "2441477.460609999995",
                "2547738.320480000004",
                "2622722.88544",
                "382485.651540000001"
            ],
            "reactive_inductive": [
                "818424.400980000001",
                "1476323.215460000006",
                "1543837.236059999992",
                "1591995.232299999993",
                "232537.26633"
            ],
            "reactive_capacitive": [
                "0",
                "0",
                "0",
                "0",
                "0"
            ],
            "exported": [
                "62.9650161415457537",
                "109.8321555802136019",
                "109.8960030656128952",
                "108.7723699849513862",
                "15.4760673520280662"
            ],
            "period": [
                "2023-06-01 - 2023-06-04",
                "2023-06-05 - 2023-06-11",
                "2023-06-12 - 2023-06-18",
                "2023-06-19 - 2023-06-25",
                "2023-06-26 - 2023-06-26"
            ]
        }
    ]
}
```

**Solicitud de ejemplo consumo diario para 2 medidores**

`http://$host/consumption?start_date=2023-06-01&end_date=2023-06-10&kind_period=daily&meters_ids=1&meters_ids=2`

**HTTP Code** : `200 OK`
```json
{
    "data_graph": [
        {
            "meter_id": 1,
            "address": "some address mock for meter id 1",
            "active": [
                "139088.88105",
                "139560.0447800000005",
                "140171.94727",
                "140941.282909999998",
                "141462.6982799999995",
                "141939.417770000001",
                "142388.264219999999",
                "142848.7589799999995",
                "143391.37723",
                "143925.17501"
            ],
            "reactive_inductive": [
                "15298.7634",
                "15336.4705099999999",
                "15384.31042",
                "15423.22822",
                "15452.3674000000007",
                "15512.1825100000001",
                "15550.22421",
                "15597.34263",
                "15624.38856",
                "15674.20633"
            ],
            "reactive_capacitive": [
                "0.014000000000000002",
                "1.127000000000000042",
                "0",
                "1.05900000000000004",
                "1.38500000000000004",
                "0.709",
                "0.292999999999999978",
                "0.133",
                "0.014",
                "0.195"
            ],
            "exported": [
                "194.196302689726036",
                "194.397164944953429",
                "194.672437023284353",
                "195.317910924426674",
                "195.71451028729238",
                "195.604529034759507",
                "195.760240324925666",
                "195.804729044676751",
                "196.257761386062451",
                "196.37543835885101"
            ],
            "period": [
                "Jun 01 2023",
                "Jun 02 2023",
                "Jun 03 2023",
                "Jun 04 2023",
                "Jun 05 2023",
                "Jun 06 2023",
                "Jun 07 2023",
                "Jun 08 2023",
                "Jun 09 2023",
                "Jun 10 2023"
            ]
        },
        {
            "meter_id": 2,
            "address": "some address mock for meter id 2",
            "active": [
                "335944.0161",
                "337849.61211",
                "339756.28815",
                "341663.386980000001",
                "343419.533259999998",
                "345075.887249999998",
                "346828.5856",
                "348678.86968",
                "350549.98928",
                "352485.85058"
            ],
            "reactive_inductive": [
                "202785.8657",
                "203975.68317",
                "205207.156610000001",
                "206455.6955",
                "207590.818510000001",
                "208584.288380000001",
                "209636.620340000002",
                "210797.02562",
                "211983.56153",
                "213232.050000000001"
            ],
            "reactive_capacitive": [
                "0",
                "0",
                "0",
                "0",
                "0",
                "0",
                "0",
                "0",
                "0",
                "0"
            ],
            "exported": [
                "15.7594637415215435",
                "15.7517568853752587",
                "15.7362005015771862",
                "15.7175950130717653",
                "15.7034376612131079",
                "15.7049137155118815",
                "15.7062637307156052",
                "15.6983474316334477",
                "15.6879923225122906",
                "15.6734988098048871"
            ],
            "period": [
                "Jun 01 2023",
                "Jun 02 2023",
                "Jun 03 2023",
                "Jun 04 2023",
                "Jun 05 2023",
                "Jun 06 2023",
                "Jun 07 2023",
                "Jun 08 2023",
                "Jun 09 2023",
                "Jun 10 2023"
            ]
        }
    ]
}
```

## Respuesta a solicitudes invalidas

`http://$host/consumption?start_date=2023-06-017&end_date=2023-06-107&kind_period=daily7&meters_ids=1s`

**HTTP Code** : `400 Bad Request`
```json
{
	"status": "Invalid request.",
	"error": "kind_period is invalid : invalid start date format : invalid end date format : meters_ids are invalid"
}
```
