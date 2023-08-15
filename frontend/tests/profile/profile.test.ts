import { expect, test } from '@playwright/test';
import { getUser } from '../mocks/users';
import { auth } from '../mocks/auth';
import { getGlobalCategories } from '../mocks/global';

test.beforeEach(async ({ page }) => {
	await auth(page);
	await getGlobalCategories(page);

	await page.goto('/profile');
});

test('profile page loaded correctly', async ({ page }) => {
	await expect(page.getByRole('heading').getByText('Profile')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
