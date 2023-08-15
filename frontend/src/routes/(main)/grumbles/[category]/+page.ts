import { browser } from '$app/environment';
import AuthService from '$lib/services/AuthService.js';
import GrumbleService from '$lib/services/GrumbleService';
import type { User } from '$lib/stores/userStore';

export async function load({ params }) {
	if (browser) {
		let user: User | undefined;
		user = await AuthService.auth();

		const grumbles = await GrumbleService.list(params.category);
		let categories = await GrumbleService.categories('friends');
		categories = categories?.filter((category) => category.people.includes(user?.id ?? '?'));

		return {
			grumbles: grumbles,
			categories: categories,
			currentCategory: params.category
		};
	}
}
