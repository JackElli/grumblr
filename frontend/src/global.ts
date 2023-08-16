import { page } from '$app/stores';
import { get } from 'svelte/store';

export const version = '0.1.0';

let ENV = 'dev';
export const IP =
	ENV === 'prod' ? 'ec2-16-170-201-214.eu-north-1.compute.amazonaws.com' : 'localhost';

export function dateDiff(date: string): string {
	const d = new Date(Date.parse(date)).getTime();
	const now = new Date().getTime();
	const secondDiff = (now - d) / 1000;
	if (secondDiff < 60) {
		return `${Math.round(secondDiff)} seconds ago`;
	}
	const minuteDiff = (now - d) / 60000;
	if (minuteDiff < 60) {
		return `${Math.round(minuteDiff)} minutes ago`;
	}
	const hourDiff = (now - d) / 3600000;
	if (hourDiff < 24) {
		return `${Math.round(hourDiff)} hours ago`;
	}
	return `${Math.round((now - d) / (3600 * 1000 * 24))} days ago`;
}

export function userIconText(username: string): string {
	return username[0].toUpperCase();
}

export function scrollToLastPos() {
	const _page = get(page);
	const scrollTo = _page.url.searchParams.get('scrollTo');
	setTimeout(function () {
		window.scrollTo({
			top: parseInt(scrollTo ?? '0')
		});
	}, 0);
}
