import { browser } from '$app/environment';
import AuthService from '$lib/services/AuthService.js';
import GrumbleService from '$lib/services/GrumbleService';
import { userStore } from '$lib/stores/userStore';
import { get } from 'svelte/store';

export async function load({ params }) {
	if (browser) {
		const grumbles = await GrumbleService.list(params.category);
		let categories = await GrumbleService.categories('friends');
		const user = get(userStore);
		categories = categories?.filter((category) => category.people.includes(user?.id ?? '?'));

		return {
			grumbles: grumbles,
			categories: categories,
			currentCategory: params.category
		};
	}
}
