// Store es un "servicio" de Svelte, que nos va a permitir almacenar
// datos y compartirlos entre componentes.
import { writable } from 'svelte/store';

// El objeto "user" es un "store" de Svelte, donde vamos a almacenar
// los datos del usuario que se ha logueado en la aplicación.
export const user = writable({
	token: '', // Token de autenticación
	role: '', // Rol del usuario
	isAuthenticated: false // Indica si el usuario está autenticado
});
