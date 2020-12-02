package sql

import (
	"log"
)

func spAutorizarCompra() {
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION autorizarcompra(_nrotarjeta char(16),_codseguridad char(4),_nrocomercio int,_monto decimal(7,2)) returns bool as $$
		 DECLARE

		 BEGIN
			PERFORM * FROM tarjeta WHERE nrotarjeta = _nrotarjeta and estado = 'vigente';
			if (not found) THEN
				PERFORM rechazo_invalida(CAST(_nrotarjeta as char(16)),CAST(_nrocomercio as int),CAST(current_timestamp as timestamp),CAST(_monto as decimal(7,2)));
				return False;
			END IF;

			PERFORM * FROM tarjeta WHERE nrotarjeta = _nrotarjeta and codseguridad = _codseguridad;
			if (not found) THEN
				PERFORM rechazo_codigo(CAST(_nrotarjeta as char(16)),CAST(_nrocomercio as int),CAST(current_timestamp as timestamp),CAST(_monto as decimal(7,2)));
				return False;
			END IF;

			PERFORM * FROM tarjeta WHERE nrotarjeta = _nrotarjeta and estado = 'suspendida';
			if (found) THEN
				PERFORM rechazo_suspendida(CAST(_nrotarjeta as char(16)),CAST(_nrocomercio as int),CAST(current_timestamp as timestamp));
				return False;
			END IF;

			INSERT INTO compra(nrotarjeta, nrocomercio, fecha, monto, pagado) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto,False);
			return True;
		END;
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}

func spAgregarRechazos() {
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION rechazo_invalida(_nrotarjeta char(16),_nrocomercio int, _fecha timestamp,_monto decimal(7,2)) returns void as $$
		DECLARE
			numerorechazo int;
		BEGIN
			INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto, 'Tarjeta no válida')
			RETURNING nrorechazo INTO numerorechazo;
		END;
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION rechazo_codigo(_nrotarjeta char(16),_nrocomercio int, _fecha timestamp,_monto decimal(7,2)) returns void as $$
		DECLARE
			numerorechazo int;
		BEGIN
			INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto, 'Código de seguridad inválido')
			RETURNING nrorechazo INTO numerorechazo;
		END;
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION rechazo_suspendida(_nrotarjeta char(16),_nrocomercio int, _fecha timestamp,_monto decimal(7,2)) returns void as $$
		DECLARE
			numerorechazo int;
		BEGIN
			INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto, 'La tarjeta se encuentra suspendida')
			RETURNING nrorechazo INTO numerorechazo;
		END;
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}


