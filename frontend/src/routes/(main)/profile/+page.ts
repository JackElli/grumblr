import { browser } from '$app/environment';
import AuthService from '$lib/services/AuthService';
import GrumbleService from '$lib/services/GrumbleService';

export async function load() {
	if (browser) {
		const user = await AuthService.auth();
		const grumblesByUser = await GrumbleService.listByUser(user.id);
		return {
			user: user,
			grumbles: grumblesByUser
		};
	}
}
