import { expect, test } from '@playwright/test';

test('friends grumbles page loaded correctly', async ({ page }) => {
    await page.goto('/grumbles');
    await expect(page.getByRole("heading").getByText("Friends grumbles")).toBeVisible();
    await expect(page).toHaveScreenshot()
});
