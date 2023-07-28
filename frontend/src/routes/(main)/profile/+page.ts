import AuthService from '$lib/services/AuthService';

export async function load() {
	await AuthService.auth();
}
