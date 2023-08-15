import { browser } from '$app/environment';
import GrumbleService from '$lib/services/GrumbleService';

export async function load({ params }) {
	if (browser) {
		const grumbles = GrumbleService.listGlobal(params.category);
		const categories = GrumbleService.categories('global');

		const [g, c] = await Promise.all([grumbles, categories]);

		return {
			grumbles: g,
			categories: c,
			currentCategory: params.category
		};
	}
}
