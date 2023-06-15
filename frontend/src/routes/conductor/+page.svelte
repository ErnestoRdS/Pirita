<script lang="ts">
	import Cash from 'svelte-bootstrap-icons/lib/Cash.svelte';
	import Globe from 'svelte-bootstrap-icons/lib/Globe.svelte';
	import BarChartLine from 'svelte-bootstrap-icons/lib/BarChartLine.svelte';

	// Rangos de Conductor (svg)
	import Pirita from '$lib/assets/pirita.svg';
	import Oro from '$lib/assets/oro.svg';
	import Platino from '$lib/assets/platino.svg';
	import Premium from '$lib/assets/premium.svg';

	const usuario = 'Juan';
	// Crear una constante "usuario_id" que sea un UUID
	const usuario_id = '123e4567-e89b-12d3-a456-426614174000';

	// Ganancias y Viajes
	let ganancias = '$1,000.00 MXN';
	let viajes = 10;

	// Rango del conductor (puede ser uno: "Pirita" <  "Oro" < "Platino" < "Premium")
	let rango = 'Pirita';

	function calcularRango(viajes: number) {
		if (viajes >= 0 && viajes <= 10) {
			rango = 'Pirita';
		} else if (viajes >= 11 && viajes <= 20) {
			rango = 'Oro';
		} else if (viajes >= 21 && viajes <= 30) {
			rango = 'Platino';
		} else if (viajes >= 31) {
			rango = 'Premium';
		}
	}

	calcularRango(viajes);
</script>

<section class="pirita-body container-fluid my-5">
	<div class="container my-3 text-center">
		<div class="container m-3">
			<svg
				class="bi border my-3 rounded-circle p-3"
				xmlns="http://www.w3.org/2000/svg"
				width="128"
				height="128"
				viewBox="0 0 24 24"
				><path
					fill="currentColor"
					d="M7.39 16.539a8 8 0 1 1 9.221 0l2.083 4.76a.5.5 0 0 1-.459.701H5.765a.5.5 0 0 1-.459-.7l2.083-4.761Zm.729-5.569a4.002 4.002 0 0 0 7.763 0l-1.941-.485a2 2 0 0 1-3.882 0l-1.94.485Z"
				/></svg
			>
		</div>

		<p class="h1 fw-bold text-light">Bienvenid@ {usuario}</p>
		<small class="text-muted">Nombre completo: Jose Joestrella López Pérez</small>
		<br />
		<small class="text-muted small">ID de Conductor: {usuario_id}</small>

		<div class="container my-2">
			<span class="h1 fw-bold">Tu rango actual es:</span>
			{#if rango == 'Pirita'}
				<img class="img-fluid" src={Pirita} alt="Pirita" width="300" height="300" />
			{:else if rango == 'Oro'}
				<img class="img-fluid" src={Oro} alt="Oro" width="300" height="300" />
			{:else if rango == 'Platino'}
				<img class="img-fluid" src={Platino} alt="Platino" width="300" height="300" />
			{:else if rango == 'Premium'}
				<img class="img-fluid" src={Premium} alt="Premium" width="300" height="300" />
			{/if}
		</div>
	</div>

	<!-- TODO: Arreglar estas tarjetas en el centro, con espaciado y los
         datos en medio con un ícono de Bootstrap
    -->
	<div class="container text-center p-2">
		<div class="row gx-5 gy-2">
			<div class="col bg-light rounded-3 p-4 shadow">
				<div class="container">
					<BarChartLine height={50} width={50} />
				</div>
				Mis Ganancias: {ganancias}
			</div>
			<div class="col bg-light rounded-3 p-4 shadow">
				<div class="container">
					<Globe height={50} width={50} />
				</div>
				Mis Viajes: {viajes}
			</div>
		</div>
	</div>
</section>

<section class="container">
	<section class="container text-center shadow-lg p-5 rounded-4">
		<p class="h1">¿Qué deseas hacer?</p>
		<p class="small">Presiona un botón para elegir una opción.</p>

		<div class="d-grid gap-2">
			<button class="btn btn-success" type="button"
				>Registrar Viaje <Cash height={24} width={24} /></button
			>
			<button class="btn btn-primary" type="button">Ver mis pagos</button>
		</div>
	</section>
</section>

<style>
	.pirita-body {
		background-color: #839ab4;
	}
</style>
