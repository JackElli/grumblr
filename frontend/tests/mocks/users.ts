import type { Page } from '@playwright/test';

export async function getUser(page: Page) {
	await page.route('http://localhost:3200/user/*', async (route) => {
		const user = {
			id: '1f21823a-8682-4900-b627-d6bd39e1b95b',
			username: 'test1',
			password: 'test',
			dateCreated: '2023-06-13T19:48:57.520664792Z',
			friends: [],
			welcome: true
		};
		await route.fulfill({
			body: JSON.stringify(user)
		});
	});
}
