<script lang="ts">
	/* Import Bootstrap 5 */
	import 'bootstrap/dist/css/bootstrap.min.css';

	const year: number = new Date().getFullYear();
</script>

<!-- Page Content -->
<slot />

<!-- Universal Footer -->
<footer class="d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top">
	<div class="col-md-4 d-flex align-items-center">
		<span class="mb-3 mb-md-0 text-body-secondary"
			>© {year} Pirita - Creado con ❤ por UpVent Technologies.
		</span>
	</div>
</footer>
