package sql

import (
	"log"
)
//store procedure
func generarCierres() {
	_, err = db.Query(
		`CREATE OR REPLACE function generarCierres(anio int) returns void as $$
		Declare
			fdesde date;
			fhasta date;
			fvto date;
		BEGIN
			FOR tarjeta IN 0 .. 9 BY 1
			LOOP
				SELECT INTO fdesde to_date((anio - 1)::text || '12' || (select 23 + trunc(random() * 4))::text, 'YYYYMMDD');
				SELECT INTO fhasta fdesde + cast((select 29 + trunc(random() * 2))::text || ' days' as interval);
				SELECT INTO fvto fhasta + cast('10 days' as interval);to_date((anio - 1)::text


				FOR mes IN 1 .. 12 BY 1
				LOOP			
					INSERT INTO cierre VALUES(anio,mes,tarjeta,fdesde,fhasta,fvto);
					SELECT INTO fdesde fhasta + cast('1 days' as interval);
					SELECT INTO fhasta fdesde + cast((select 29 + trunc(random() * 2))::text || ' days' as interval);
					SELECT INTO fvto fhasta + cast('10 days' as interval);
				END LOOP;

			END LOOP;
		END;
		$$ language plpgsql;`)

	if err != nil {
		log.Fatal(err)
	}
}

