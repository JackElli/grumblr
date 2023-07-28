export const version = '0.0.2';
// export const IP = 'ec2-16-171-174-3.eu-north-1.compute.amazonaws.com'
export const IP = 'localhost'

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
