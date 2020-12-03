package sql

import {
	"log"
}

func generarResumen(){
	_, err = db.Query(`create function generarResumen(cliente int, anio int,mes int )) returns bool as $$
			declare
				idResumen int;
				totalPagar decimal(7,2) : = 0;
				tarjeta char(16);
					
				begin
				
				

				FOR tarjeta IN 
					Select nrotarjeta From tarjeta Where nrocliente = cliente and estado = 'vigente'
				LOOP

					-- GENERO LA CABECERA
				
					INSERT INTO cabecera (nombre,apellido,domicilio,nrotarjeta,desde,hasta,vence)
					SELECT cl.nombre,cl.apellido,cl.domicilio,t.nrotarjeta,c.fechainicio,c.fechacierre,c.fechavto
						FROM public.tarjeta t, public.cierre c, public.cliente cl WHERE SUBSTRING (t.nrotarjeta),1)::int = c.terminacion
						and cl.nrocliente = t.nrocliente
						and t.nrotarjeta = tarjeta
						and c.año = anio
						and c.mes = mes
					RETURNING nroresumen INTO idResumen
					
					if (idResumen is null) then
						raise 'No se pudo generar el resumen, Cliente inexistente';
						return False;
					end if;

					--GENERAR DETALLE

					INSERT INTO detalle (nroresumen,nrolinea,fecha,nombrecomercio,monto)
					SELECT idResumen, ROW_NUMBER() OVER (ORDER BY t.nrotarjeta as nrolinea,co.fecha,com.nombre,co.monto
					FROM public.tarjeta t, public.cierre c, public.compra co, public.comercio com
					WHERE SUBSTRING (t.nrotarjeta, LENGTH(t.nrotarjeta),1)::int = c.terminacion
					and co.nrotarjeta = t.nrotarjeta
					and com.nrocomercio = co.nrocomercio
					and t.nrotarjeta = tarjeta
					and c.año  = anio
					and c.mes = mes
					and co.fecha >= c.fechainicio
					and co.fecha <= c.fechacierre;

					if (lastval() is NULL) then
						raise 'No se pudo generar el resumen';
						return False;
					end if;

					-- ACTUALIZAR RESUMEN

					-- coalesce devuelve el primer valor no null 

					totalPagar := (SELECT SUM(monto) FROM detalle WHERE nroresumen = idResumen GROUP BY nroresumen);

					UPDATE cabecera
					set total  = COALESCE (NULLIF(totalPagar, 0),0) WHERE nroresumen = idResumen;

					--CAMBIAR PAGADO A TRUE

					FOR _linea in SELECT * FROM public.tarjeta t, public.cierre c, public.compra co, public.comercio com WHERE SUBSTRING 
						(t.nrotarjeta, LENGTH(t.nrotarjeta),1) :: int = c.terminacion 
						and co.nrotarjeta = t.nrotarjeta 
						and com.nrocomercio = co.nrocomercio 
						and t.nrotarjeta = tarjeta
				LOOP
					UPDATE compra set pagado = True where nrotarjeta=_linea.nrotarjeta and monto = _linea.monto;
				END LOOP;

				return TRUE;

					end;
			$$ language plpgsql;`)


		if err != nil {
			log.Fatal(err)
		}

		

					
					

				
				





}
