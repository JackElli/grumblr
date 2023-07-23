import { userStore, type User } from '$lib/stores/userStore';
import { redirect } from '@sveltejs/kit';

export async function Auth(): Promise<User> {
	const userId = '1f21823a-8682-4900-b627-d6bd39e1b95b';
	try {
		const resp = await fetch(`http://localhost:3200/user/${userId}`, {
			method: 'GET',
			credentials: 'include'
		});
		const user = await resp.json();
		userStore.set(user);
		return user;
	} catch (e) {
		console.error(e);
		throw redirect(302, '/login');
	}
}
