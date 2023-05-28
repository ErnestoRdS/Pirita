interface Conductor {
  id: number;
  nombre: string;
  apellidos: string;
  curp: string;
  clave_ine: string;
  salario: number;
  estado: string;
  pagos?: Pago[];
  usuario: string;
  correo: string;
  password: string;
}

export default Conductor;
