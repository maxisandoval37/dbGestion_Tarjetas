package sql

import (
	"log"
)

func spGenerarCierres() {
	_, err = db.Query(
		`CREATE OR REPLACE function generarCierres(anio int) returns void as $$
		Declare
			fdesde date;
			fhasta date;
			fvto date;
			anioanterior int;
		BEGIN
			anioanterior = (anio-1);
			FOR ultimo IN 0 .. 9 BY 1
			LOOP				
				FOR mes IN 1 .. 12 BY 1
				LOOP	
					SELECT INTO fdesde to_date( anioanterior::text || mes, 'YYYYMM' ) 
						+ CAST(ultimo::text ||' days' AS interval);
					SELECT INTO fhasta fdesde::date + CAST('1 months' AS interval);
					SELECT INTO fvto fhasta::date + cast('7 days' as interval);
					
					INSERT INTO cierre VALUES(anio, mes, ultimo, fdesde, fhasta, fvto);
				END LOOP;
			END LOOP;
		END;
		$$ language plpgsql;`)

	if err != nil {
		log.Fatal(err)
	}
}

func cierres2020() {
	_, err = db.Exec(`select generarCierres(2020);`)
	if err != nil {
		log.Fatal(err)
	}
}

func Cierres(){
	spGenerarCierres();
	cierres2020();
}

