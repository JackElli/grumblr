import type { Page } from '@playwright/test';

export async function getGlobalGrumbles(page: Page) {
	await page.route('http://localhost:3200/global/*', async (route) => {
		await route.fulfill({
			json: [
				{
					createdBy: 'user:1',
					message: 'this is the very first global grumble',
					dateCreated: '2023-06-03T20:24:35.060986337Z',
					type: 'global',
					agrees: {},
					disagrees: {},
					comments: [],
					dataType: 'text'
				}
			]
		});
	});
}

export async function getGlobalCategories(page: Page) {
	await page.route('http://localhost:3200/grumbles/info/categories/global', async (route) => {
		await route.fulfill({
			json: [
				{
					id: '',
					type: 'global',
					people: [],
					name: 'Amazing'
				},
				{
					id: '',
					type: 'global',
					people: [],
					name: 'Random'
				}
			]
		});
	});
}
