package sql



func cargarDatos() {
	
	cargarClientes()
	cargarComercios()
	cargarTarjetas()
	
	
	if err != nil {
		log.Fatal(err)
	}
}


func cargarClientes() {
	_ , err = db.Exec(` INSERT INTO cliente VALUES (81635249);
						INSERT INTO cliente VALUES (97824536);
						INSERT INTO cliente VALUES (16495823);
						INSERT INTO cliente VALUES (87512694);
						INSERT INTO cliente VALUES (58214936);
						INSERT INTO cliente VALUES (87219364);
						INSERT INTO cliente VALUES (69254381);
						INSERT INTO cliente VALUES (24318769);
						INSERT INTO cliente VALUES (67918245);
						INSERT INTO cliente VALUES (15624837);
						INSERT INTO cliente VALUES (93527468);
						INSERT INTO cliente VALUES (82974315);
						INSERT INTO cliente VALUES (48129563);
						INSERT INTO cliente VALUES (65317289);
						INSERT INTO cliente VALUES (49815267);
						INSERT INTO cliente VALUES (51362874);
						INSERT INTO cliente VALUES (83691452);
						INSERT INTO cliente VALUES (64723591);
						INSERT INTO cliente VALUES (93167854);
						INSERT INTO cliente VALUES (84396721);`)
						
		if err != nil {
		log.Fatal(err)
		}
}
	
func cargarComercios(){
	 _ , err = db.Exec(`INSERT INTO comercio VALUES (95169);
						INSERT INTO comercio VALUES (29981);
						INSERT INTO comercio VALUES (82211);
						INSERT INTO comercio VALUES (79701);
						INSERT INTO comercio VALUES (21724);
						INSERT INTO comercio VALUES (51249);
						INSERT INTO comercio VALUES (87682);
						INSERT INTO comercio VALUES (59460);
						INSERT INTO comercio VALUES (34039);
						INSERT INTO comercio VALUES (32694);
						INSERT INTO comercio VALUES (77289);
						INSERT INTO comercio VALUES (68463);
						INSERT INTO comercio VALUES (75374);
						INSERT INTO comercio VALUES (79751);
						INSERT INTO comercio VALUES (38833);
						INSERT INTO comercio VALUES (68199);
						INSERT INTO comercio VALUES (63129);
						INSERT INTO comercio VALUES (90949);
						INSERT INTO comercio VALUES (68806);
						INSERT INTO comercio VALUES (61351);`)
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




