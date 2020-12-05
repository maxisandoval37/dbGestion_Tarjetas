package sql

import (
	"log"
)




func GenerarResumenesPrincipal () {
	generarResumen()	
	generarResumenes2()
}


func generarResumenes2() {
	_, err = db.Query(`
				SELECT generar_resumen(81635249, 2021, 5);
				SELECT generar_resumen(97824536, 2021, 12);
				SELECT generar_resumen(16495823, 2021, 12);
				SELECT generar_resumen(87512694, 2021, 12);
				SELECT generar_resumen(58214936, 2021, 12);
				SELECT generar_resumen(87219364, 2021, 6);
				SELECT generar_resumen(67918245, 2021, 6);
				SELECT generar_resumen(93527468, 2021, 6);
				SELECT generar_resumen(84396721, 2021, 11);
				`)
	if err != nil {
		log.Fatal(err)
	}
}


func generarResumen() {
	_, err = db.Query(
		`create or replace function generar_resumen(n_cliente cliente.nrocliente%type,
						anio_par int,
						mes_par int) returns void as $$
	declare
		cliente_encontrado record;
		compra_aux record;
		tarjeta_aux record;
		cierre_aux record;
		total_aux cabecera.total%type;
		nroresumen_aux cabecera.nroresumen%type;
		nombre_comercio comercio.nombre%type;
		cont int := 1;

	begin
		SELECT * INTO cliente_encontrado FROM cliente WHERE nrocliente = n_cliente;
		  if not found then
	      		  raise 'Cliente % no existe.', n_cliente;
  		  end if;
		
		for tarjeta_aux in select * FROM tarjeta WHERE nrocliente = n_cliente loop

			total_aux := 0;
			SELECT * INTO cierre_aux FROM cierre cie WHERE cie.año = anio_par and cie.mes = mes_par and cie.terminacion = substring(tarjeta_aux.nrotarjeta, 16, 1)::int;

			INSERT INTO cabecera(nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence) 
					values (cliente_encontrado.nombre, cliente_encontrado.apellido, cliente_encontrado.domicilio,tarjeta_aux.nrotarjeta, cierre_aux.fechainicio, cierre_aux.fechacierre,cierre_aux.fechavto);

			SELECT INTO nroresumen_aux nroresumen FROM cabecera where nrotarjeta = tarjeta_aux.nrotarjeta
									and desde = cierre_aux.fechainicio
									and hasta = cierre_aux.fechacierre;

			for compra_aux in SELECT * FROM compra WHERE nrotarjeta = tarjeta_aux.nrotarjeta 
								and fecha::date >= (cierre_aux.fechainicio)::date 
								and fecha::date <= (cierre_aux.fechacierre)::date
								and pagado = false loop
				
				nombre_comercio := (SELECT nombre FROM comercio where nrocomercio = compra_aux.nrocomercio);
				INSERT INTO detalle values (nroresumen_aux, cont, compra_aux.fecha, nombre_comercio, compra_aux.monto);
				total_aux := total_aux + compra_aux.monto;
				cont := cont + 1;
				update compra set pagado = true where nrooperacion = compra_aux.nrooperacion;
		
			end loop;
			
			update cabecera set total = total_aux where nrotarjeta = tarjeta_aux.nrotarjeta
									and desde = cierre_aux.fechainicio
									and hasta = cierre_aux.fechacierre;
		end loop;
	end;
$$ language plpgsql;`)
	if err != nil {
		log.Fatal(err)
	}
}







