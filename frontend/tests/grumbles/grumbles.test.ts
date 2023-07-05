import { expect, test } from '@playwright/test';
import { getGrumbles, getLongGrumbles, getNoGrumbles } from '../mocks/grumbles';
import { getUser } from '../mocks/users';

test.beforeEach(async ({ page }) => {
	await getUser(page)
	await getGrumbles(page);
});

test('HAPPY friends grumbles page loaded correctly', async ({ page }) => {
	await page.goto('/grumbles');
	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY long grumble', async ({ page }) => {
	await getLongGrumbles(page);
	await page.goto('/grumbles');
	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY new grumble button opens modal', async ({ page }) => {
	await page.goto('/grumbles');
	await page.getByRole('button').getByText('New grumble').click();
	await expect(page.getByText('Add your grumble text')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('NEGATIVE no grumbles available', async ({ page }) => {
	await getNoGrumbles(page);
	await page.goto('/grumbles');
	await expect(page.getByText('No grumbles found here.')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
