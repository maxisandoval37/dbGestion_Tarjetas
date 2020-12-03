package sql

import (
	"log"
)

func spAutorizarCompra() {
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION autorizarcompra(_nrotarjeta char(16), _codseguridad char(4), _nrocomercio int, _monto decimal(7,2)) returns bool as $$
		 DECLARE
			totalPendiente decimal(8,2);
			montoMaximo decimal(8,2);
			fechaVenceTarjeta int;
			fechaVence date;

		 BEGIN
			PERFORM * FROM tarjeta WHERE nrotarjeta = _nrotarjeta and estado = 'vigente';
			if (not found) THEN
				PERFORM agregarrechazo(CAST(_nrotarjeta as char(16)), CAST(_nrocomercio as int), CAST(current_timestamp as timestamp), CAST(_monto as decimal(7,2)), CAST('Tarjeta no válida' as text));
				return False;
			END IF;

			PERFORM * FROM tarjeta WHERE nrotarjeta=_nrotarjeta and codseguridad=_codseguridad;
			if (not found) THEN
				PERFORM agregarrechazo(CAST(_nrotarjeta as char(16)), CAST(_nrocomercio as int), CAST(current_timestamp as timestamp), CAST(_monto as decimal(7,2)), CAST('Código de seguridad inválido' as text));
				return False;
			END IF;

			totalPendiente:= (SELECT sum(monto) FROM compra WHERE nrotarjeta =_nrotarjeta and pagado=False);
			montoMaximo:= (SELECT limitecompra FROM tarjeta WHERE nrotarjeta=_nrotarjeta);
			if(totalpPendiente is null and _monto > montoMaximo OR totalPendiente is not null and totalPendiente + _monto > montoMaximo) THEN
				PERFORM agregarrechazo(CAST(_nrotarjeta as char(16)), CAST(_nrocomercio as int), CAST(current_timestamp as timestamp), CAST(_monto as decimal(7,2)), CAST('Supera límite de tarjeta' as text));
				return False;
			END IF;

			SELECT validahasta INTO fechaVenceTarjeta FROM tarjeta WHERE nrotarjeta=_nrotarjeta;
			SELECT INTO FechaVence TO_DATE(fechaVenceTarjeta ||'01','YYYYMMDD');
			SELECT INTO FechaVence (FechaVence +  interval '1 month')::date;
			if (FechaVence < current_date) THEN
			PERFORM agregarrechazo(CAST(_nrotarjeta as char(16)), CAST(_nrocomercio as int), CAST(current_timestamp as timestamp), CAST(_monto as decimal(7,2)), CAST('Plazo de vigencia expirado' as text));
				return False;
			END IF;

			PERFORM * FROM tarjeta WHERE nrotarjeta=_nrotarjeta and estado='suspendida';
			if (found) THEN
				PERFORM agregarrechazo(CAST(_nrotarjeta as char(16)), CAST(_nrocomercio as int), CAST(current_timestamp as timestamp), CAST(_monto as decimal(7,2)), CAST('La tarjeta se encuentra suspendida' as text));
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

func spAgregarRechazo() {
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION agregarrechazo(_nrotarjeta char(16),_nrocomercio int, _fecha timestamp,_monto decimal(7,2),_motivo text) returns void as $$
		DECLARE
			numerorechazo int;
		BEGIN
			INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto, _motivo)
			RETURNING nrorechazo INTO numerorechazo;
		END;

	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}
