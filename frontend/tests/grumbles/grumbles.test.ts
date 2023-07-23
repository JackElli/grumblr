import { expect, test } from '@playwright/test';
import { getCatetories, getGrumbles, getLongGrumbles, getNoGrumbles } from '../mocks/grumbles';
import { getUser } from '../mocks/users';

test.beforeEach(async ({ page }) => {
	await getUser(page);
	await getGrumbles(page);
	await getCatetories(page);
	await page.goto('/grumbles');
});

test('HAPPY friends grumbles page loaded correctly', async ({ page }) => {
	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY long grumble', async ({ page }) => {
	await getLongGrumbles(page);

	await expect(page.getByRole('heading').getByText('Friends grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY new grumble button opens modal', async ({ page }) => {
	await page.getByRole('button').getByText('New grumble').click();
	await expect(page.getByText('Add your grumble text')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('NEGATIVE no grumbles available', async ({ page }) => {
	await getNoGrumbles(page);
	await expect(page.getByText('No grumbles found here.')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY comments button shows comments', async ({ page }) => {
	await page.getByText('1 comment').click();
	await expect(page.getByText('Viewing grumble')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
