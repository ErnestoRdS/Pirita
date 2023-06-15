<script lang="ts">
	import PiritaText from '$lib/components/PiritaText.svelte';

	import Speedometer2 from 'svelte-bootstrap-icons/lib/Speedometer2.svelte';
	import TaxiFront from 'svelte-bootstrap-icons/lib/TaxiFront.svelte';
	import People from 'svelte-bootstrap-icons/lib/People.svelte';
	import CashCoin from 'svelte-bootstrap-icons/lib/CashCoin.svelte';
	import Compass from 'svelte-bootstrap-icons/lib/Compass.svelte';
</script>

<header>
	<div class="px-3 py-2 text-bg-dark border-bottom">
		<div class="container">
			<div
				class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start"
			>
				<a
					class="d-flex align-items-center my-2 my-lg-0 me-lg-auto text-white text-decoration-none"
					href="/panel"
				>
					<PiritaText />
				</a>
				<ul class="nav col-12 col-lg-auto my-2 justify-content-center my-md-0 text-small">
					<li>
						<a class="nav-link link-light text-decoration-none" href="#">
							<Speedometer2 class="bi d-block mx-auto mb-1" height={24} width={24} />
							Panel
						</a>
					</li>
					<li>
						<a class="nav-link link-light text-decoration-none" href="#">
							<TaxiFront class="bi d-block mx-auto mb-1" height={24} width={24} />
							Unidades
						</a>
					</li>
					<li>
						<a class="nav-link link-light text-decoration-none" href="#">
							<People class="bi d-block mx-auto mb-1" height={24} width={24} />
							Conductores
						</a>
					</li>
					<li>
						<a class="nav-link link-light text-decoration-none" href="#">
							<CashCoin class="bi d-block mx-auto mb-1" height={24} width={24} />
							Pagos
						</a>
					</li>
					<li>
						<a class="nav-link link-light text-decoration-none" href="#">
							<Compass class="bi d-block mx-auto mb-1" height={24} width={24} />
							Viajes
						</a>
					</li>
				</ul>
			</div>
		</div>
	</div>
</header>

<slot />
