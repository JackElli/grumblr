import type { Page } from '@playwright/test';
import type { _Grumble } from '../../src/routes/(main)/grumbles/grumbles';

export async function getGlobalGrumbles(page: Page) {
	await page.route('http://localhost:3200/global', async (route) => {
		const grumbles = [
			{
				createdBy: 'user:1',
				message: 'this is the very first global grumble',
				dateCreated: '2023-06-03T20:24:35.060986337Z',
				type: 'global'
			}
		];
		await route.fulfill({
			body: JSON.stringify(grumbles)
		});
	});
}
