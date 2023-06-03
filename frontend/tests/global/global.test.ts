import { expect, test } from '@playwright/test';

test('global grumbles page loaded correctly', async ({ page }) => {
	await page.goto('/global');
	await expect(page.getByRole('heading').getByText('Global grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
