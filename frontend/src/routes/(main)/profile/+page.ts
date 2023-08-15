import { browser } from '$app/environment';
import AuthService from '$lib/services/AuthService';

export async function load() {
	if (browser) {
		const user = await AuthService.auth();
		return {
			user: user
		};
	}
}
