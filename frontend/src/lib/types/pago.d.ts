interface Pago {
	id: number;
	conductor_id: number;
	fecha: string;
	cantidad: number;
	notas: string;
}

export default Pago;
