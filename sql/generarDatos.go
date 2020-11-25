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
	_ , err = db.Exec(` INSERT INTO cliente VALUES (81635249,'Juan','Perez','Corrientes 159','271542570123');
						INSERT INTO cliente VALUES (97824536,'Sergio','Denis','Av Lib San Martín 2130','307692152314');
						INSERT INTO cliente VALUES (16495823,'Diego Armando','Maradona','Av Esteco 419','997410050001');
						INSERT INTO cliente VALUES (87512694,'Alberto','Fernandez','Ocampo 248','279062550540');
						INSERT INTO cliente VALUES (58214936,'Cristina','Kirchner','Av J R Vidal 1751','565034300512');
						INSERT INTO cliente VALUES (87219364,'Juan Domingo','Perón','Av Cabildo 2554','321914144530');
						INSERT INTO cliente VALUES (69254381,'Nestór','Kirchner','Rioja 951','350375190034');
						INSERT INTO cliente VALUES (24318769,'Maximo','Kirchner','Primera Junta 889','567869941940');
						INSERT INTO cliente VALUES (67918245,'Mariano','Pavone','Av A Palacios 1545','284033656589');
						INSERT INTO cliente VALUES (15624837,'Juan Carlos','Olave','Almirón 301','382324807564');
						INSERT INTO cliente VALUES (93527468,'Armando','Paredes','B Mitre 178','683146119876');
						INSERT INTO cliente VALUES (82974315,'Juan Roman','Riquelme','E Santamarina 401','763110342400');
						INSERT INTO cliente VALUES (48129563,'Martin','Palermo','Av Callao 892','928847559909');
						INSERT INTO cliente VALUES (65317289,'Carlos','Tevez','Av San Martín 83','446819089687');
						INSERT INTO cliente VALUES (49815267,'Angela Patricia','Gonzales','Alberti 6062','931248588055');
						INSERT INTO cliente VALUES (51362874,'Nahuel','Rodriguez','Calle 24 1235','284529192033');
						INSERT INTO cliente VALUES (83691452,'Federico','Garcia','Castro Barros 898','894803786745');
						INSERT INTO cliente VALUES (64723591,'Matias Adrian','Perez','Río Cuarto 2020','221651370676');
						INSERT INTO cliente VALUES (93167854,'Daniela','Gomez','Lavalle 2363','180274635956');
						INSERT INTO cliente VALUES (84396721,'Martina','Botero','M De Alzaga 3972','654507732034');`)
						
		if err != nil {
		log.Fatal(err)
		}
}
	
func cargarComercios(){
	 _ , err = db.Exec(`INSERT INTO comercio VALUES (95169,'Disco','Las Heras 716','29812595','872178944032');
						INSERT INTO comercio VALUES (29981,'Carrefour','Av T A Edison 555','37466130','457755059512');
						INSERT INTO comercio VALUES (82211,'Garbarino','Avenida 44 2049','39275046','625050693943');
						INSERT INTO comercio VALUES (79701,'Coto','Aberastain Sur 163','60991299','709777835612');
						INSERT INTO comercio VALUES (21724,'Falabella','Camargo 775','57347558','138988075243');
						INSERT INTO comercio VALUES (51249,'Grido Helado','H Yrigoyen 1659','30179225','361518160053');
						INSERT INTO comercio VALUES (87682,'Fravega','Tagle 3383','68627468','925173628765');
						INSERT INTO comercio VALUES (59460,'Sodimac','Av Entre Ríos 1072','87860378','915082365243');
						INSERT INTO comercio VALUES (34039,'Panaderia y confiteria','Av Colón 1012','43369535','634483423800');
						INSERT INTO comercio VALUES (32694,'Verduleria el tomate','Calle 55 3067','49628779','713519136410');
						INSERT INTO comercio VALUES (77289,'Fiambreria el gaucho','Acc Alte Brown 971','17473237','726870604311');
						INSERT INTO comercio VALUES (68463,'Panaderia el pan','C Melo 4708','48005807','988294290165');
						INSERT INTO comercio VALUES (75374,'Carniceria el carni','Brown 308','51589936','284719172711');
						INSERT INTO comercio VALUES (79751,'Panaderia maxi','Av R S Ortiz 160','83138226','267532588765');
						INSERT INTO comercio VALUES (38833,'Supermercado asia','Thompson 309','32787317','777139421953');
						INSERT INTO comercio VALUES (68199,'Compumundo','P Molina 133','74727863','6708988871');
						INSERT INTO comercio VALUES (63129,'Verduleria sandra','Bogotá 2842','39276629','292802798211');
						INSERT INTO comercio VALUES (90949,'Solo deportes','Garibaldi 155','36548082','333826274011');
						INSERT INTO comercio VALUES (68806,'Panaderia pan dulce','Chile 329','87303986','946007831954');
						INSERT INTO comercio VALUES (61351,'Fiambreria luna','Av L Quaranta 7091','46863583','339534886414');`)
		if err != nil {
		log.Fatal(err)
		}
	
}
	
