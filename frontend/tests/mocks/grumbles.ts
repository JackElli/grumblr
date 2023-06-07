import type { Page } from '@playwright/test';
import type { _Grumble } from '../../src/routes/(main)/grumbles/grumbles';

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

export async function getNoGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles', async (route) => {
		const grumbles: _Grumble[] = [];
		await route.fulfill({
			body: JSON.stringify(grumbles)
		});
	});
}

export async function getLongGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles', async (route) => {
		const grumbles = [
			{
				createdBy: 'user:1',
				message:
					'This is a very long grumble This is a very long grumble This is a very long grumble This is a very long grumble This is a very long grumble',
				dateCreated: '2023-06-03T20:24:35.060986337Z'
			}
		];
		await route.fulfill({
			body: JSON.stringify(grumbles)
		});
	});
}
