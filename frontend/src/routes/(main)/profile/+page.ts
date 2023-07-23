import { Auth } from '$lib/services/AuthService';

export async function load() {
	const user = Auth();
	return {
		user: user
	};
}
