import { expect, test } from '@playwright/test';
import { getGlobalGrumbles } from '../mocks/global';
import { getUser } from '../mocks/users';
import { getCatetories } from '../mocks/grumbles';
test.beforeEach(async ({ page }) => {
	await getUser(page);
	await getGlobalGrumbles(page);
	await getCatetories(page);
	await page.goto('/global');
});

test('HAPPY global grumbles page loaded correctly', async ({ page }) => {
	await expect(page.getByRole('heading').getByText('Global grumbles')).toBeVisible();
	await expect(page).toHaveScreenshot();
});

test('HAPPY new global grumble button opens modal', async ({ page }) => {
	await page.getByRole('button').getByText('New grumble').click();
	await expect(page.getByText('New global grumble')).toBeVisible();
	await expect(page).toHaveScreenshot();
});
