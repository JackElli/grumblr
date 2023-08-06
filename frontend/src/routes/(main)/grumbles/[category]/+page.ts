import { browser } from '$app/environment';
import AuthService from '$lib/services/AuthService.js';
import GrumbleService from '$lib/services/GrumbleService';

export async function load({ params }) {
	if (browser) {
		// if user doesn't currently have any friends, they cannot
		// see friends grumbles
		const user = await AuthService.auth();
		if (!user.friends || user.friends.length == 0) {
			return {
				friends: 0
			};
		}

		// Can we run the next two things in parallel?
		let grumbles;
		try {
			grumbles = await GrumbleService.list(params.category);
		} catch (e) {
			console.log(e);
			return {
				error: `Unable to fetch grumbles ${e}`
			};
		}

		let categories;
		try {
			categories = await GrumbleService.getCategories('friends');
		} catch (e) {
			console.log(e);
			return {
				error: `Unable to fetch categories ${e}`
			};
		}

		categories = categories.filter((category) => category.people.includes(user.id));

		return {
			grumbles: grumbles ?? [],
			categories: categories ?? [],
			currentCategory: params.category
		};
	}
}
