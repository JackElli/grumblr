import type { Page } from '@playwright/test';

export async function getUser(page: Page) {
	await page.route('http://localhost:3200/user/*', async (route) => {
		await route.fulfill({
			json: {
				id: '1f21823a-8682-4900-b627-d6bd39e1b95b',
				username: 'test1',
				friends: [
					{
						test: 'asdasd'
					}
				],
				welcome: true
			}
		});
	});
}

export async function getUserNoFriends(page: Page) {
	await page.route('http://localhost:3200/user/*', async (route) => {
		await route.fulfill({
			json: {
				id: '1f21823a-8682-4900-b627-d6bd39e1b95b',
				username: 'test1',
				friends: [],
				welcome: true
			}
		});
	});
}
