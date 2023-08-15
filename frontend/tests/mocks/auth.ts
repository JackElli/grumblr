import type { Page } from '@playwright/test';

export async function auth(page: Page) {
	await page.route('http://localhost:3200/auth', async (route) => {
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
