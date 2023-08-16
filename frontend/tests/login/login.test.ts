import { expect, test } from '@playwright/test';

test.beforeEach(async ({ page }) => {
	await page.goto('/login');
});

test('login page loaded correctly', async ({ page }) => {
	await expect(page.getByRole('heading').getByText('What will you grumble about?')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
