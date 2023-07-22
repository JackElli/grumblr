import { expect, test } from '@playwright/test';
import { getUser } from '../mocks/users';

test.beforeEach(async ({ page }) => {
	await getUser(page);
	await page.goto('/profile');
});

test('profile page loaded correctly', async ({ page }) => {
	await page.goto('/profile');
	await expect(page.getByRole('heading').getByText('Profile')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
