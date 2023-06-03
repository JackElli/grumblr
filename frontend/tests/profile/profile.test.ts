import { expect, test } from '@playwright/test';

test('profile page loaded correctly', async ({ page }) => {
	await page.goto('/profile');
	await expect(page.getByRole('heading').getByText('Profile')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
