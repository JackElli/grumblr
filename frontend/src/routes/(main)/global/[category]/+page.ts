import { browser } from '$app/environment';
import GrumbleService from '$lib/services/GrumbleService';

export async function load({ params }) {
	if (browser) {
		let grumbles;
		try {
			grumbles = await GrumbleService.listGlobal(params.category);
		} catch (e) {
			console.log(e);
			return {
				error: `Unable to fetch grumbles ${e}`
			};
		}

		let categories;
		try {
			categories = await GrumbleService.getCategories('global');
		} catch (e) {
			console.log(e);
			return {
				error: `Unable to fetch categories ${e}`
			};
		}

		return {
			grumbles: grumbles ?? [],
			categories: categories ?? [],
			currentCategory: params.category
		};
	}
}
