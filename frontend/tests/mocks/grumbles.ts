import type { Page } from '@playwright/test';
import type { _Grumble } from '../../src/routes/(main)/grumbles/grumbles';

export async function getGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles/*', async (route) => {
		await route.fulfill({
			json: [
				{
					id: 'test123',
					createdBy: 'user:1',
					message: 'this is the very first grumble',
					dateCreated: '2023-06-03T20:24:35.060986337Z',
					comments: [
						{
							createdBy: 'asdasd',
							message: 'test comment',
							dateCreated: '023-07-22T14:21:36.92541872Z'
						}
					],
					agrees: {},
					disagrees: {},
					category: 'testing'
				}
			]
		});
	});
}

export async function getCatetories(page: Page) {
	await page.route('http://localhost:3200/grumbles/info/categories/*', async (route) => {
		await route.fulfill({
			json: [
				{
					id: 'testcat1',
					type: 'friends',
					people: ['jack'],
					name: 'Weather'
				}
			]
		});
	});
}

export async function getNoGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles/*', async (route) => {
		await route.fulfill({
			json: []
		});
	});
}

export async function getLongGrumbles(page: Page) {
	await page.route('http://localhost:3200/grumbles/*', async (route) => {
		await route.fulfill({
			json: [
				{
					createdBy: 'user:1',
					message:
						'This is a very long grumble This is a very long grumble This is a very long grumble This is a very long grumble This is a very long grumble',
					dateCreated: '2023-06-03T20:24:35.060986337Z',
					comments: [
						{
							createdBy: 'asdasd',
							message: 'test comment',
							dateCreated: '023-07-22T14:21:36.92541872Z'
						}
					],
					agrees: {},
					disagrees: {},
				}
			]
		});
	});
}
