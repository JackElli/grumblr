import type { Page } from '@playwright/test';

export async function getGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles', async (route) => {
		const grumbles = [
			{
				createdBy: 'user:1',
				message: 'this is the very first grumble',
				dateCreated: '2023-06-03T20:24:35.060986337Z'
			}
		];
		await route.fulfill({
			body: JSON.stringify(grumbles)
		});
	});
}
