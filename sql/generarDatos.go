package sql

import (
	"log"
)

func CargarDatos() {
	
	CargarClientes()
	CargarComercios()
	CargarTarjetas()
	
	
	if err != nil {
		log.Fatal(err)
	}
}


func CargarClientes() { 
	_, err = db.Exec(` INSERT INTO cliente VALUES (81635249,'Juan','Perez','Corrientes 159','271542570123');
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
	
func CargarComercios(){
	 _ , err = db.Exec(`INSERT INTO comercio VALUES (95169,'Disco','Las Heras 716','A9812CAV','872178944032');
						INSERT INTO comercio VALUES (29981,'Carrefour','Av T A Edison 555','B7466EWR','457755059512');
						INSERT INTO comercio VALUES (82211,'Garbarino','Avenida 44 2049','C9275CVB','625050693943');
						INSERT INTO comercio VALUES (79701,'Coto','Aberastain Sur 163','D0991LKJ','709777835612');
						INSERT INTO comercio VALUES (21724,'Falabella','Camargo 775','E7347CVF','138988075243');
						INSERT INTO comercio VALUES (51249,'Grido Helado','H Yrigoyen 1659','E0179VAG','361518160053');
						INSERT INTO comercio VALUES (87682,'Fravega','Tagle 3383','F8627MAX','925173628765');
						INSERT INTO comercio VALUES (59460,'Sodimac','Av Entre Ríos 1072','G7860TOM','915082365243');
						INSERT INTO comercio VALUES (34039,'Panaderia y confiteria','Av Colón 1012','H3369LUN','634483423800');
						INSERT INTO comercio VALUES (32694,'Verduleria el tomate','Calle 55 3067','I9628VBF','713519136410');
						INSERT INTO comercio VALUES (77289,'Fiambreria el gaucho','Acc Alte Brown 971','J7473HEG','726870604311');
						INSERT INTO comercio VALUES (68463,'Panaderia el pan','C Melo 4708','K8005LVB','988294290165');
						INSERT INTO comercio VALUES (75374,'Carniceria el carni','Brown 308','L1589MCB','284719172711');
						INSERT INTO comercio VALUES (79751,'Panaderia maxi','Av R S Ortiz 160','M3138MMM','267532588765');
						INSERT INTO comercio VALUES (38833,'Supermercado asia','Thompson 309','O2787AAA','777139421953');
						INSERT INTO comercio VALUES (68199,'Compumundo','P Molina 133','P4724PMJ','670898887134');
						INSERT INTO comercio VALUES (63129,'Verduleria sandra','Bogotá 2842','Q9276KMK','292802798211');
						INSERT INTO comercio VALUES (90949,'Solo deportes','Garibaldi 155','R6548CXV','333826274011');
						INSERT INTO comercio VALUES (68806,'Panaderia pan dulce','Chile 329','S7303BXE','946007831954');
						INSERT INTO comercio VALUES (61351,'Fiambreria luna','Av L Quaranta 7091','T6863CSD','339534886414');`)
		if err != nil {
			log.Fatal(err)
		}
	
}
	
func CargarTarjetas(){ //decimal(t,c) total-cant decimales
	_, err = db.Exec(`INSERT INTO tarjeta VALUES ('4382420954476737',81635249,'201205','201306','9184',50,'vigente');
						INSERT INTO tarjeta VALUES ('7836666357653320',97824536,'201106','201205','1817',40,'vigente');
						INSERT INTO tarjeta VALUES ('2732199710583851',16495823,'201804','202005','6701',45,'vigente');
						INSERT INTO tarjeta VALUES ('9530652367572720',87512694,'201905','202104','4728',70,'anulada');
						INSERT INTO tarjeta VALUES ('3695274119339368',58214936,'201611','201811','6423',10,'vigente');
						INSERT INTO tarjeta VALUES ('6294033816643938',87219364,'200903','201103','2473',34,'vigente');
						INSERT INTO tarjeta VALUES ('6374432605814140',69254381,'201012','201310','3116',17,'anulada');
						INSERT INTO tarjeta VALUES ('7947654982802386',24318769,'200806','201010','8178',30,'suspendida');
						INSERT INTO tarjeta VALUES ('8743165676937175',67918245,'200403','200504','6915',10,'vigente');
						INSERT INTO tarjeta VALUES ('8174730839100196',15624837,'201801','202001','5206',30,'suspendida');
						INSERT INTO tarjeta VALUES ('4343577717377484',93527468,'201410','201501','8670',10,'vigente');
						INSERT INTO tarjeta VALUES ('6025188452991960',82974315,'201305','201405','4864',90,'vigente');
						INSERT INTO tarjeta VALUES ('3016480348260525',48129563,'201601','201701','4568',15,'anulada');
						INSERT INTO tarjeta VALUES ('9330693747869828',65317289,'201004','201405','5387',52,'vigente');
						INSERT INTO tarjeta VALUES ('9420306211523591',49815267,'200712','200901','6194',70,'suspendida');
						INSERT INTO tarjeta VALUES ('7229894669781604',51362874,'200510','200701','4894',10,'vigente');
						INSERT INTO tarjeta VALUES ('2155972533112753',83691452,'201901','202301','4310',20,'suspendida');
						INSERT INTO tarjeta VALUES ('6924033286851784',64723591,'201703','202001','8850',14,'vigente');
						INSERT INTO tarjeta VALUES ('4486467155848418',93167854,'200501','200701','1054',00,'vigente');
						INSERT INTO tarjeta VALUES ('9184549155934952',93167854,'201801','201905','1218',32,'vigente');
						INSERT INTO tarjeta VALUES ('2779243321116675',84396721,'201503','201706','1778',00,'vigente');
						INSERT INTO tarjeta VALUES ('5333311040348954',84396721,'201608','201904','7991',50,'anulada');`)
		if err != nil {
			log.Fatal(err)
		}
}

						
//Cliente 84396721 tiene dos tarjetas
//Cliente 93167854 tiene dos tarjetas