func cargarTarjetas(){ //FALTA VERIFICAR SI EL FORMATO DE LA FECHA ESTA BIEN Y FALTA EL LIMITE DE COMPRA
	_, err = db.Exec(`INSERT INTO tarjeta VALUES ('4382420954',81635249);
						INSERT INTO tarjeta VALUES ('7836666357653320',97824536,'201106','201108','1817','vigente');
						INSERT INTO tarjeta VALUES ('2732199710583851',16495823,'010503','010605','6701','vigente');
						INSERT INTO tarjeta VALUES ('9530652367572720',87512694,'300105','020408','4728','anulada');
						INSERT INTO tarjeta VALUES ('3695274119339368',58214936,'040510','230511','6423','vigente');
						INSERT INTO tarjeta VALUES ('6294033816643938',87219364,'050311','200312','2473','suspendida');
						INSERT INTO tarjeta VALUES ('6374432605814140',69254381,'020509','020510','3116','anulada');
						INSERT INTO tarjeta VALUES ('7947654982802386',24318769,'110507','110509','8178','suspendida');
						INSERT INTO tarjeta VALUES ('8743165676937175',67918245,'051107','051109','6915','vigente');
						INSERT INTO tarjeta VALUES ('8174730839100196',15624837,'230110','240111','5206','suspendida');
						INSERT INTO tarjeta VALUES ('4343577717377484',93527468,'111111','111212','8670','vigente');
						INSERT INTO tarjeta VALUES ('6025188452991960',82974315,'020310','030412','4864','vigente');
						INSERT INTO tarjeta VALUES ('3016480348260525',48129563,'041112','051113','4568','anulada');
						INSERT INTO tarjeta VALUES ('9330693747869828',65317289,'030305','030408','5387','anulada');
						INSERT INTO tarjeta VALUES ('9420306211523591',49815267,'110402','010403','6194','suspendida');
						INSERT INTO tarjeta VALUES ('7229894669781604',51362874,'111102','121103','4894','vigente');
						INSERT INTO tarjeta VALUES ('2155972533112753',83691452,'290405','290406','4310','suspendida');
						INSERT INTO tarjeta VALUES ('6924033286851784',64723591,'110411','110412','8850','suspendida');
						INSERT INTO tarjeta VALUES ('4486467155848418',93167854,'230110','230212','1054','vigente');
						INSERT INTO tarjeta VALUES ('2779243321116675',84396721,'050311','050412','1778','anulada');
						INSERT INTO tarjeta VALUES ('9184549155934952',93167854,'230110','230211','1218','vigente');
						INSERT INTO tarjeta VALUES ('5333311040348954',84396721,'010503','010604','7991','anulada');`)
		if err != nil {
		log.Fatal(err)
		}
}




