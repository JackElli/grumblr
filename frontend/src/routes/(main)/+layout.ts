import AuthService from '$lib/services/AuthService';
import { userStore } from '$lib/stores/userStore';
import { get } from 'svelte/store';

export async function load() {
	// This makes sure the user stays 'logged in'
	const user = get(userStore);
	if (!user) {
		try {
			const user = await AuthService.auth();
			userStore.set(user);
		} catch (error) {
			console.error(error);
		}
	}
}
