package sql



func cargarDatos() {
	
	cargarClientes()
	cargarComercios()
	cargarTarjetas()
	
	
	if err != nil {
		log.Fatal(err)
	}
}


func cargarClientes() { //FALTA CARGAR TELEFONOS
	_ , err = db.Exec(` INSERT INTO cliente VALUES (81635249,'Juan','Perez','Corrientes 159',);
						INSERT INTO cliente VALUES (97824536,'Sergio','Denis','Av Lib San Martín 2130');
						INSERT INTO cliente VALUES (16495823,'Diego Armando','Maradona','Av Esteco 419');
						INSERT INTO cliente VALUES (87512694,'Alberto','Fernandez','Ocampo 248');
						INSERT INTO cliente VALUES (58214936,'Cristina','Kirchner','Av J R Vidal 1751');
						INSERT INTO cliente VALUES (87219364,'Juan Domingo','Perón','Av Cabildo 2554');
						INSERT INTO cliente VALUES (69254381,'Nestór','Kirchner','Rioja 951');
						INSERT INTO cliente VALUES (24318769,'Maximo','Kirchner','Primera Junta 889');
						INSERT INTO cliente VALUES (67918245,'Mariano','Pavone','Av A Palacios 1545');
						INSERT INTO cliente VALUES (15624837,'Juan Carlos','Olave','Almirón 301');
						INSERT INTO cliente VALUES (93527468,'Armando','Paredes','B Mitre 178');
						INSERT INTO cliente VALUES (82974315,'Juan Roman','Riquelme','E Santamarina 401');
						INSERT INTO cliente VALUES (48129563,'Martin','Palermo','Av Callao 892');
						INSERT INTO cliente VALUES (65317289,'Carlos','Tevez','Av San Martín 83');
						INSERT INTO cliente VALUES (49815267,'Angela Patricia','Gonzales','Alberti 6062');
						INSERT INTO cliente VALUES (51362874,'Nahuel','Rodriguez','Calle 24 1235');
						INSERT INTO cliente VALUES (83691452,'Federico','Garcia','Castro Barros 898');
						INSERT INTO cliente VALUES (64723591,'Matias Adrian','Perez','Río Cuarto 2020');
						INSERT INTO cliente VALUES (93167854,'Daniela','Gomez','Lavalle 2363');
						INSERT INTO cliente VALUES (84396721,'Martina','Botero','M De Alzaga 3972');`)
						
		if err != nil {
		log.Fatal(err)
		}
}
	
func cargarComercios(){
	 _ , err = db.Exec(`INSERT INTO comercio VALUES (95169,'Disco','Las Heras 716');
						INSERT INTO comercio VALUES (29981,'Carrefour','Av T A Edison 555');
						INSERT INTO comercio VALUES (82211,'Garbarino','Avenida 44 2049');
						INSERT INTO comercio VALUES (79701,'Coto','Aberastain Sur 163');
						INSERT INTO comercio VALUES (21724,'Falabella','Camargo 775');
						INSERT INTO comercio VALUES (51249,'Grido Helado','H Yrigoyen 1659');
						INSERT INTO comercio VALUES (87682,'Fravega','Tagle 3383');
						INSERT INTO comercio VALUES (59460,'Sodimac','Av Entre Ríos 1072');
						INSERT INTO comercio VALUES (34039,'Panaderia y confiteria','Av Colón 1012');
						INSERT INTO comercio VALUES (32694,'Verduleria el tomate','Calle 55 3067');
						INSERT INTO comercio VALUES (77289,'Fiambreria el gaucho','Acc Alte Brown 971');
						INSERT INTO comercio VALUES (68463,'Panaderia el pan','C Melo 4708');
						INSERT INTO comercio VALUES (75374,'Carniceria el carni','Brown 308');
						INSERT INTO comercio VALUES (79751,'Panaderia maxi','Av R S Ortiz 160');
						INSERT INTO comercio VALUES (38833,'Supermercado asia','Thompson 309');
						INSERT INTO comercio VALUES (68199,'Compumundo','P Molina 133');
						INSERT INTO comercio VALUES (63129,'Verduleria sandra','Bogotá 2842');
						INSERT INTO comercio VALUES (90949,'Solo deportes','Garibaldi 155');
						INSERT INTO comercio VALUES (68806,'Panaderia pan dulce','Chile 329');
						INSERT INTO comercio VALUES (61351,'Fiambreria luna','Av L Quaranta 7091');`)
		if err != nil {
		log.Fatal(err)
		}
	
}
	
func cargarTarjetas(){
	_, err = db.Exec(`INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();
						INSERT INTO tarjeta VALUES ();`)
		if err != nil {
		log.Fatal(err)
		}
}




