<script lang="ts">
	import { user } from '$lib/helpers/store';
	import { goto } from '$app/navigation';

	import PiritaCar from '$lib/components/PiritaCar.svelte';
	import PiritaText from '$lib/components/PiritaText.svelte';
	import PiritaTextSecondary from '$lib/components/PiritaTextSecondary.svelte';

	let usuario = '';
	let password = '';

	const login = async () => {
		const response = await fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				usuario,
				password
			})
		});

		const data = await response.json();

		if (data.token) {
			user.update(() => {
				return {
					token: data.token,
					role: data.type,
					isAuthenticated: true
				};
			});
			if (data.type === 'admin') {
				goto('/panel');
			} else if (data.type === 'conductor') {
				goto('/conductor');
			}
		} else {
			alert('Inicio de sesión fallido: Usuario o contraseña incorrectos');
		}
	};
</script>

<section class="h-50 text-center pirita-head">
	<div class="circle" />
</section>

<section class="d-flex text-center pirita-body">
	<section class="container-fluid">
		<div class="container my-5">
			<PiritaCar />
			<PiritaText />
			<PiritaTextSecondary />
		</div>
		<div class="container py-4">
			<form>
				<input
					class="form-control form-control-lg my-1"
					bind:value={usuario}
					type="text"
					placeholder="Usuario"
				/>
				<input
					class="form-control form-control-lg my-2"
					bind:value={password}
					type="password"
					placeholder="Contraseña"
				/>
				<button class="btn btn-dark btn-lg mb-3" on:click={login}>Iniciar sesión</button>
			</form>
		</div>
		<div class="container">
			<hr />
			<h3 class="text-light fw-bold small">¿AÚN NO TIENES CUENTA?</h3>
			<h4 class="text-light fw-bold small">Pide a tu patrón que te registre.</h4>
		</div>
	</section>
</section>

<style>
	.pirita-body {
		background-color: #839ab4;
	}

	.pirita-head {
		background-color: #fafafa;
	}

	.circle {
		width: 7em;
		height: 7em;
		background-color: #a57d31;
		border-radius: 50%;
		margin: 0 auto;
		margin-top: 5%;
		margin-bottom: -2em;
		/* Always on top */
		position: relative;
		z-index: 1;
	}
</style>
