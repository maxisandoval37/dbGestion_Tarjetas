package sql

import (
	"log"
)


func GenerarResumen() {
	_, err = db.Query(
		`create or replace function generarResumen(cliente int, anioR int, mesR int) returns bool as $$
		declare
			   idResumen int;
			   totalPagar decimal(7,2) := 0;
			   _linea record;
			   tarjeta char(16);
		   
				begin		
				
				FOR tarjeta IN 
					Select nrotarjeta
					From Tarjeta
					Where nrocliente = cliente
					and estado = 'vigente'
				LOOP
				
				-- 	Generar Cabecera
					INSERT INTO cabecera (nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence) 
					SELECT cli.nombre, cli.apellido, cli.domicilio, t.nrotarjeta, c.fechainicio, c.fechacierre, c.fechavto
						FROM public.tarjeta t, public.cierre c, public.cliente cli
						WHERE SUBSTRING (t.nrotarjeta, LENGTH(t.nrotarjeta), 1)::int = c.terminacion
						and cli.nrocliente = t.nrocliente
						and t.nrotarjeta = tarjeta
						and c.año = anioR
						and c.mes = mesR
					RETURNING nroresumen INTO idResumen;
				
					if (idResumen is null) then
						raise 'No se pudo generar el resumen, Cliente inexistente';
						return False;
					end if;	
		
				-- Generar detalle	
					INSERT INTO detalle (nroresumen, nrolinea, fecha, nombrecomercio, monto) 
					SELECT idResumen, ROW_NUMBER () OVER (ORDER BY t.nrotarjeta) as nrolinea, co.fecha, com.nombre , co.monto
					FROM public.tarjeta t, public.cierre c, public.compra co, public.comercio com
					WHERE SUBSTRING (t.nrotarjeta, LENGTH(t.nrotarjeta), 1)::int = c.terminacion
					and co.nrotarjeta = t.nrotarjeta
					and com.nrocomercio = co.nrocomercio
					and t.nrotarjeta = tarjeta
					and c.año = anioR
					and c.mes = mesR
					and co.fecha >= c.fechainicio 
					and co.fecha <= c.fechacierre;	
				
					if (lastval() is NULL) then
						raise 'No se pudo generar el resumen';
						return False;
					end if;	
			
				-- Actualizar Resumen
					totalPagar := (SELECT SUM(monto) 
								  FROM detalle 
								  WHERE nroresumen = idResumen
								  GROUP BY nroresumen);
					 
					UPDATE cabecera 
					set total = COALESCE(NULLIF(totalPagar, 0), 0)
					WHERE nroresumen = idResumen;	
					
				--Cambiar pagado a True
					FOR _linea in SELECT * FROM public.tarjeta t, public.cierre c, public.compra co, public.comercio com
						WHERE SUBSTRING (t.nrotarjeta, LENGTH(t.nrotarjeta), 1)::int = c.terminacion
						and co.nrotarjeta = t.nrotarjeta
						and com.nrocomercio = co.nrocomercio
						and t.nrotarjeta = tarjeta 
					LOOP
						UPDATE compra set pagado = True where nrotarjeta=_linea.nrotarjeta and monto=_linea.monto;									
					END LOOP;
				
				END LOOP;
				
				return True;
				
				   end;
		$$ language plpgsql;`)
	if err != nil {
		log.Fatal(err)
	}
}
