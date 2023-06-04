import { expect, test } from '@playwright/test';
import { getGrumbles, getLongGrumbles } from '../mocks/grumbles';

test('friends grumbles page loaded correctly', async ({ page }) => {
	await getGrumbles(page);
	await page.goto('/grumbles');
	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('long grumble', async ({ page }) => {
	await getLongGrumbles(page);
	await page.goto('/grumbles');
	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
